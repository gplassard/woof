package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateCurrentUserApplicationKeyCmd = &cobra.Command{
	Use:   "create-current-user-application-key",
	
	Short: "Create an application key for current user",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.CreateCurrentUserApplicationKey(client.NewContext(apiKey, appKey, site), datadogV2.ApplicationKeyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-current-user-application-key: %v", err)
		}

		cmdutil.PrintJSON(res, "application_keys")
	},
}

func init() {
	Cmd.AddCommand(CreateCurrentUserApplicationKeyCmd)
}
