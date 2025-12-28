package cloud_cost_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCustomCostsFilesCmd = &cobra.Command{
	Use:   "list-custom-costs-files",
	
	Short: "List Custom Costs files",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCustomCostsFiles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-custom-costs-files")

		cmdutil.PrintJSON(res, "cloud_cost_management")
	},
}

func init() {
	Cmd.AddCommand(ListCustomCostsFilesCmd)
}
