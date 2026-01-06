package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateTagPipelinesRulesetCmd = &cobra.Command{
	Use: "create-tag-pipelines-ruleset",

	Short: "Create tag pipeline ruleset",
	Long: `Create tag pipeline ruleset
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-tag-pipelines-ruleset`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RulesetResp
		var err error

		var body datadogV2.CreateRulesetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-tag-pipelines-ruleset")

		cmd.Println(cmdutil.FormatJSON(res, "ruleset"))
	},
}

func init() {

	CreateTagPipelinesRulesetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateTagPipelinesRulesetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateTagPipelinesRulesetCmd)
}
