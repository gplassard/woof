package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetResourceEvaluationFiltersCmd = &cobra.Command{
	Use: "get-resource-evaluation-filters",

	Short: "List resource filters",
	Long: `List resource filters
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#get-resource-evaluation-filters`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetResourceEvaluationFiltersResponse
		var err error

		optionalParams := datadogV2.NewGetResourceEvaluationFiltersOptionalParameters()

		if cmd.Flags().Changed("cloud-provider") {
			val, _ := cmd.Flags().GetString("cloud-provider")
			optionalParams.WithCloudProvider(val)
		}

		if cmd.Flags().Changed("account-id") {
			val, _ := cmd.Flags().GetString("account-id")
			optionalParams.WithAccountId(val)
		}

		if cmd.Flags().Changed("skip-cache") {
			val, _ := cmd.Flags().GetString("skip-cache")
			optionalParams.WithSkipCache(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetResourceEvaluationFilters(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to get-resource-evaluation-filters")

		fmt.Println(cmdutil.FormatJSON(res, "csm_resource_filter"))
	},
}

func init() {

	GetResourceEvaluationFiltersCmd.Flags().String("cloud-provider", "", "Filter resource filters by cloud provider (e.g. aws, gcp, azure).")

	GetResourceEvaluationFiltersCmd.Flags().String("account-id", "", "Filter resource filters by cloud provider account ID. This parameter is only valid when provider is specified.")

	GetResourceEvaluationFiltersCmd.Flags().String("skip-cache", "", "Skip cache for resource filters.")

	Cmd.AddCommand(GetResourceEvaluationFiltersCmd)
}
