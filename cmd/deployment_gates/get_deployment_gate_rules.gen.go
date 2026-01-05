package deployment_gates

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDeploymentGateRulesCmd = &cobra.Command{
	Use:     "get-deployment-gate-rules [gate_id]",
	Aliases: []string{"get-rules"},
	Short:   "Get rules for a deployment gate",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.GetDeploymentGateRules(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-deployment-gate-rules")

		cmd.Println(cmdutil.FormatJSON(res, "list_deployment_rules"))
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentGateRulesCmd)
}
