package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateTagPipelinesRulesetCmd = &cobra.Command{
	Use: "update-tag-pipelines-ruleset [ruleset_id] [payload]",

	Short: "Update tag pipeline ruleset",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.UpdateRulesetRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		res, _, err := api.UpdateTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-tag-pipelines-ruleset")

		cmd.Println(cmdutil.FormatJSON(res, "ruleset"))
	},
}

func init() {
	Cmd.AddCommand(UpdateTagPipelinesRulesetCmd)
}
