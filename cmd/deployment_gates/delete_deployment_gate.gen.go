package deployment_gates

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDeploymentGateCmd = &cobra.Command{
	Use:     "delete-deployment-gate [id]",
	Aliases: []string{"delete"},
	Short:   "Delete deployment gate",
	Long: `Delete deployment gate
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#delete-deployment-gate`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteDeploymentGate(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-deployment-gate")

	},
}

func init() {

	Cmd.AddCommand(DeleteDeploymentGateCmd)
}
