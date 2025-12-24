package deployment_gates

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDeploymentGateCmd = &cobra.Command{
	Use:   "getdeploymentgate [id]",
	Short: "Get deployment gate",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDeploymentGatesApi(client.NewAPIClient())
		res, _, err := api.GetDeploymentGate(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getdeploymentgate: %v", err)
		}

		cmdutil.PrintJSON(res, "deployment_gates")
	},
}

func init() {
	Cmd.AddCommand(GetDeploymentGateCmd)
}
