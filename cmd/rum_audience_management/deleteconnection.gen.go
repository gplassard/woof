package rum_audience_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteConnectionCmd = &cobra.Command{
	Use:   "deleteconnection [id] [entity]",
	Short: "Delete connection",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		_, err := api.DeleteConnection(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to deleteconnection: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteConnectionCmd)
}
