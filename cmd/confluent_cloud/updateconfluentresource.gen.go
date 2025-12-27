package confluent_cloud

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateConfluentResourceCmd = &cobra.Command{
	Use:   "updateconfluentresource [account_id] [resource_id]",
	Short: "Update resource in Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.UpdateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.ConfluentResourceRequest{})
		if err != nil {
			log.Fatalf("failed to updateconfluentresource: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentResourceCmd)
}
