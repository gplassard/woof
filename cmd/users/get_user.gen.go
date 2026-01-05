package users

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetUserCmd = &cobra.Command{
	Use:     "get-user [user_id]",
	Aliases: []string{"get"},
	Short:   "Get user details",
	Long: `Get user details
Documentation: https://docs.datadoghq.com/api/latest/users/#get-user`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UserResponse
		var err error

		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err = api.GetUser(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-user")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {

	Cmd.AddCommand(GetUserCmd)
}
