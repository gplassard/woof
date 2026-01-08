package roles

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var AddUserToRoleCmd = &cobra.Command{
	Use:     "add-user-to-role [role_id]",
	Aliases: []string{"add-user-to"},
	Short:   "Add a user to a role",
	Long: `Add a user to a role
Documentation: https://docs.datadoghq.com/api/latest/roles/#add-user-to-role`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UsersResponse
		var err error

		var body datadogV2.RelationshipToUser
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewRolesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.AddUserToRole(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to add-user-to-role")

		fmt.Println(cmdutil.FormatJSON(res, "roles"))
	},
}

func init() {

	AddUserToRoleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	AddUserToRoleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(AddUserToRoleCmd)
}
