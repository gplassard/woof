package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDeploymentRuleCmd = &cobra.Command{
	Use: "create-deployment-rule [gate_id]",

	Short: "Create deployment rule",
	Long: `Create deployment rule
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#create-deployment-rule`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentRuleResponse
		var err error

		var body datadogV2.CreateDeploymentRuleParams
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err = api.CreateDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-deployment-rule")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_rule"))
	},
}

func init() {

	CreateDeploymentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDeploymentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDeploymentRuleCmd)
}
