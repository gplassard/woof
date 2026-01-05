package cloud_cost_management

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderTagPipelinesRulesetsCmd = &cobra.Command{
	Use: "reorder-tag-pipelines-rulesets",

	Short: "Reorder tag pipeline rulesets",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err := api.ReorderTagPipelinesRulesets(client.NewContext(apiKey, appKey, site), datadogV2.ReorderRulesetResourceArray{})
		cmdutil.HandleError(err, "failed to reorder-tag-pipelines-rulesets")

	},
}

func init() {
	Cmd.AddCommand(ReorderTagPipelinesRulesetsCmd)
}
