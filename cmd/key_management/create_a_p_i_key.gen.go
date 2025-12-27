package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateAPIKeyCmd = &cobra.Command{
	Use:   "create_a_p_i_key",
	Short: "Create an API key",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.CreateAPIKey(client.NewContext(apiKey, appKey, site), datadogV2.APIKeyCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_a_p_i_key: %v", err)
		}

		cmdutil.PrintJSON(res, "key_management")
	},
}

func init() {
	Cmd.AddCommand(CreateAPIKeyCmd)
}
