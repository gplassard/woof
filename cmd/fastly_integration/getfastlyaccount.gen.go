package fastly_integration

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetFastlyAccountCmd = &cobra.Command{
	Use:   "getfastlyaccount [account_id]",
	Short: "Get Fastly account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewFastlyIntegrationApi(client.NewAPIClient())
		res, _, err := api.GetFastlyAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getfastlyaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "fastly_integration")
	},
}

func init() {
	Cmd.AddCommand(GetFastlyAccountCmd)
}
