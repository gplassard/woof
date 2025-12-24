package rum_audience_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateConnectionCmd = &cobra.Command{
	Use:   "updateconnection [entity]",
	Short: "Update connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		_, err := api.UpdateConnection(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateConnectionRequest{})
		if err != nil {
			log.Fatalf("failed to updateconnection: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UpdateConnectionCmd)
}
