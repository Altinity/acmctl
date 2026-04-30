package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "Manage cluster volumes and object storage",
}

var storageObjectAddCmd = &cobra.Command{
	Use:   "object-add <cluster-id>",
	Short: "Add object storage to a cluster",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"type": "type", "bucket": "bucket", "region": "region",
			"access-key": "accessKey", "secret-key": "secretKey",
			"create-bucket": "createBucket",
		})
		result, err := apiClient.AddObjectStorage(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var storageObjectDeleteCmd = &cobra.Command{
	Use:   "object-delete <cluster-id> <name>",
	Short: "Remove object storage from a cluster",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.RemoveObjectStorage(args[0], args[1]); err != nil {
			return err
		}
		fmt.Printf("Object storage %s removed from cluster %s.\n", args[1], args[0])
		return nil
	},
}

var storageVolumeListCmd = &cobra.Command{
	Use:   "volumes <cluster-id>",
	Short: "List cluster volumes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ListClusterVolumes(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var storageVolumeModifyCmd = &cobra.Command{
	Use:   "volumes-modify <cluster-id>",
	Short: "Modify cluster volumes (bulk)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"type": "type", "size": "size", "throughput": "throughput", "iops": "iops",
		})
		result, err := apiClient.ModifyClusterVolumes(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var storageVolumeUpdateCmd = &cobra.Command{
	Use:   "volume-update <volume-id>",
	Short: "Modify a single cluster volume",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		params := flagsToParams(cmd, map[string]string{
			"size": "size", "throughput": "throughput", "type": "type", "iops": "iops",
		})
		result, err := apiClient.UpdateClusterVolume(args[0], params)
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var storageVolumeDeleteCmd = &cobra.Command{
	Use:   "volume-delete <volume-id>",
	Short: "Remove a cluster volume",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.DeleteClusterVolume(args[0]); err != nil {
			return err
		}
		fmt.Printf("Volume %s deleted.\n", args[0])
		return nil
	},
}

var storageVolumeCordonCmd = &cobra.Command{
	Use:   "volume-cordon <volume-id>",
	Short: "Set cordon value for a cluster volume",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cordon, _ := cmd.Flags().GetString("cordon")
		if err := apiClient.CordonClusterVolume(args[0], cordon); err != nil {
			return err
		}
		fmt.Printf("Volume %s cordon set.\n", args[0])
		return nil
	},
}

var storageVolumeFreeCmd = &cobra.Command{
	Use:   "volume-free <volume-id>",
	Short: "Free up a cluster volume",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.FreeClusterVolume(args[0]); err != nil {
			return err
		}
		fmt.Printf("Volume %s free initiated.\n", args[0])
		return nil
	},
}

var storageVolumeValidateCmd = &cobra.Command{
	Use:   "volume-validate <volume-id>",
	Short: "Validate volume modification based on PVC status",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := apiClient.ValidateClusterVolumeModification(args[0])
		if err != nil {
			return err
		}
		return printJSON(result)
	},
}

var storageInterruptFreeCmd = &cobra.Command{
	Use:   "interrupt-free <cluster-id>",
	Short: "Interrupt running free-volume queries",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := apiClient.InterruptFreeVolumes(args[0]); err != nil {
			return err
		}
		fmt.Printf("Free-volume queries interrupted for cluster %s.\n", args[0])
		return nil
	},
}

func init() {
	for _, name := range []string{"type", "bucket", "region", "access-key", "secret-key", "create-bucket"} {
		storageObjectAddCmd.Flags().String(name, "", "")
	}

	for _, name := range []string{"type", "size", "throughput", "iops"} {
		storageVolumeModifyCmd.Flags().String(name, "", "")
		storageVolumeUpdateCmd.Flags().String(name, "", "")
	}

	storageVolumeCordonCmd.Flags().String("cordon", "", "cordon value")

	for _, c := range []*cobra.Command{
		storageObjectAddCmd, storageObjectDeleteCmd, storageVolumeListCmd, storageVolumeModifyCmd,
		storageVolumeUpdateCmd, storageVolumeDeleteCmd, storageVolumeCordonCmd, storageVolumeFreeCmd,
		storageVolumeValidateCmd, storageInterruptFreeCmd,
	} {
		storageCmd.AddCommand(c)
	}
	rootCmd.AddCommand(storageCmd)
}
