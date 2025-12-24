package deployment_gates

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDeploymentRuleCmd = &cobra.Command{
	Use:   "createdeploymentrule [gate_id]",
	Short: "Create deployment rule",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.CreateDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], datadogV2.CreateDeploymentRuleParams{})
		if err != nil {
			log.Fatalf("failed to createdeploymentrule: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gates")
	},
}

func init() {
	Cmd.AddCommand(CreateDeploymentRuleCmd)
}
