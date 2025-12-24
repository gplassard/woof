package confluent_cloud

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListConfluentAccountCmd = &cobra.Command{
	Use:   "listconfluentaccount",
	Short: "List Confluent accounts",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.ListConfluentAccount(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listconfluentaccount: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(ListConfluentAccountCmd)
}
