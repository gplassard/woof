package users

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DisableUserCmd = &cobra.Command{
	Use:   "disableuser [user_id]",
	Short: "Disable a user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewUsersApi(client.NewAPIClient())
		_, err := api.DisableUser(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to disableuser: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DisableUserCmd)
}
