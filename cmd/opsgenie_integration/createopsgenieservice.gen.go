package opsgenie_integration

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateOpsgenieServiceCmd = &cobra.Command{
	Use:   "createopsgenieservice",
	Short: "Create a new service object",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewOpsgenieIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateOpsgenieService(client.NewContext(apiKey, appKey, site), datadogV2.OpsgenieServiceCreateRequest{})
		if err != nil {
			log.Fatalf("failed to createopsgenieservice: %v", err)
		}

		cmdutil.PrintJSON(res, "opsgenie_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateOpsgenieServiceCmd)
}
