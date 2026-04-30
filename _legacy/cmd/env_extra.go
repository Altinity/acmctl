package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var envApproveCmd = &cobra.Command{
	Use:   "approve <id>",
	Short: "Approve an environment setup request",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		reason, _ := cmd.Flags().GetString("reason")
		if err := apiClient.ApproveEnvironment(args[0], reason); err != nil {
			return err
		}
		fmt.Printf("Environment %s approved.\n", args[0])
		return nil
	},
}

var envAccCheckCmd = &cobra.Command{
	Use:   "acc-check <id>",
	Short: "Check Cloud Connector ↔ Cloud Controller connection",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		noWait, _ := cmd.Flags().GetBool("no-wait")
		result, err := apiClient.AccCheck(args[0], noWait)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envAccConnectCmd = &cobra.Command{
	Use:   "acc-connect <id>",
	Short: "Set environment into awaiting-connection state",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resources, _ := cmd.Flags().GetString("resources")
		result, err := apiClient.AccConnect(args[0], resources)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envAccTokenCmd = &cobra.Command{
	Use:   "acc-token <id>",
	Short: "Get a Cloud Connector connection token",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetAccToken(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envResetCmd = &cobra.Command{
	Use:   "reset <id>",
	Short: "Reset an environment to its initial state",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.ResetEnvironment(args[0]); err != nil {
			return err
		}
		fmt.Printf("Environment %s reset.\n", args[0])
		return nil
	},
}

var envKubeUpdateCmd = &cobra.Command{
	Use:   "kube-update <id>",
	Short: "Refresh Kubernetes environment details from acm-env-details ConfigMap",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.RefreshKubeEnvironment(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envAlertsCmd = &cobra.Command{
	Use:   "alerts <id>",
	Short: "List alerts for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		resolved, _ := cmd.Flags().GetString("resolved")
		result, err := apiClient.ListEnvironmentAlerts(args[0], resolved)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envBucketsCmd = &cobra.Command{
	Use:   "buckets <id>",
	Short: "List buckets in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		bucketType, _ := cmd.Flags().GetString("type")
		result, err := apiClient.ListEnvironmentBuckets(args[0], bucketType)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envClusterLaunchValidityCmd = &cobra.Command{
	Use:   "launch-validity <id>",
	Short: "Check whether cluster launch is valid",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetClusterLaunchValidity(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envExportCmd = &cobra.Command{
	Use:   "export <id>",
	Short: "Export environment specification (JSON or YAML)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		format, _ := cmd.Flags().GetString("format")
		result, err := apiClient.ExportEnvironment(args[0], format)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envLogCmd = &cobra.Command{
	Use:   "log <id>",
	Short: "Show environment action log",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"source": "source", "page": "page", "limit": "limit",
			"filter": "filter", "order": "order",
		})
		result, err := apiClient.GetEnvironmentLog(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envIcebergCmd = &cobra.Command{
	Use:   "iceberg <id>",
	Short: "Get Iceberg catalog settings for an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetIcebergSettings(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envInviteDetailsCmd = &cobra.Command{
	Use:   "invite-details <id>",
	Short: "Show default org assignment for an invited user",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetInviteDetails(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envInviteCmd = &cobra.Command{
	Use:   "invite <id>",
	Short: "Invite a user to an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"email": "email", "role": "id_role"})
		result, err := apiClient.InviteUser(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envResourcesCmd = &cobra.Command{
	Use:   "resources <id>",
	Short: "List environment resources",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		limits, _ := cmd.Flags().GetString("limits")
		result, err := apiClient.GetEnvironmentResources(args[0], limits)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envResourceCmd = &cobra.Command{
	Use:   "resource <id>",
	Short: "Get a kubernetes resource spec",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"kind": "kind", "name": "name", "api-version": "apiVersion",
		})
		result, err := apiClient.GetEnvironmentResource(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envUsageCmd = &cobra.Command{
	Use:   "usage <id>",
	Short: "Show environment resource usage",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		skipID, _ := cmd.Flags().GetString("skip-cluster-id")
		result, err := apiClient.GetEnvironmentUsage(args[0], skipID)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

// CHOP / Kubernetes management

var envChopListCmd = &cobra.Command{
	Use:   "chop-list <id>",
	Short: "List CHOP configurations",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListChopConfigurations(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envChopAddCmd = &cobra.Command{
	Use:   "chop-add <id>",
	Short: "Add a CHOP configuration",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		spec, _ := cmd.Flags().GetString("spec")
		resolvedSpec, err := resolveValue(spec)
		if err != nil {
			return err
		}
		result, err := apiClient.AddChopConfiguration(args[0], map[string]string{
			"name": name, "spec": resolvedSpec,
		})
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envChopPatchCmd = &cobra.Command{
	Use:   "chop-patch <id> <name>",
	Short: "Patch a CHOP configuration",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, _ := cmd.Flags().GetString("spec")
		resolvedSpec, err := resolveValue(spec)
		if err != nil {
			return err
		}
		result, err := apiClient.PatchChopConfiguration(args[0], args[1], map[string]string{"spec": resolvedSpec})
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envChopDeleteCmd = &cobra.Command{
	Use:   "chop-delete <id> <name>",
	Short: "Delete a CHOP configuration",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteChopConfiguration(args[0], args[1]); err != nil {
			return err
		}
		fmt.Printf("CHOP configuration %s/%s deleted.\n", args[0], args[1])
		return nil
	},
}

var envChopApplyCmd = &cobra.Command{
	Use:   "chop-apply <id>",
	Short: "Apply all CHOP configuration changes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.ApplyChopConfiguration(args[0]); err != nil {
			return err
		}
		fmt.Printf("CHOP configurations applied for %s.\n", args[0])
		return nil
	},
}

var envConfigApplyCmd = &cobra.Command{
	Use:   "config-apply <id>",
	Short: "Apply CHOP/ConfigMap/Template changes by restarting pods",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ApplyKubeConfig(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envConfigTemplateListCmd = &cobra.Command{
	Use:   "config-templates <id>",
	Short: "List Kubernetes configuration templates",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListConfigurationTemplates(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envConfigTemplatePatchCmd = &cobra.Command{
	Use:   "config-template-patch <id> <name>",
	Short: "Patch a Kubernetes configuration template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		spec, _ := cmd.Flags().GetString("spec")
		resolved, err := resolveValue(spec)
		if err != nil {
			return err
		}
		result, err := apiClient.PatchConfigurationTemplate(args[0], args[1], map[string]string{"spec": resolved})
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envKubeConfigmapsListCmd = &cobra.Command{
	Use:   "configmaps <id>",
	Short: "List Kubernetes ConfigMaps in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListKubeConfigmaps(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envKubeConfigmapPatchCmd = &cobra.Command{
	Use:   "configmap-patch <id> <name>",
	Short: "Patch a Kubernetes ConfigMap",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, _ := cmd.Flags().GetString("data")
		resolved, err := resolveValue(data)
		if err != nil {
			return err
		}
		result, err := apiClient.PatchKubeConfigmap(args[0], args[1], map[string]string{"data": resolved})
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envKubeMapCmd = &cobra.Command{
	Use:   "kube-map <id>",
	Short: "Show Kubernetes resources map",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.GetKubeMap(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envKubeChoCmd = &cobra.Command{
	Use:   "kube-cho <id>",
	Short: "Install / remove ClickHouse Operator inside the environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{"action": "action", "version": "version"})
		result, err := apiClient.HandleKubeCho(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envDiscoverCmd = &cobra.Command{
	Use:   "discover <id>",
	Short: "Discover Kubernetes clusters in an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.DiscoverEnvironment(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envDiscoverConfirmCmd = &cobra.Command{
	Use:   "discover-confirm <id>",
	Short: "Confirm which discovered Kubernetes clusters to check out",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		clusters, _ := cmd.Flags().GetString("clusters")
		result, err := apiClient.ConfirmDiscovery(args[0], clusters)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

// Backups (external bucket inspection)

var envBackupsCmd = &cobra.Command{
	Use:   "external-backups",
	Short: "List backup entries from an external bucket (no environment context)",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"type": "type", "provider": "provider",
			"access-key": "accessKey", "secret-key": "secretKey",
			"arn": "arn", "endpoint": "endpoint", "region": "region",
			"bucket": "bucket", "path": "path", "check": "check",
		})
		result, err := apiClient.ListExternalBackups(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envEnvBackupsCmd = &cobra.Command{
	Use:   "env-backups <id>",
	Short: "List backup entries from an environment-scoped external bucket",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"type": "type", "path": "path", "check": "check", "schedule": "schedule",
		})
		result, err := apiClient.ListEnvironmentBackups(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

// Provisioning

var envConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect an existing environment to ACM",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"name": "name", "type": "type", "host": "host", "port": "port",
			"user": "user", "pass": "pass",
			"ssh-host": "sshHost", "ssh-port": "sshPort",
			"ssh-user": "sshUser", "ssh-pass": "sshPass",
			"kube-api-url": "kubeAPIUrl", "kube-token": "kubeToken",
			"kube-auth-options": "kubeAuthOptions",
			"kube-namespace": "kubeNamespace",
			"kube-namespace-manage": "kubeNamespaceManage",
		})
		result, err := apiClient.ConnectEnvironment(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import an environment from a JSON source",
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.ImportEnvironment(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Request provisioning of a new environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"name": "name", "cloud-provider": "cloud_provider",
			"aws-region": "aws_region", "gcp-region": "gcp_region",
			"azure-region": "azure_region", "hcloud-region": "hcloud_region",
			"first": "first",
		})
		result, err := apiClient.RequestEnvironment(params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

// Cluster CRUD via environment

var envClusterAddCmd = &cobra.Command{
	Use:   "cluster-add <env-id>",
	Short: "Add a cluster descriptor to an environment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.AddClusterToEnvironment(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envClusterImportCmd = &cobra.Command{
	Use:   "cluster-import <env-id>",
	Short: "Import a cluster from a JSON source",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.ImportCluster(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var envClusterRestoreCmd = &cobra.Command{
	Use:   "cluster-restore <env-id>",
	Short: "Restore a cluster from external storage",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params, err := collectFieldFlags(cmd)
		if err != nil {
			return err
		}
		result, err := apiClient.RestoreClusterIntoEnvironment(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

func init() {
	envApproveCmd.Flags().String("reason", "", "approval reason")

	envAccCheckCmd.Flags().Bool("no-wait", false, "skip wait")

	envAccConnectCmd.Flags().String("resources", "", "resources spec")

	envAlertsCmd.Flags().String("resolved", "", "filter resolved")

	envBucketsCmd.Flags().String("type", "", "bucket type")

	envExportCmd.Flags().String("format", "", "json or yaml")

	for _, name := range []string{"source", "page", "limit", "filter", "order"} {
		envLogCmd.Flags().String(name, "", "")
	}

	envInviteCmd.Flags().String("email", "", "email")
	envInviteCmd.Flags().String("role", "", "role ID")

	envResourcesCmd.Flags().String("limits", "", "include limits")
	for _, name := range []string{"kind", "name", "api-version"} {
		envResourceCmd.Flags().String(name, "", "")
	}
	envUsageCmd.Flags().String("skip-cluster-id", "", "skip cluster ID")

	envChopAddCmd.Flags().String("name", "", "config name")
	envChopAddCmd.Flags().String("spec", "", "spec (use @file or @- for stdin)")
	envChopPatchCmd.Flags().String("spec", "", "spec (use @file or @- for stdin)")

	envConfigTemplatePatchCmd.Flags().String("spec", "", "spec (use @file or @- for stdin)")
	envKubeConfigmapPatchCmd.Flags().String("data", "", "data (use @file or @- for stdin)")

	envKubeChoCmd.Flags().String("action", "", "install/remove")
	envKubeChoCmd.Flags().String("version", "", "operator version")

	envDiscoverConfirmCmd.Flags().String("clusters", "", "clusters spec")

	for _, name := range []string{"type", "provider", "access-key", "secret-key",
		"arn", "endpoint", "region", "bucket", "path", "check"} {
		envBackupsCmd.Flags().String(name, "", "")
	}
	for _, name := range []string{"type", "path", "check", "schedule"} {
		envEnvBackupsCmd.Flags().String(name, "", "")
	}

	for _, name := range []string{"name", "type", "host", "port", "user", "pass",
		"ssh-host", "ssh-port", "ssh-user", "ssh-pass",
		"kube-api-url", "kube-token", "kube-auth-options",
		"kube-namespace", "kube-namespace-manage"} {
		envConnectCmd.Flags().String(name, "", "")
	}

	envImportCmd.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")

	for _, name := range []string{"name", "cloud-provider", "aws-region", "gcp-region",
		"azure-region", "hcloud-region", "first"} {
		envRequestCmd.Flags().String(name, "", "")
	}

	for _, c := range []*cobra.Command{envClusterAddCmd, envClusterImportCmd, envClusterRestoreCmd} {
		c.Flags().StringSliceP("field", "F", nil, "key=value (repeatable)")
	}

	for _, c := range []*cobra.Command{
		envApproveCmd, envAccCheckCmd, envAccConnectCmd, envAccTokenCmd,
		envResetCmd, envKubeUpdateCmd,
		envAlertsCmd, envBucketsCmd, envClusterLaunchValidityCmd,
		envExportCmd, envLogCmd, envIcebergCmd, envInviteDetailsCmd, envInviteCmd,
		envResourcesCmd, envResourceCmd, envUsageCmd,
		envChopListCmd, envChopAddCmd, envChopPatchCmd, envChopDeleteCmd, envChopApplyCmd,
		envConfigApplyCmd, envConfigTemplateListCmd, envConfigTemplatePatchCmd,
		envKubeConfigmapsListCmd, envKubeConfigmapPatchCmd, envKubeMapCmd, envKubeChoCmd,
		envDiscoverCmd, envDiscoverConfirmCmd,
		envBackupsCmd, envEnvBackupsCmd,
		envConnectCmd, envImportCmd, envRequestCmd,
		envClusterAddCmd, envClusterImportCmd, envClusterRestoreCmd,
	} {
		envCmd.AddCommand(c)
	}
}
