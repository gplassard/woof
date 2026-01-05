package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateResourceEvaluationFiltersCmd = &cobra.Command{
	Use: "update-resource-evaluation-filters [payload]",

	Short: "Update resource filters",
	Long: `Update resource filters
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#update-resource-evaluation-filters`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpdateResourceEvaluationFiltersResponse
		var err error

		var body datadogV2.UpdateResourceEvaluationFiltersRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err = api.UpdateResourceEvaluationFilters(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to update-resource-evaluation-filters")

		cmd.Println(cmdutil.FormatJSON(res, "csm_resource_filter"))
	},
}

func init() {
	Cmd.AddCommand(UpdateResourceEvaluationFiltersCmd)
}
