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
	Long: `List Custom Costs files
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#list-custom-costs-files`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CustomCostsFileListResponse
		var err error

		optionalParams := datadogV2.NewListCustomCostsFilesOptionalParameters()

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("page-size") {
			val, _ := cmd.Flags().GetInt64("page-size")
			optionalParams.WithPageSize(val)
		}

		if cmd.Flags().Changed("filter-status") {
			val, _ := cmd.Flags().GetString("filter-status")
			optionalParams.WithFilterStatus(val)
		}

		if cmd.Flags().Changed("sort") {
			val, _ := cmd.Flags().GetString("sort")
			optionalParams.WithSort(val)
		}

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCustomCostsFiles(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-custom-costs-files")

		cmd.Println(cmdutil.FormatJSON(res, "cloud_cost_management"))
	},
}

func init() {

	ListCustomCostsFilesCmd.Flags().Int64("page-number", 0, "Page number for pagination")

	ListCustomCostsFilesCmd.Flags().Int64("page-size", 0, "Page size for pagination")

	ListCustomCostsFilesCmd.Flags().String("filter-status", "", "Filter by file status")

	ListCustomCostsFilesCmd.Flags().String("sort", "", "Sort key with optional descending prefix")

	Cmd.AddCommand(ListCustomCostsFilesCmd)
}
