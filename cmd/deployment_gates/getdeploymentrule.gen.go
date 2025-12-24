package deployment_gates

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDeploymentRuleCmd = &cobra.Command{
	Use:   "getdeploymentrule [gate_id] [id]",
	Short: "Get deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.GetDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to getdeploymentrule: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gates")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentRuleCmd)
}
