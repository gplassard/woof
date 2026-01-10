package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetFleetDeploymentCmd = &cobra.Command{
	Use: "get-fleet-deployment [deployment_id]",

	Short: "Get a configuration deployment by ID",
	Long: `Get a configuration deployment by ID
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#get-fleet-deployment`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentResponse
		var err error

		optionalParams := datadogV2.NewGetFleetDeploymentOptionalParameters()

		if cmd.Flags().Changed("limit") {
			val, _ := cmd.Flags().GetInt64("limit")
			optionalParams.WithLimit(val)
		}

		if cmd.Flags().Changed("page") {
			val, _ := cmd.Flags().GetInt64("page")
			optionalParams.WithPage(val)
		}

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetFleetDeployment(client.NewContext(apiKey, appKey, site), args[0], *optionalParams)
		cmdutil.HandleError(err, "failed to get-fleet-deployment")

		fmt.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {

	GetFleetDeploymentCmd.Flags().Int64("limit", 0, "Maximum number of hosts to return per page. Default is 50, maximum is 100.")

	GetFleetDeploymentCmd.Flags().Int64("page", 0, "Page index for pagination (zero-based). Use this to retrieve subsequent pages of hosts.")

	Cmd.AddCommand(GetFleetDeploymentCmd)
}
