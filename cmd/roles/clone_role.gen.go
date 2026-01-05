package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CloneRoleCmd = &cobra.Command{
	Use:     "clone-role [role_id]",
	Aliases: []string{"clone"},
	Short:   "Create a new role by cloning an existing role",
	Long: `Create a new role by cloning an existing role
Documentation: https://docs.datadoghq.com/api/latest/roles/#clone-role`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RoleResponse
		var err error

		var body datadogV2.RoleCloneRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.CloneRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to clone-role")

		cmd.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	CloneRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CloneRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CloneRoleCmd)
}
