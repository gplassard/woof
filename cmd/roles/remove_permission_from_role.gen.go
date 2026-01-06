package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var RemovePermissionFromRoleCmd = &cobra.Command{
	Use:     "remove-permission-from-role [role_id]",
	Aliases: []string{"remove-permission-from"},
	Short:   "Revoke permission",
	Long: `Revoke permission
Documentation: https://docs.datadoghq.com/api/latest/roles/#remove-permission-from-role`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.PermissionsResponse
		var err error

		var body datadogV2.RelationshipToPermission
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.RemovePermissionFromRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to remove-permission-from-role")

		cmd.Println(cmdutil.FormatJSON(res, "permissions"))
	},
}

func init() {

	RemovePermissionFromRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	RemovePermissionFromRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(RemovePermissionFromRoleCmd)
}
