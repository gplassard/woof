package cloud_cost_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTagPipelinesRulesetCmd = &cobra.Command{
	Use: "delete-tag-pipelines-ruleset [ruleset_id]",

	Short: "Delete tag pipeline ruleset",
	Long: `Delete tag pipeline ruleset
Documentation: https://docs.datadoghq.com/api/latest/cloud-cost-management/#delete-tag-pipelines-ruleset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCloudCostManagementApi(client.NewAPIClient())
		_, err = api.DeleteTagPipelinesRuleset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-tag-pipelines-ruleset")

	},
}

func init() {
	Cmd.AddCommand(DeleteTagPipelinesRulesetCmd)
}
