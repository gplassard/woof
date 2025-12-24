package deployment_gates

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDeploymentGateCmd = &cobra.Command{
	Use:   "createdeploymentgate",
	Short: "Create deployment gate",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.CreateDeploymentGate(client.NewContext(apiKey, appKey, site), datadogV2.CreateDeploymentGateParams{})
		if err != nil {
			log.Fatalf("failed to createdeploymentgate: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gates")
	},
}

func init() {
	Cmd.AddCommand(CreateDeploymentGateCmd)
}
