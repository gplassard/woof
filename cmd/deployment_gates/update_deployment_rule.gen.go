package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDeploymentRuleCmd = &cobra.Command{
	Use: "update-deployment-rule [gate_id] [id]",

	Short: "Update deployment rule",
	Long: `Update deployment rule
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#update-deployment-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentRuleResponse
		var err error

		var body datadogV2.UpdateDeploymentRuleParams
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err = api.UpdateDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-deployment-rule")

		cmd.Println(cmdutil.FormatJSON(res, "deployment_rule"))
	},
}

func init() {

	UpdateDeploymentRuleCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDeploymentRuleCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDeploymentRuleCmd)
}
