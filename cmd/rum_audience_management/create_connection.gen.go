package rum_audience_management

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateConnectionCmd = &cobra.Command{
	Use:   "create-connection [entity]",
	
	Short: "Create connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		_, err := api.CreateConnection(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CreateConnectionRequest{})
		if err != nil {
			log.Fatalf("failed to create-connection: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(CreateConnectionCmd)
}
