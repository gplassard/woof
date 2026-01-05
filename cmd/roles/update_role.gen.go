package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateRoleCmd = &cobra.Command{
	Use:     "update-role [role_id]",
	Aliases: []string{"update"},
	Short:   "Update a role",
	Long: `Update a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#update-role`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleUpdateResponse
		var err error

		var body datadogV2.RoleUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.UpdateRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	UpdateRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateRoleCmd)
}
