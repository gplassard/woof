package action_connection

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteActionConnectionCmd = &cobra.Command{
	Use:   "deleteactionconnection [connection_id]",
	Short: "Delete an existing Action Connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewActionConnectionApi(client.NewAPIClient())
		_, err := api.DeleteActionConnection(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deleteactionconnection: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteActionConnectionCmd)
}
