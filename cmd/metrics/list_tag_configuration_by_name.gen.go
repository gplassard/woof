package metrics

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagConfigurationByNameCmd = &cobra.Command{
	Use: "list-tag-configuration-by-name [metric_name]",

	Short: "List tag configuration by name",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewMetricsApi(client.NewAPIClient())
		res, _, err := api.ListTagConfigurationByName(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-tag-configuration-by-name")

		cmd.Println(cmdutil.FormatJSON(res, "manage_tags"))
	},
}

func init() {
	Cmd.AddCommand(ListTagConfigurationByNameCmd)
}
