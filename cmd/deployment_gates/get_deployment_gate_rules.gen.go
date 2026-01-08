package deployment_gates

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDeploymentGateRulesCmd = &cobra.Command{
	Use:     "get-deployment-gate-rules [gate_id]",
	Aliases: []string{"get-rules"},
	Short:   "Get rules for a deployment gate",
	Long: `Get rules for a deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#get-deployment-gate-rules`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGateRulesResponse
		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDeploymentGateRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-deployment-gate-rules")

		fmt.Println(cmdutil.FormatJSON(res, "deployment_gate_rule"))
	},
}

func init() {

	Cmd.AddCommand(GetDeploymentGateRulesCmd)
}
