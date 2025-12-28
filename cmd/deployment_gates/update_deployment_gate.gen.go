package deployment_gates

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDeploymentGateCmd = &cobra.Command{
	Use:   "update_deployment_gate [id]",
	Short: "Update deployment gate",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.UpdateDeploymentGate(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateDeploymentGateParams{})
		if err != nil {
			log.Fatalf("failed to update_deployment_gate: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gate")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeploymentGateCmd)
}
