package confluent_cloud

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetConfluentResourceCmd = &cobra.Command{
	Use:   "getconfluentresource [account_id] [resource_id]",
	Short: "Get resource from Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.GetConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1])
		if err != nil {
			log.Fatalf("failed to getconfluentresource: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(GetConfluentResourceCmd)
}
