package deployment_gates

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDeploymentGatesCmd = &cobra.Command{
	Use:     "list-deployment-gates",
	Aliases: []string{"list"},
	Short:   "Get all deployment gates",
	Long: `Get all deployment gates
Documentation: https://docs.datadoghq.com/api/latest/deployment-gates/#list-deployment-gates`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DeploymentGatesListResponse
		var err error

		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDeploymentGates(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-deployment-gates")

		fmt.Println(cmdutil.FormatJSON(res, "deployment_gate"))
	},
}

func init() {

	Cmd.AddCommand(ListDeploymentGatesCmd)
}
