package fleet_automation

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFleetDeploymentUpgradeCmd = &cobra.Command{
	Use: "create-fleet-deployment-upgrade",

	Short: "Upgrade hosts",
	Long: `Upgrade hosts
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#create-fleet-deployment-upgrade`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentResponse
		var err error

		var body datadogV2.FleetDeploymentPackageUpgradeCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err = api.CreateFleetDeploymentUpgrade(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-fleet-deployment-upgrade")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {

	CreateFleetDeploymentUpgradeCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateFleetDeploymentUpgradeCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateFleetDeploymentUpgradeCmd)
}
