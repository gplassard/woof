package confluent_cloud

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListConfluentResourceCmd = &cobra.Command{
	Use:   "listconfluentresource [account_id]",
	Short: "List Confluent Account resources",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.ListConfluentResource(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listconfluentresource: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(ListConfluentResourceCmd)
}
