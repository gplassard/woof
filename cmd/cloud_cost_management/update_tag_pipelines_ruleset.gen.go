package cloud_cost_management

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateTagPipelinesRulesetCmd = &cobra.Command{
	Use: "update-tag-pipelines-ruleset [ruleset_id]",

	Short: "Update tag pipeline ruleset",
	Long: `Update tag pipeline ruleset
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#update-tag-pipelines-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RulesetResp
		var err error

		var body datadogV2.UpdateRulesetRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-tag-pipelines-ruleset")

		fmt.Println(cmdutil.FormatJSON(res, "ruleset"))
	},
}

func init() {

	UpdateTagPipelinesRulesetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateTagPipelinesRulesetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateTagPipelinesRulesetCmd)
}
