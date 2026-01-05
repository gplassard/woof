package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateTagPipelinesRulesetCmd = &cobra.Command{
	Use: "create-tag-pipelines-ruleset [payload]",

	Short: "Create tag pipeline ruleset",
	Long: `Create tag pipeline ruleset
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#create-tag-pipelines-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RulesetResp
		var err error

		var body datadogV2.CreateRulesetRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err = api.CreateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-tag-pipelines-ruleset")

		cmd.Println(cmdutil.FormatJSON(res, "ruleset"))
	},
}

func init() {
	Cmd.AddCommand(CreateTagPipelinesRulesetCmd)
}
