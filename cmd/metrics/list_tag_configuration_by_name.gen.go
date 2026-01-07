package metrics

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagConfigurationByNameCmd = &cobra.Command{
	Use: "list-tag-configuration-by-name [metric_name]",

	Short: "List tag configuration by name",
	Long: `List tag configuration by name
Documentation: https://docs.datadoghq.com/api/latest/metrics/#list-tag-configuration-by-name`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricTagConfigurationResponse
		var err error

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListTagConfigurationByName(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-tag-configuration-by-name")

		fmt.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {

	Cmd.AddCommand(ListTagConfigurationByNameCmd)
}
