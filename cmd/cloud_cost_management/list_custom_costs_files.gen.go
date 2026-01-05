package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCustomCostsFilesCmd = &cobra.Command{
	Use: "list-custom-costs-files",

	Short: "List Custom Costs files",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListCustomCostsFiles(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-custom-costs-files")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {
	Cmd.AddCommand(ListCustomCostsFilesCmd)
}
