package fastly_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateFastlyServiceCmd = &cobra.Command{
	Use:   "createfastlyservice [account_id]",
	Short: "Add Fastly service",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.CreateFastlyService(client.NewContext(apiKey, appKey, site), args[0], datadogV2.FastlyServiceRequest{})
		if err != nil {
			log.Fatalf("failed to createfastlyservice: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly_integration")
	},
}

func init() {
	Cmd.AddCommand(CreateFastlyServiceCmd)
}
