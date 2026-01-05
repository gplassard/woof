package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListTagPipelinesRulesetsCmd = &cobra.Command{
	Use: "list-tag-pipelines-rulesets",

	Short: "List tag pipeline rulesets",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.ListTagPipelinesRulesets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-tag-pipelines-rulesets")

		cmd.Println(cmdutil.FormatJSON(res, "ruleset"))
	},
}

func init() {
	Cmd.AddCommand(ListTagPipelinesRulesetsCmd)
}
