package deployment_gates

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var GetDeploymentGatesEvaluationResultCmd = &cobra.Command{
	Use:     "get-deployment-gates-evaluation-result [id]",
	Aliases: []string{"get-evaluation-result"},
	Short:   "Get a deployment gate evaluation result",
	Long: `Get a deployment gate evaluation result
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#get-deployment-gates-evaluation-result`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGatesEvaluationResultResponse
		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDeploymentGatesEvaluationResult(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to get-deployment-gates-evaluation-result")

		fmt.Println(cmdutil.FormatJSON(res, "deployment_gates_evaluation_result_response"))
	},
}

func init() {

	Cmd.AddCommand(GetDeploymentGatesEvaluationResultCmd)
}
