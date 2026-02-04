package users

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListUserOrganizationsCmd = &cobra.Command{
	Use:     "list-user-organizations [user_id]",
	Aliases: []string{"list-organizations"},
	Short:   "Get a user organization",
	Long: `Get a user organization
Documentation: https://docs.datadoghq.com/api/latest/users/#list-user-organizations`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserResponse
		var err error

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListUserOrganizations(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-user-organizations")

		fmt.Println(cmdutil.FormatJSON(res, "user_organization"))
	},
}

func init() {

	Cmd.AddCommand(ListUserOrganizationsCmd)
}
