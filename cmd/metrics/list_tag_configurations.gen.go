package metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagConfigurationsCmd = &cobra.Command{
	Use: "list-tag-configurations",

	Short: "Get a list of metrics",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagConfigurations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-tag-configurations")

		cmd.Println(cmdutil.FormatJSON(res, "metrics"))
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationsCmd)
}
