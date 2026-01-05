package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDeploymentRuleCmd = &cobra.Command{
	Use: "create-deployment-rule [gate_id] [payload]",

	Short: "Create deployment rule",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.CreateDeploymentRuleParams
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.CreateDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-deployment-rule")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_rule"))
	},
}

func init() {
	Cmd.AddCommand(CreateDeploymentRuleCmd)
}
