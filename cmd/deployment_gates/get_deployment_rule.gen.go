package deployment_gates

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDeploymentRuleCmd = &cobra.Command{
	Use:   "get-deployment-rule [gate_id] [id]",
	
	Short: "Get deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.GetDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-deployment-rule")

		cmdutil.PrintJSON(res, "deployment_rule")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentRuleCmd)
}
