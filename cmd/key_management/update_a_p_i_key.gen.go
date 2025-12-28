package key_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateAPIKeyCmd = &cobra.Command{
	Use:   "update_a_p_i_key [api_key_id]",
	Short: "Edit an API key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewKeyManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateAPIKey(client.NewContext(apiKey, appKey, site), args[0], datadogV2.APIKeyUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_a_p_i_key: %v", err)
		}

		cmdutil.PrintJSON(res, "api_keys")
	},
}

func init() {
	Cmd.AddCommand(UpdateAPIKeyCmd)
}
