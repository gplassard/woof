package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddPermissionToRoleCmd = &cobra.Command{
	Use:     "add-permission-to-role [role_id]",
	Aliases: []string{"add-permission-to"},
	Short:   "Grant permission to a role",
	Long: `Grant permission to a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#add-permission-to-role`,
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
		res, _, err = api.AddPermissionToRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-permission-to-role")

		fmt.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	AddPermissionToRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AddPermissionToRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AddPermissionToRoleCmd)
}
