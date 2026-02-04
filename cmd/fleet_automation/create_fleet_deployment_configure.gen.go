package fleet_automation

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFleetDeploymentConfigureCmd = &cobra.Command{
	Use: "create-fleet-deployment-configure",

	Short: "Create a configuration deployment",
	Long: `Create a configuration deployment
Documentation: https://docs.datadoghq.com/api/latest/fleet-automation/#create-fleet-deployment-configure`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.FleetDeploymentResponse
		var err error

		var body datadogV2.FleetDeploymentConfigureCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateFleetDeploymentConfigure(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-fleet-deployment-configure")

		fmt.Println(cmdutil.FormatJSON(res, "fleet_deployment_configure"))
	},
}

func init() {

	CreateFleetDeploymentConfigureCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateFleetDeploymentConfigureCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateFleetDeploymentConfigureCmd)
}
