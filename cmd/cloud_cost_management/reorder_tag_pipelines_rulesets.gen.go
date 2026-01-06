package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ReorderTagPipelinesRulesetsCmd = &cobra.Command{
	Use: "reorder-tag-pipelines-rulesets",

	Short: "Reorder tag pipeline rulesets",
	Long: `Reorder tag pipeline rulesets
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#reorder-tag-pipelines-rulesets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.ReorderRulesetResourceArray
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.ReorderTagPipelinesRulesets(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to reorder-tag-pipelines-rulesets")

	},
}

func init() {

	ReorderTagPipelinesRulesetsCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	ReorderTagPipelinesRulesetsCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(ReorderTagPipelinesRulesetsCmd)
}
