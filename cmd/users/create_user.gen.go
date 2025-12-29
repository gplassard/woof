package users

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateUserCmd = &cobra.Command{
	Use:   "create-user",
	Aliases: []string{ "create", },
	Short: "Create a user",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.CreateUser(client.NewContext(apiKey, appKey, site), datadogV2.UserCreateRequest{})
		cmdutil.HandleError(err, "failed to create-user")

		cmd.Println(cmdutil.FormatJSON(res, "users"))
	},
}

func init() {
	Cmd.AddCommand(CreateUserCmd)
}
