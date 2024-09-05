package karpenter

// JSON: https://github.com/aws/karpenter-provider-aws/blob/main/website/content/en/preview/getting-started/getting-started-with-karpenter/karpenter-capacity-dashboard.json
//
//	kubectl create configmap karpenter-capacity \
//	  --from-file=./karpenter-capacity-dashboard.json \
//	  --dry-run=client -o yaml > capacity_dashboard.yaml
//
// Must double escape because textTemplate transforms and Helm uses the same Golang template
// TextTemplate turns: "{{ "{{" }} {{ "{{cluster}}" | printf "%q" }} {{ "}}" }}" _to_ "{{ "{{cluster}}" }}"
// and Helm turns it _to_ {{cluster}}
const capacityDashboardTemplate = `---
apiVersion: v1
kind: ConfigMap
metadata:
  name: karpenter-capacity
data:
  karpenter-capacity-dashboard.json: |
    {
        "annotations": {
          "list": [
            {
              "builtIn": 1,
              "datasource": "-- Grafana --",
              "enable": true,
              "hide": true,
              "iconColor": "rgba(0, 211, 255, 1)",
              "name": "Annotations & Alerts",
              "target": {
                "limit": 100,
                "matchAny": false,
                "tags": [],
                "type": "dashboard"
              },
              "type": "dashboard"
            }
          ]
        },
        "editable": true,
        "fiscalYearStartMonth": 0,
        "graphTooltip": 0,
        "id": 32,
        "links": [],
        "liveNow": true,
        "panels": [
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "description": "",
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "Nodes rate",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "none"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "min": 0,
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                },
                "unit": "none"
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byFrameRefID",
                    "options": "B"
                  },
                  "properties": [
                    {
                      "id": "custom.axisPlacement",
                      "value": "right"
                    },
                    {
                      "id": "custom.axisLabel",
                      "value": "Nodes total"
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 5,
              "w": 24,
              "x": 0,
              "y": 0
            },
            "id": 14,
            "options": {
              "legend": {
                "calcs": [
                  "lastNotNull"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum(rate(karpenter_nodes_created_total{nodepool=~\"$nodepool\"}[$__rate_interval])) by(cluster,nodepool)",
                "format": "time_series",
                "legendFormat": "{{ "{{" }} {{ "{{cluster}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by(cluster,nodepool) (karpenter_nodes_created_total{nodepool=~\"$nodepool\"})",
                "hide": false,
                "instant": false,
                "legendFormat": "__auto",
                "range": true,
                "refId": "B"
              }
            ],
            "title": "Nodes Created by nodepool",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "description": "",
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "Nodes rate",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "none"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "min": 0,
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                },
                "unit": "none"
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byFrameRefID",
                    "options": "B"
                  },
                  "properties": [
                    {
                      "id": "custom.axisPlacement",
                      "value": "right"
                    },
                    {
                      "id": "custom.axisLabel",
                      "value": "Nodes total"
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 5,
              "w": 24,
              "x": 0,
              "y": 5
            },
            "id": 15,
            "options": {
              "legend": {
                "calcs": [
                  "lastNotNull"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum(rate(karpenter_nodes_terminated_total{nodepool=~\"$nodepool\"}[$__rate_interval])) by(cluster,nodepool)",
                "format": "time_series",
                "legendFormat": "__auto",
                "range": true,
                "refId": "A"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by(cluster,nodepool) (karpenter_nodes_terminated_total{nodepool=~\"$nodepool\"})",
                "hide": false,
                "instant": false,
                "legendFormat": "__auto",
                "range": true,
                "refId": "B"
              }
            ],
            "title": "Nodes Terminated by nodepool",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "none"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": []
            },
            "gridPos": {
              "h": 6,
              "w": 24,
              "x": 0,
              "y": 10
            },
            "id": 12,
            "options": {
              "legend": {
                "calcs": [
                  "last"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "sum by(reason)(karpenter_voluntary_disruption_eligible_nodes)",
                "legendFormat": "reason={{ "{{" }} {{ "{{reason}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              }
            ],
            "title": "Nodes Eligibale for Disruptions by \"reason\"",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "description": "",
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "Disruptions rate",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "normal"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byFrameRefID",
                    "options": "B"
                  },
                  "properties": [
                    {
                      "id": "custom.axisPlacement",
                      "value": "right"
                    },
                    {
                      "id": "custom.axisLabel",
                      "value": "Disruptions total"
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 6,
              "w": 24,
              "x": 0,
              "y": 16
            },
            "id": 17,
            "options": {
              "legend": {
                "calcs": [
                  "last"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "sum(rate(karpenter_nodeclaims_disrupted_total{nodepool=~\"$nodepool\"}[$__rate_interval])) by(cluster,nodepool)",
                "legendFormat": "{{ "{{" }} {{ "{{label_name}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by(cluster,nodepool)(karpenter_nodeclaims_disrupted_total{nodepool=~\"$nodepool\"})",
                "hide": false,
                "instant": false,
                "legendFormat": "{{ "{{" }} {{ "{{label_name}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "B"
              }
            ],
            "title": "Node Disruptions by \"$distribution_filter\"",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "Disruptions rate",
                  "axisPlacement": "left",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "normal"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byFrameRefID",
                    "options": "B"
                  },
                  "properties": [
                    {
                      "id": "custom.axisPlacement",
                      "value": "right"
                    },
                    {
                      "id": "custom.axisLabel",
                      "value": "Disruptions total"
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 6,
              "w": 24,
              "x": 0,
              "y": 22
            },
            "id": 19,
            "options": {
              "legend": {
                "calcs": [
                  "last"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": true,
                "expr": "sum(rate(karpenter_voluntary_disruption_decisions_total[$__rate_interval])) by(cluster,nodepool,consolidation_type,decision)",
                "hide": false,
                "legendFormat": "type={{ "{{" }} {{ "{{consolidation_type}}" | printf "%q" }} {{ "}}" }}, decision={{ "{{" }} {{ "{{decision}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by(consolidation_type,decision)(karpenter_voluntary_disruption_decisions_total)",
                "hide": true,
                "instant": false,
                "legendFormat": "{{ "{{" }} {{ "{{label_name}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "B"
              }
            ],
            "title": "Disruption Actions performed by \"type\" and \"decision\"",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "normal"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": []
            },
            "gridPos": {
              "h": 8,
              "w": 24,
              "x": 0,
              "y": 28
            },
            "id": 6,
            "options": {
              "legend": {
                "calcs": [],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by ($distribution_filter)(\n    karpenter_pods_state{arch=~\"$arch\", capacity_type=~\"$capacity_type\", instance_type=~\"$instance_type\", nodepool=~\"$nodepool\"}\n)",
                "legendFormat": "{{ "{{" }} {{ "{{label_name}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              }
            ],
            "title": "Pod Distribution by \"$distribution_filter\"",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 10,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "normal"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": []
            },
            "gridPos": {
              "h": 6,
              "w": 24,
              "x": 0,
              "y": 36
            },
            "id": 20,
            "options": {
              "legend": {
                "calcs": [
                  "last"
                ],
                "displayMode": "table",
                "placement": "right",
                "showLegend": true
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "sum by(phase)(karpenter_pods_state)",
                "legendFormat": "{{ "{{" }} {{ "{{label_name}}" | printf "%q" }} {{ "}}" }}",
                "range": true,
                "refId": "A"
              }
            ],
            "title": "Pod Phase",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "continuous-RdYlGr"
                },
                "custom": {
                  "align": "left",
                  "cellOptions": {
                    "type": "auto"
                  },
                  "inspect": false
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byRegexp",
                    "options": ".*Utilization$"
                  },
                  "properties": [
                    {
                      "id": "custom.cellOptions",
                      "value": {
                        "mode": "gradient",
                        "type": "gauge"
                      }
                    },
                    {
                      "id": "min",
                      "value": 0
                    },
                    {
                      "id": "max",
                      "value": 1
                    },
                    {
                      "id": "unit",
                      "value": "percentunit"
                    }
                  ]
                },
                {
                  "matcher": {
                    "id": "byName",
                    "options": "Memory Provisioned"
                  },
                  "properties": [
                    {
                      "id": "unit",
                      "value": "bytes"
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 11,
              "w": 18,
              "x": 0,
              "y": 42
            },
            "id": 10,
            "options": {
              "cellHeight": "sm",
              "footer": {
                "countRows": false,
                "fields": "",
                "reducer": [
                  "sum"
                ],
                "show": false
              },
              "showHeader": true
            },
            "pluginVersion": "11.1.0",
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "karpenter_nodepool_usage{resource_type=\"cpu\"} / karpenter_nodepool_limit{resource_type=\"cpu\"}",
                "format": "table",
                "instant": true,
                "legendFormat": "CPU Limit Utilization",
                "range": false,
                "refId": "CPU Limit Utilization"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "count by (nodepool)(karpenter_nodes_allocatable{nodepool!=\"N/A\",resource_type=\"cpu\"}) # Selects a single resource type to get node count",
                "format": "table",
                "hide": false,
                "instant": true,
                "range": false,
                "refId": "Node Count"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "karpenter_nodepool_usage{resource_type=\"memory\"} / karpenter_nodepool_limit{resource_type=\"memory\"}",
                "format": "table",
                "hide": false,
                "instant": true,
                "legendFormat": "Memory Limit Utilization",
                "range": false,
                "refId": "Memory Limit Utilization"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "sum by (nodepool)(karpenter_nodes_allocatable{nodepool!=\"N/A\",resource_type=\"cpu\"})",
                "format": "table",
                "hide": false,
                "instant": true,
                "range": false,
                "refId": "CPU Capacity"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "sum by (nodepool)(karpenter_nodes_allocatable{nodepool!=\"N/A\",resource_type=\"memory\"})",
                "format": "table",
                "hide": false,
                "instant": true,
                "range": false,
                "refId": "Memory Capacity"
              }
            ],
            "title": "Nodepool Summary",
            "transformations": [
              {
                "id": "seriesToColumns",
                "options": {
                  "byField": "nodepool"
                }
              },
              {
                "id": "organize",
                "options": {
                  "excludeByName": {
                    "Time": true,
                    "Time 1": true,
                    "Time 2": true,
                    "Time 3": true,
                    "Time 4": true,
                    "Time 5": true,
                    "__name__": true,
                    "instance": true,
                    "instance 1": true,
                    "instance 2": true,
                    "job": true,
                    "job 1": true,
                    "job 2": true,
                    "resource_type": true,
                    "resource_type 1": true,
                    "resource_type 2": true
                  },
                  "indexByName": {
                    "Time 1": 6,
                    "Time 2": 7,
                    "Time 3": 11,
                    "Time 4": 15,
                    "Time 5": 16,
                    "Value #CPU Capacity": 2,
                    "Value #CPU Limit Utilization": 3,
                    "Value #Memory Capacity": 4,
                    "Value #Memory Limit Utilization": 5,
                    "Value #Node Count": 1,
                    "instance 1": 8,
                    "instance 2": 12,
                    "job 1": 9,
                    "job 2": 13,
                    "nodepool": 0,
                    "resource_type 1": 10,
                    "resource_type 2": 14
                  },
                  "renameByName": {
                    "Time 1": "",
                    "Value": "CPU Utilization",
                    "Value #CPU Capacity": "CPU Provisioned",
                    "Value #CPU Limit Utilization": "CPU Limit Utilization",
                    "Value #CPU Utilization": "CPU Limit Utilization",
                    "Value #Memory Capacity": "Memory Provisioned",
                    "Value #Memory Limit Utilization": "Memory Limit Utilization",
                    "Value #Memory Utilization": "Memory Utilization",
                    "Value #Node Count": "Node Count",
                    "instance": "",
                    "instance 1": "",
                    "job": "",
                    "nodepool": "Nodepool"
                  }
                }
              }
            ],
            "type": "table"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "palette-classic"
                },
                "custom": {
                  "axisBorderShow": false,
                  "axisCenteredZero": false,
                  "axisColorMode": "text",
                  "axisLabel": "",
                  "axisPlacement": "auto",
                  "barAlignment": 0,
                  "drawStyle": "line",
                  "fillOpacity": 0,
                  "gradientMode": "none",
                  "hideFrom": {
                    "legend": false,
                    "tooltip": false,
                    "viz": false
                  },
                  "insertNulls": false,
                  "lineInterpolation": "linear",
                  "lineWidth": 1,
                  "pointSize": 5,
                  "scaleDistribution": {
                    "type": "linear"
                  },
                  "showPoints": "auto",
                  "spanNulls": false,
                  "stacking": {
                    "group": "A",
                    "mode": "none"
                  },
                  "thresholdsStyle": {
                    "mode": "off"
                  }
                },
                "mappings": [],
                "max": 1,
                "min": 0,
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                },
                "unit": "percentunit"
              },
              "overrides": []
            },
            "gridPos": {
              "h": 11,
              "w": 6,
              "x": 18,
              "y": 42
            },
            "id": 8,
            "options": {
              "legend": {
                "calcs": [],
                "displayMode": "list",
                "placement": "bottom",
                "showLegend": false
              },
              "tooltip": {
                "mode": "single",
                "sort": "none"
              }
            },
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "expr": "(count(karpenter_nodes_allocatable{arch=~\"$arch\",capacity_type=\"spot\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"}) or vector(0)) / count(karpenter_nodes_allocatable{arch=~\"$arch\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"})",
                "legendFormat": "Percentage",
                "range": true,
                "refId": "A"
              }
            ],
            "title": "Spot Node Percentage",
            "type": "timeseries"
          },
          {
            "datasource": {
              "type": "prometheus",
              "uid": "${datasource}"
            },
            "fieldConfig": {
              "defaults": {
                "color": {
                  "mode": "continuous-RdYlGr"
                },
                "custom": {
                  "align": "left",
                  "cellOptions": {
                    "type": "auto"
                  },
                  "inspect": false
                },
                "mappings": [],
                "thresholds": {
                  "mode": "absolute",
                  "steps": [
                    {
                      "color": "green",
                      "value": null
                    },
                    {
                      "color": "red",
                      "value": 80
                    }
                  ]
                }
              },
              "overrides": [
                {
                  "matcher": {
                    "id": "byName",
                    "options": "node_name"
                  },
                  "properties": [
                    {
                      "id": "custom.width",
                      "value": 333
                    }
                  ]
                },
                {
                  "matcher": {
                    "id": "byRegexp",
                    "options": ".*Utilization"
                  },
                  "properties": [
                    {
                      "id": "custom.cellOptions",
                      "value": {
                        "mode": "gradient",
                        "type": "gauge"
                      }
                    },
                    {
                      "id": "unit",
                      "value": "percentunit"
                    },
                    {
                      "id": "min",
                      "value": 0
                    },
                    {
                      "id": "thresholds",
                      "value": {
                        "mode": "percentage",
                        "steps": [
                          {
                            "color": "green",
                            "value": null
                          },
                          {
                            "color": "red",
                            "value": 75
                          }
                        ]
                      }
                    },
                    {
                      "id": "max",
                      "value": 1
                    }
                  ]
                },
                {
                  "matcher": {
                    "id": "byName",
                    "options": "Uptime"
                  },
                  "properties": [
                    {
                      "id": "unit",
                      "value": "s"
                    },
                    {
                      "id": "decimals",
                      "value": 0
                    }
                  ]
                }
              ]
            },
            "gridPos": {
              "h": 9,
              "w": 24,
              "x": 0,
              "y": 53
            },
            "id": 4,
            "options": {
              "cellHeight": "sm",
              "footer": {
                "countRows": false,
                "fields": "",
                "reducer": [
                  "sum"
                ],
                "show": false
              },
              "showHeader": true,
              "sortBy": [
                {
                  "desc": true,
                  "displayName": "Uptime"
                }
              ]
            },
            "pluginVersion": "11.1.0",
            "targets": [
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "((karpenter_nodes_total_daemon_requests{resource_type=\"cpu\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"} or karpenter_nodes_allocatable*0) + \n(karpenter_nodes_total_pod_requests{resource_type=\"cpu\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"} or karpenter_nodes_allocatable*0)) / \nkarpenter_nodes_allocatable{resource_type=\"cpu\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"}",
                "format": "table",
                "hide": false,
                "instant": true,
                "legendFormat": "CPU Utilization",
                "range": false,
                "refId": "CPU Utilization"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "((karpenter_nodes_total_daemon_requests{resource_type=\"memory\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"} or karpenter_nodes_allocatable*0) + \n(karpenter_nodes_total_pod_requests{resource_type=\"memory\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"} or karpenter_nodes_allocatable*0)) / \nkarpenter_nodes_allocatable{resource_type=\"memory\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"}",
                "format": "table",
                "hide": false,
                "instant": true,
                "legendFormat": "Memory Utilization",
                "range": false,
                "refId": "Memory Utilization"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "karpenter_nodes_total_daemon_requests{resource_type=\"pods\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"} + \nkarpenter_nodes_total_pod_requests{resource_type=\"pods\",arch=~\"$arch\",capacity_type=~\"$capacity_type\",instance_type=~\"$instance_type\",nodepool=~\"$nodepool\",zone=~\"$zone\"}",
                "format": "table",
                "hide": false,
                "instant": true,
                "legendFormat": "Memory Utilization",
                "range": false,
                "refId": "Pod Count"
              },
              {
                "datasource": {
                  "type": "prometheus",
                  "uid": "${datasource}"
                },
                "editorMode": "code",
                "exemplar": false,
                "expr": "label_replace(\n    sum by (node)(node_time_seconds) - sum by (node)(node_boot_time_seconds),\n    \"node_name\", \"$1\", \"node\", \"(.+)\"\n)",
                "format": "table",
                "hide": false,
                "instant": true,
                "legendFormat": "Uptime",
                "range": false,
                "refId": "Uptime"
              }
            ],
            "title": "Node Summary",
            "transformations": [
              {
                "id": "seriesToColumns",
                "options": {
                  "byField": "node_name"
                }
              },
              {
                "id": "organize",
                "options": {
                  "excludeByName": {
                    "Time": true,
                    "Time 1": true,
                    "Time 2": true,
                    "Time 3": true,
                    "Time 4": true,
                    "Value": false,
                    "Value #Pod Count": false,
                    "__name__": true,
                    "arch": true,
                    "arch 1": true,
                    "arch 2": true,
                    "arch 3": true,
                    "capacity_type 2": true,
                    "capacity_type 3": true,
                    "instance": true,
                    "instance 1": true,
                    "instance 2": true,
                    "instance 3": true,
                    "instance_category 1": true,
                    "instance_category 2": true,
                    "instance_category 3": true,
                    "instance_cpu": true,
                    "instance_cpu 1": true,
                    "instance_cpu 2": true,
                    "instance_cpu 3": true,
                    "instance_family": true,
                    "instance_family 1": true,
                    "instance_family 2": true,
                    "instance_family 3": true,
                    "instance_generation 1": true,
                    "instance_generation 2": true,
                    "instance_generation 3": true,
                    "instance_gpu_count": true,
                    "instance_gpu_count 1": true,
                    "instance_gpu_count 2": true,
                    "instance_gpu_count 3": true,
                    "instance_gpu_manufacturer": true,
                    "instance_gpu_manufacturer 1": true,
                    "instance_gpu_manufacturer 2": true,
                    "instance_gpu_manufacturer 3": true,
                    "instance_gpu_memory": true,
                    "instance_gpu_memory 1": true,
                    "instance_gpu_memory 2": true,
                    "instance_gpu_memory 3": true,
                    "instance_gpu_name": true,
                    "instance_gpu_name 1": true,
                    "instance_gpu_name 2": true,
                    "instance_gpu_name 3": true,
                    "instance_hypervisor": true,
                    "instance_hypervisor 1": true,
                    "instance_hypervisor 2": true,
                    "instance_hypervisor 3": true,
                    "instance_local_nvme 1": true,
                    "instance_local_nvme 2": true,
                    "instance_local_nvme 3": true,
                    "instance_memory": true,
                    "instance_memory 1": true,
                    "instance_memory 2": true,
                    "instance_memory 3": true,
                    "instance_pods": true,
                    "instance_pods 1": true,
                    "instance_pods 2": true,
                    "instance_pods 3": true,
                    "instance_size": true,
                    "instance_size 1": true,
                    "instance_size 2": true,
                    "instance_size 3": true,
                    "instance_type 1": false,
                    "instance_type 2": true,
                    "instance_type 3": true,
                    "job": true,
                    "job 1": true,
                    "job 2": true,
                    "job 3": true,
                    "node": true,
                    "nodepool 1": false,
                    "nodepool 2": true,
                    "nodepool 3": true,
                    "os": true,
                    "os 1": true,
                    "os 2": true,
                    "os 3": true,
                    "resource_type": true,
                    "resource_type 1": true,
                    "resource_type 2": true,
                    "resource_type 3": true,
                    "zone 1": false,
                    "zone 2": true,
                    "zone 3": true
                  },
                  "indexByName": {
                    "Time 1": 1,
                    "Time 2": 25,
                    "Time 3": 45,
                    "Time 4": 65,
                    "Value #CPU Utilization": 10,
                    "Value #Memory Utilization": 11,
                    "Value #Pod Count": 9,
                    "Value #Uptime": 8,
                    "arch 1": 5,
                    "arch 2": 26,
                    "arch 3": 46,
                    "capacity_type 1": 6,
                    "capacity_type 2": 27,
                    "capacity_type 3": 47,
                    "instance 1": 4,
                    "instance 2": 28,
                    "instance 3": 48,
                    "instance_cpu 1": 12,
                    "instance_cpu 2": 29,
                    "instance_cpu 3": 49,
                    "instance_family 1": 13,
                    "instance_family 2": 30,
                    "instance_family 3": 50,
                    "instance_gpu_count 1": 14,
                    "instance_gpu_count 2": 31,
                    "instance_gpu_count 3": 51,
                    "instance_gpu_manufacturer 1": 15,
                    "instance_gpu_manufacturer 2": 32,
                    "instance_gpu_manufacturer 3": 52,
                    "instance_gpu_memory 1": 16,
                    "instance_gpu_memory 2": 33,
                    "instance_gpu_memory 3": 53,
                    "instance_gpu_name 1": 17,
                    "instance_gpu_name 2": 34,
                    "instance_gpu_name 3": 54,
                    "instance_hypervisor 1": 18,
                    "instance_hypervisor 2": 35,
                    "instance_hypervisor 3": 55,
                    "instance_memory 1": 19,
                    "instance_memory 2": 36,
                    "instance_memory 3": 56,
                    "instance_pods 1": 20,
                    "instance_pods 2": 37,
                    "instance_pods 3": 57,
                    "instance_size 1": 21,
                    "instance_size 2": 38,
                    "instance_size 3": 58,
                    "instance_type 1": 3,
                    "instance_type 2": 39,
                    "instance_type 3": 59,
                    "job 1": 22,
                    "job 2": 40,
                    "job 3": 60,
                    "node": 66,
                    "node_name": 0,
                    "nodepool 1": 2,
                    "nodepool 2": 42,
                    "nodepool 3": 62,
                    "os 1": 23,
                    "os 2": 41,
                    "os 3": 61,
                    "resource_type 1": 24,
                    "resource_type 2": 43,
                    "resource_type 3": 63,
                    "zone 1": 7,
                    "zone 2": 44,
                    "zone 3": 64
                  },
                  "renameByName": {
                    "Time": "",
                    "Time 1": "",
                    "Value": "CPU Utilization",
                    "Value #Allocatable": "",
                    "Value #CPU Utilization": "CPU Utilization",
                    "Value #Memory Utilization": "Memory Utilization",
                    "Value #Pod CPU": "",
                    "Value #Pod Count": "Pods",
                    "Value #Uptime": "Uptime",
                    "arch": "Architecture",
                    "arch 1": "Arch",
                    "capacity_type": "Capacity Type",
                    "capacity_type 1": "Capacity Type",
                    "instance 1": "Instance",
                    "instance_cpu 1": "vCPU",
                    "instance_type": "Instance Type",
                    "instance_type 1": "Instance Type",
                    "node_name": "Node Name",
                    "nodepool 1": "Nodepool",
                    "zone 1": "Zone"
                  }
                }
              }
            ],
            "type": "table"
          }
        ],
        "refresh": "10s",
        "schemaVersion": 39,
        "tags": [],
        "templating": {
          "list": [
            {
              "current": {
                "selected": false,
                "text": "Prometheus",
                "value": "prometheus"
              },
              "hide": 0,
              "includeAll": false,
              "label": "Data Source",
              "multi": false,
              "name": "datasource",
              "options": [],
              "query": "prometheus",
              "refresh": 1,
              "regex": "",
              "skipUrlSync": false,
              "type": "datasource"
            },
            {
              "current": {
                "selected": true,
                "text": [
                  "All"
                ],
                "value": [
                  "$__all"
                ]
              },
              "datasource": {
                "type": "prometheus",
                "uid": "${datasource}"
              },
              "definition": "label_values(karpenter_nodes_allocatable, nodepool)",
              "hide": 0,
              "includeAll": true,
              "multi": true,
              "name": "nodepool",
              "options": [],
              "query": {
                "query": "label_values(karpenter_nodes_allocatable, nodepool)",
                "refId": "StandardVariableQuery"
              },
              "refresh": 2,
              "regex": "",
              "skipUrlSync": false,
              "sort": 1,
              "type": "query"
            },
            {
              "current": {
                "selected": true,
                "text": [
                  "All"
                ],
                "value": [
                  "$__all"
                ]
              },
              "datasource": {
                "type": "prometheus",
                "uid": "${datasource}"
              },
              "definition": "label_values(karpenter_nodes_allocatable, zone)",
              "hide": 0,
              "includeAll": true,
              "multi": true,
              "name": "zone",
              "options": [],
              "query": {
                "query": "label_values(karpenter_nodes_allocatable, zone)",
                "refId": "StandardVariableQuery"
              },
              "refresh": 2,
              "regex": "",
              "skipUrlSync": false,
              "sort": 1,
              "type": "query"
            },
            {
              "current": {
                "selected": true,
                "text": [
                  "All"
                ],
                "value": [
                  "$__all"
                ]
              },
              "datasource": {
                "type": "prometheus",
                "uid": "${datasource}"
              },
              "definition": "label_values(karpenter_nodes_allocatable, arch)",
              "hide": 0,
              "includeAll": true,
              "multi": true,
              "name": "arch",
              "options": [],
              "query": {
                "query": "label_values(karpenter_nodes_allocatable, arch)",
                "refId": "StandardVariableQuery"
              },
              "refresh": 2,
              "regex": "",
              "skipUrlSync": false,
              "sort": 1,
              "type": "query"
            },
            {
              "current": {
                "selected": true,
                "text": [
                  "All"
                ],
                "value": [
                  "$__all"
                ]
              },
              "datasource": {
                "type": "prometheus",
                "uid": "${datasource}"
              },
              "definition": "label_values(karpenter_nodes_allocatable, capacity_type)",
              "hide": 0,
              "includeAll": true,
              "multi": true,
              "name": "capacity_type",
              "options": [],
              "query": {
                "query": "label_values(karpenter_nodes_allocatable, capacity_type)",
                "refId": "StandardVariableQuery"
              },
              "refresh": 2,
              "regex": "",
              "skipUrlSync": false,
              "sort": 1,
              "type": "query"
            },
            {
              "current": {
                "selected": true,
                "text": [
                  "All"
                ],
                "value": [
                  "$__all"
                ]
              },
              "datasource": {
                "type": "prometheus",
                "uid": "${datasource}"
              },
              "definition": "label_values(karpenter_nodes_allocatable, instance_type)",
              "hide": 0,
              "includeAll": true,
              "multi": true,
              "name": "instance_type",
              "options": [],
              "query": {
                "query": "label_values(karpenter_nodes_allocatable, instance_type)",
                "refId": "StandardVariableQuery"
              },
              "refresh": 2,
              "regex": "",
              "skipUrlSync": false,
              "sort": 1,
              "type": "query"
            },
            {
              "current": {
                "selected": true,
                "text": "nodepool",
                "value": "nodepool"
              },
              "hide": 0,
              "includeAll": false,
              "multi": false,
              "name": "distribution_filter",
              "options": [
                {
                  "selected": false,
                  "text": "arch",
                  "value": "arch"
                },
                {
                  "selected": false,
                  "text": "capacity_type",
                  "value": "capacity_type"
                },
                {
                  "selected": false,
                  "text": "instance_type",
                  "value": "instance_type"
                },
                {
                  "selected": false,
                  "text": "namespace",
                  "value": "namespace"
                },
                {
                  "selected": false,
                  "text": "node",
                  "value": "node"
                },
                {
                  "selected": true,
                  "text": "nodepool",
                  "value": "nodepool"
                },
                {
                  "selected": false,
                  "text": "zone",
                  "value": "zone"
                }
              ],
              "query": "arch,capacity_type,instance_type,namespace,node,nodepool,zone",
              "queryValue": "",
              "skipUrlSync": false,
              "type": "custom"
            }
          ]
        },
        "time": {
          "from": "now-3h",
          "to": "now"
        },
        "timepicker": {},
        "timezone": "",
        "title": "Karpenter Capacity v1",
        "uid": "ta8I9Q67Z",
        "version": 7,
        "weekStart": ""
      }
`