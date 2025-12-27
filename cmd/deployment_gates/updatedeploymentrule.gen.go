package deployment_gates

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDeploymentRuleCmd = &cobra.Command{
	Use:   "updatedeploymentrule [gate_id] [id]",
	Short: "Update deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.UpdateDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.UpdateDeploymentRuleParams{})
		if err != nil {
			log.Fatalf("failed to updatedeploymentrule: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gates")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeploymentRuleCmd)
}
