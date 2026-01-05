package metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTagConfigurationCmd = &cobra.Command{
	Use: "update-tag-configuration [metric_name]",

	Short: "Update a tag configuration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.UpdateTagConfiguration(client.NewContext(apiKey, appKey, site), args[0], datadogV2.MetricTagConfigurationUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-tag-configuration")

		cmd.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTagConfigurationCmd)
}
