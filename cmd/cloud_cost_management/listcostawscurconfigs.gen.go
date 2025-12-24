package cloud_cost_management

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCostAWSCURConfigsCmd = &cobra.Command{
	Use:   "listcostawscurconfigs",
	Short: "List Cloud Cost Management AWS CUR configs",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCostAWSCURConfigs(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listcostawscurconfigs: %v", err)
		}

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(ListCostAWSCURConfigsCmd)
}
