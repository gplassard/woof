package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteCustomCostsFileCmd = &cobra.Command{
	Use:   "deletecustomcostsfile [file_id]",
	Short: "Delete Custom Costs file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.DeleteCustomCostsFile(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to deletecustomcostsfile: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteCustomCostsFileCmd)
}
