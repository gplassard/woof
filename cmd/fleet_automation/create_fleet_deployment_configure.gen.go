package fleet_automation

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateFleetDeploymentConfigureCmd = &cobra.Command{
	Use: "create-fleet-deployment-configure",

	Short: "Create a configuration deployment",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewFleetAutomationApi(client.NewAPIClient())
		res, _, err := api.CreateFleetDeploymentConfigure(client.NewContext(apiKey, appKey, site), datadogV2.FleetDeploymentConfigureCreateRequest{})
		cmdutil.HandleError(err, "failed to create-fleet-deployment-configure")

		cmd.Println(cmdutil.FormatJSON(res, "deployment"))
	},
}

func init() {
	Cmd.AddCommand(CreateFleetDeploymentConfigureCmd)
}
