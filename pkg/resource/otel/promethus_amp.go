package otel

import (
	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/eksdemo/pkg/aws"
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/manifest"
	"github.com/awslabs/eksdemo/pkg/resource"
	"github.com/awslabs/eksdemo/pkg/resource/amp_workspace"
	"github.com/awslabs/eksdemo/pkg/template"
)

type PrometheusAMPOptions struct {
	resource.CommonOptions

	AmpAlias    string
	AmpEndpoint string
}

// As of 9/8/24 this is broken, errors with:
//
//	Error: failed to resolve config: cannot resolve the configuration: cannot convert the confmap.Conf:
//	environment variable "1" has invalid name: must match regex ^[a-zA-Z_][a-zA-Z0-9_]*$
//
// May be fixed in v0.105.0 in PR https://github.com/open-telemetry/opentelemetry-collector/pull/10560
func NewPrometheusAMPCollector() *resource.Resource {
	options := &PrometheusAMPOptions{
		CommonOptions: resource.CommonOptions{
			Name:           "prometheus-amp",
			Namespace:      "adot-system",
			NamespaceFlag:  true,
			ServiceAccount: "adot-collector",
		},
	}

	return &resource.Resource{
		Command: cmd.Command{
			Name:        "prometheus-amp",
			Description: "Prometheus AMP Collector",
			Aliases:     []string{"prom-amp", "promamp"},
		},

		CreateFlags: []cmd.Flag{
			&cmd.StringFlag{
				CommandFlag: cmd.CommandFlag{
					Name:        "alias",
					Description: "AMP workspace alias",
					Required:    true,
				},
				Option: &options.AmpAlias,
			},
		},

		Manager: &manifest.ResourceManager{
			Template: &template.TextTemplate{
				Template: promAMPCollectorTemplate,
			},
		},

		Options: options,
	}
}

