package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagConfigurationsCmd = &cobra.Command{
	Use: "list-tag-configurations",

	Short: "Get a list of metrics",
	Long: `Get a list of metrics
Documentation: https://docs.datadoghq.com/api/latest/metrics/#list-tag-configurations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricsAndMetricTagConfigurationsResponse
		var err error

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err = api.ListTagConfigurations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-tag-configurations")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationsCmd)
}
