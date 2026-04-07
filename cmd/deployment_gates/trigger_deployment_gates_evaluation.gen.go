package deployment_gates

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var TriggerDeploymentGatesEvaluationCmd = &cobra.Command{
	Use:     "trigger-deployment-gates-evaluation",
	Aliases: []string{"trigger-evaluation"},
	Short:   "Trigger a deployment gate evaluation",
	Long: `Trigger a deployment gate evaluation
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#trigger-deployment-gates-evaluation`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGatesEvaluationResponse
		var err error

		var body datadogV2.DeploymentGatesEvaluationRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.TriggerDeploymentGatesEvaluation(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to trigger-deployment-gates-evaluation")

		fmt.Println(cmdutil.FormatJSON(res, "deployment_gates_evaluation_response"))
	},
}

func init() {

	TriggerDeploymentGatesEvaluationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	TriggerDeploymentGatesEvaluationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(TriggerDeploymentGatesEvaluationCmd)
}
