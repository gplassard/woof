package roles

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var AddUserToRoleCmd = &cobra.Command{
	Use:     "add-user-to-role [role_id] [payload]",
	Aliases: []string{"add-user-to"},
	Short:   "Add a user to a role",
	Long: `Add a user to a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#add-user-to-role`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsersResponse
		var err error

		var body datadogV2.RelationshipToUser
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		res, _, err = api.AddUserToRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-user-to-role")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(AddUserToRoleCmd)
}
