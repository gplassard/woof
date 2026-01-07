package security_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateResourceEvaluationFiltersCmd = &cobra.Command{
	Use: "update-resource-evaluation-filters",

	Short: "Update resource filters",
	Long: `Update resource filters
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-resource-evaluation-filters`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateResourceEvaluationFiltersResponse
		var err error

		var body datadogV2.UpdateResourceEvaluationFiltersRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateResourceEvaluationFilters(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-resource-evaluation-filters")

		fmt.Println(cmdutil.FormatJSON(res, "csm_resource_filter"))
	},
}

func init() {

	UpdateResourceEvaluationFiltersCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateResourceEvaluationFiltersCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateResourceEvaluationFiltersCmd)
}
