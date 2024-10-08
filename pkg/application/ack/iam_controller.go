package ack

import (
	"github.com/awslabs/eksdemo/pkg/application"
	"github.com/awslabs/eksdemo/pkg/cmd"
	"github.com/awslabs/eksdemo/pkg/installer"
	"github.com/awslabs/eksdemo/pkg/resource"
	"github.com/awslabs/eksdemo/pkg/resource/irsa"
	"github.com/awslabs/eksdemo/pkg/template"
)

// Docs:    https://aws-controllers-k8s.github.io/community/docs/community/overview/
// Docs:    https://aws-controllers-k8s.github.io/community/reference/
// GitHub:  https://github.com/aws-controllers-k8s/iam-controller
// Helm:    https://github.com/aws-controllers-k8s/iam-controller/tree/main/helm
// Chart:   https://gallery.ecr.aws/aws-controllers-k8s/iam-chart
// Repo:    https://gallery.ecr.aws/aws-controllers-k8s/iam-controller
// Version: Latest is v1.3.12 (as of 9/5/24)

func NewIAMController() *application.Application {
	return &application.Application{
		Command: cmd.Command{
			Parent:      "ack",
			Name:        "iam-controller",
			Description: "ACK IAM Controller",
			Aliases:     []string{"iam"},
		},

		Dependencies: []*resource.Resource{
			irsa.NewResourceWithOptions(&irsa.IrsaOptions{
				CommonOptions: resource.CommonOptions{
					Name: "ack-iam-controller-irsa",
				},
				PolicyType: irsa.PolicyDocument,
				PolicyDocTemplate: &template.TextTemplate{
					Template: iamPolicyDocTemplate,
				},
			}),
		},

		Options: &application.ApplicationOptions{
			Namespace:      "ack-system",
			ServiceAccount: "ack-iam-controller",
			DefaultVersion: &application.LatestPrevious{
				LatestChart:   "1.3.12",
				Latest:        "1.3.12",
				PreviousChart: "1.2.1",
				Previous:      "1.2.1",
			},
		},

		Installer: &installer.HelmInstaller{
			ReleaseName:   "ack-iam-controller",
			RepositoryURL: "oci://public.ecr.aws/aws-controllers-k8s/iam-chart",
			ValuesTemplate: &template.TextTemplate{
				Template: iamValuesTemplate,
			},
		},
	}
}

// https://github.com/aws-controllers-k8s/iam-controller/blob/main/config/iam/recommended-inline-policy
const iamPolicyDocTemplate = `
Version: '2012-10-17'
Statement:
- Effect: Allow
  Action:
  - iam:GetGroup
  - iam:CreateGroup
  - iam:DeleteGroup
  - iam:UpdateGroup
  - iam:GetRole
  - iam:CreateRole
  - iam:DeleteRole
  - iam:UpdateRole
  - iam:PutRolePermissionsBoundary
  - iam:PutUserPermissionsBoundary
  - iam:GetUser
  - iam:CreateUser
  - iam:DeleteUser
  - iam:UpdateUser
  - iam:GetPolicy
  - iam:CreatePolicy
  - iam:DeletePolicy
  - iam:GetPolicyVersion
  - iam:CreatePolicyVersion
  - iam:DeletePolicyVersion
  - iam:ListPolicyVersions
  - iam:ListPolicyTags
  - iam:ListAttachedGroupPolicies
  - iam:GetGroupPolicy
  - iam:PutGroupPolicy
  - iam:AttachGroupPolicy
  - iam:DetachGroupPolicy
  - iam:DeleteGroupPolicy
  - iam:ListAttachedRolePolicies
  - iam:ListRolePolicies
  - iam:GetRolePolicy
  - iam:PutRolePolicy
  - iam:AttachRolePolicy
  - iam:DetachRolePolicy
  - iam:DeleteRolePolicy
  - iam:ListAttachedUserPolicies
  - iam:ListUserPolicies
  - iam:GetUserPolicy
  - iam:PutUserPolicy
  - iam:AttachUserPolicy
  - iam:DetachUserPolicy
  - iam:DeleteUserPolicy
  - iam:ListRoleTags
  - iam:ListUserTags
  - iam:TagPolicy
  - iam:UntagPolicy
  - iam:TagRole
  - iam:UntagRole
  - iam:TagUser
  - iam:UntagUser
  - iam:RemoveClientIDFromOpenIDConnectProvider
  - iam:ListOpenIDConnectProviderTags
  - iam:UpdateOpenIDConnectProviderThumbprint
  - iam:UntagOpenIDConnectProvider
  - iam:AddClientIDToOpenIDConnectProvider
  - iam:DeleteOpenIDConnectProvider
  - iam:GetOpenIDConnectProvider
  - iam:TagOpenIDConnectProvider
  - iam:CreateOpenIDConnectProvider
  - iam:UpdateAssumeRolePolicy
  Resource: "*"
`

// https://github.com/aws-controllers-k8s/iam-controller/blob/main/helm/values.yaml
const iamValuesTemplate = `---
image:
  tag: {{ .Version }}
fullnameOverride: ack-iam-controller
aws:
  region: {{ .Region }}
serviceAccount:
  name: {{ .ServiceAccount }}
  annotations:
    {{ .IrsaAnnotation }}
`
