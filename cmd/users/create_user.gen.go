package users

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a user",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		res, _, err := api.CreateUser(client.NewContext(apiKey, appKey, site), datadogV2.UserCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-user: %v", err)
		}

		cmdutil.PrintJSON(res, "users")
	},
}

func init() {
	Cmd.AddCommand(CreateUserCmd)
}