// https://github.com/aws-observability/aws-otel-community/blob/master/sample-configs/operator/collector-config-amp.yaml
const promAMPCollectorTemplate = `---
#
# OpenTelemetry Collector configuration
# Metrics pipeline with Prometheus Receiver and Prometheus Remote Write Exporter sending metrics to Amazon Managed Prometheus
#
---
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  namespace: {{ .Namespace }}
  name: {{ .Name }}
spec:
  mode: deployment
  serviceAccount:  {{ .ServiceAccount }}
  podAnnotations:
    prometheus.io/scrape: 'true'
    prometheus.io/port: '8888'
  resources:
    requests:
      cpu: "1"
    limits:
      cpu: "1"
  config:
    extensions:
      sigv4auth:
        region: {{ .Region }}
        service: "aps"

    receivers:
      #
      # Scrape configuration for the Prometheus Receiver
      # This is the same configuration used when Prometheus is installed using the community Helm chart
      #
      prometheus:
        config:
          global:
            scrape_interval: 15s
            scrape_timeout: 10s

          scrape_configs:
          - job_name: kubernetes-apiservers
            bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
            kubernetes_sd_configs:
            - role: endpoints
            relabel_configs:
            - action: keep
              regex: default;kubernetes;https
              source_labels:
              - __meta_kubernetes_namespace
              - __meta_kubernetes_service_name
              - __meta_kubernetes_endpoint_port_name
            scheme: https
            tls_config:
              ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
              insecure_skip_verify: true

          - job_name: kubernetes-nodes
            bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
            kubernetes_sd_configs:
            - role: node
            relabel_configs:
            - action: labelmap
              regex: __meta_kubernetes_node_label_(.+)
            - replacement: kubernetes.default.svc:443
              target_label: __address__
            - regex: (.+)
              replacement: /api/v1/nodes/$$1/proxy/metrics
              source_labels:
              - __meta_kubernetes_node_name
              target_label: __metrics_path__
            scheme: https
            tls_config:
              ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
              insecure_skip_verify: true

          - job_name: kubernetes-nodes-cadvisor
            bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token
            kubernetes_sd_configs:
            - role: node
            relabel_configs:
            - action: labelmap
              regex: __meta_kubernetes_node_label_(.+)
            - replacement: kubernetes.default.svc:443
              target_label: __address__
            - regex: (.+)
              replacement: /api/v1/nodes/$$1/proxy/metrics/cadvisor
              source_labels:
              - __meta_kubernetes_node_name
              target_label: __metrics_path__
            scheme: https
            tls_config:
              ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt
              insecure_skip_verify: true

          - job_name: kubernetes-service-endpoints
            kubernetes_sd_configs:
            - role: endpoints
            relabel_configs:
            - action: keep
              regex: true
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_scrape
            - action: replace
              regex: (https?)
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_scheme
              target_label: __scheme__
            - action: replace
              regex: (.+)
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_path
              target_label: __metrics_path__
            - action: replace
              regex: ([^:]+)(?::\d+)?;(\d+)
              replacement: $$1:$$2
              source_labels:
              - __address__
              - __meta_kubernetes_service_annotation_prometheus_io_port
              target_label: __address__
            - action: labelmap
              regex: __meta_kubernetes_service_annotation_prometheus_io_param_(.+)
              replacement: __param_$$1
            - action: labelmap
              regex: __meta_kubernetes_service_label_(.+)
            - action: replace
              source_labels:
              - __meta_kubernetes_namespace
              target_label: kubernetes_namespace
            - action: replace
              source_labels:
              - __meta_kubernetes_service_name
              target_label: kubernetes_name
            - action: replace
              source_labels:
              - __meta_kubernetes_pod_node_name
              target_label: kubernetes_node

          - job_name: kubernetes-service-endpoints-slow
            kubernetes_sd_configs:
            - role: endpoints
            relabel_configs:
            - action: keep
              regex: true
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_scrape_slow
            - action: replace
              regex: (https?)
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_scheme
              target_label: __scheme__
            - action: replace
              regex: (.+)
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_path
              target_label: __metrics_path__
            - action: replace
              regex: ([^:]+)(?::\d+)?;(\d+)
              replacement: $$1:$$2
              source_labels:
              - __address__
              - __meta_kubernetes_service_annotation_prometheus_io_port
              target_label: __address__
            - action: labelmap
              regex: __meta_kubernetes_service_annotation_prometheus_io_param_(.+)
              replacement: __param_$$1
            - action: labelmap
              regex: __meta_kubernetes_service_label_(.+)
            - action: replace
              source_labels:
              - __meta_kubernetes_namespace
              target_label: kubernetes_namespace
            - action: replace
              source_labels:
              - __meta_kubernetes_service_name
              target_label: kubernetes_name
            - action: replace
              source_labels:
              - __meta_kubernetes_pod_node_name
              target_label: kubernetes_node
            scrape_interval: 5m
            scrape_timeout: 30s

          - job_name: prometheus-pushgateway
            kubernetes_sd_configs:
            - role: service
            relabel_configs:
            - action: keep
              regex: pushgateway
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_probe

          - job_name: kubernetes-services
            kubernetes_sd_configs:
            - role: service
            metrics_path: /probe
            params:
              module:
              - http_2xx
            relabel_configs:
            - action: keep
              regex: true
              source_labels:
              - __meta_kubernetes_service_annotation_prometheus_io_probe
            - source_labels:
              - __address__
              target_label: __param_target
            - replacement: blackbox
              target_label: __address__
            - source_labels:
              - __param_target
              target_label: instance
            - action: labelmap
              regex: __meta_kubernetes_service_label_(.+)
            - source_labels:
              - __meta_kubernetes_namespace
              target_label: kubernetes_namespace
            - source_labels:
              - __meta_kubernetes_service_name
              target_label: kubernetes_name

          - job_name: kubernetes-pods
            kubernetes_sd_configs:
            - role: pod
            relabel_configs:
            - action: keep
              regex: true
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_scrape
            - action: replace
              regex: (https?)
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_scheme
              target_label: __scheme__
            - action: replace
              regex: (.+)
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_path
              target_label: __metrics_path__
            - action: replace
              regex: ([^:]+)(?::\d+)?;(\d+)
              replacement: $$1:$$2
              source_labels:
              - __address__
              - __meta_kubernetes_pod_annotation_prometheus_io_port
              target_label: __address__
            - action: labelmap
              regex: __meta_kubernetes_pod_annotation_prometheus_io_param_(.+)
              replacement: __param_$$1
            - action: labelmap
              regex: __meta_kubernetes_pod_label_(.+)
            - action: replace
              source_labels:
              - __meta_kubernetes_namespace
              target_label: kubernetes_namespace
            - action: replace
              source_labels:
              - __meta_kubernetes_pod_name
              target_label: kubernetes_pod_name
            - action: drop
              regex: Pending|Succeeded|Failed|Completed
              source_labels:
              - __meta_kubernetes_pod_phase

          - job_name: kubernetes-pods-slow
            scrape_interval: 5m
            scrape_timeout: 30s
            kubernetes_sd_configs:
            - role: pod
            relabel_configs:
            - action: keep
              regex: true
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_scrape_slow
            - action: replace
              regex: (https?)
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_scheme
              target_label: __scheme__
            - action: replace
              regex: (.+)
              source_labels:
              - __meta_kubernetes_pod_annotation_prometheus_io_path
              target_label: __metrics_path__
            - action: replace
              regex: ([^:]+)(?::\d+)?;(\d+)
              replacement: $$1:$$2
              source_labels:
              - __address__
              - __meta_kubernetes_pod_annotation_prometheus_io_port
              target_label: __address__
            - action: labelmap
              regex: __meta_kubernetes_pod_annotation_prometheus_io_param_(.+)
              replacement: __param_$1
            - action: labelmap
              regex: __meta_kubernetes_pod_label_(.+)
            - action: replace
              source_labels:
              - __meta_kubernetes_namespace
              target_label: namespace
            - action: replace
              source_labels:
              - __meta_kubernetes_pod_name
              target_label: pod
            - action: drop
              regex: Pending|Succeeded|Failed|Completed
              source_labels:
              - __meta_kubernetes_pod_phase

    processors:
      batch/metrics:
        timeout: 60s

    exporters:
      prometheusremotewrite:
        endpoint: {{ .AmpEndpoint }}api/v1/remote_write
        auth:
          authenticator: sigv4auth

    service:
      extensions: [sigv4auth]
      pipelines:
        metrics:
          receivers: [prometheus]
          processors: [batch/metrics]
          exporters: [prometheusremotewrite]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: otel-prometheus-amp-role
rules:
  - apiGroups:
      - ""
    resources:
      - nodes
      - nodes/proxy
      - services
      - endpoints
      - pods
    verbs:
      - get
      - list
      - watch
  - apiGroups:
      - extensions
    resources:
      - ingresses
    verbs:
      - get
      - list
      - watch
  - nonResourceURLs:
      - /metrics
    verbs:
      - get
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: otel-prometheus-amp-role-binding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: otel-prometheus-amp-role
subjects:
  - kind: ServiceAccount
    name: {{ .ServiceAccount }}
    namespace: {{ .Namespace }}
`

func (o *PrometheusAMPOptions) PreCreate() error {
	workspace, err := amp_workspace.NewGetter(aws.NewAMPClient()).GetAmpByAlias(o.AmpAlias)
	if err != nil {
		return err
	}

	o.AmpEndpoint = awssdk.ToString(workspace.Workspace.PrometheusEndpoint)

	return nil
}
