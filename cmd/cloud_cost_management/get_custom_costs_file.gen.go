package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetCustomCostsFileCmd = &cobra.Command{
	Use: "get-custom-costs-file [file_id]",

	Short: "Get Custom Costs file",
	Long: `Get Custom Costs file
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#get-custom-costs-file`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomCostsFileGetResponse
		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetCustomCostsFile(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-custom-costs-file")

		fmt.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	Cmd.AddCommand(GetCustomCostsFileCmd)
}
