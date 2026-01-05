package metrics

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTagConfigurationCmd = &cobra.Command{
	Use: "create-tag-configuration [metric_name] [payload]",

	Short: "Create a tag configuration",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.MetricTagConfigurationCreateRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.CreateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-tag-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {
	Cmd.AddCommand(CreateTagConfigurationCmd)
}
