package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDeploymentRuleCmd = &cobra.Command{
	Use: "get-deployment-rule [gate_id] [id]",

	Short: "Get deployment rule",
	Long: `Get deployment rule
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#get-deployment-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentRuleResponse
		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-deployment-rule")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetDeploymentRuleCmd)
}
