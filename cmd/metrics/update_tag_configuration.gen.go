package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateTagConfigurationCmd = &cobra.Command{
	Use: "update-tag-configuration [metric_name] [payload]",

	Short: "Update a tag configuration",
	Long: `Update a tag configuration
Documentation: https://docs.datadoghq.com/api/latest/metrics/#update-tag-configuration`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.MetricTagConfigurationResponse
		var err error

		var body datadogV2.MetricTagConfigurationUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err = api.UpdateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-tag-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTagConfigurationCmd)
}
