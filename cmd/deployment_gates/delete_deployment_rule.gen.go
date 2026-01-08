package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDeploymentRuleCmd = &cobra.Command{
	Use: "delete-deployment-rule [gate_id] [id]",

	Short: "Delete deployment rule",
	Long: `Delete deployment rule
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#delete-deployment-rule`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteDeploymentRule(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-deployment-rule")

	},
}

func init() {

	Cmd.AddCommand(DeleteDeploymentRuleCmd)
}
