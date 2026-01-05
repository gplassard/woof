package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTagConfigurationCmd = &cobra.Command{
	Use: "create-tag-configuration [metric_name]",

	Short: "Create a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MetricTagConfigurationCreateRequest{})
		cmdutil.HandleError(err, "failed to create-tag-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {
	Cmd.AddCommand(CreateTagConfigurationCmd)
}
