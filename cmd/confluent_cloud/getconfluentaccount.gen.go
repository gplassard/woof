package confluent_cloud

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetConfluentAccountCmd = &cobra.Command{
	Use:   "getconfluentaccount [account_id]",
	Short: "Get Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.GetConfluentAccount(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getconfluentaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(GetConfluentAccountCmd)
}
