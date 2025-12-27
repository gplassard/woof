package confluent_cloud

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateConfluentResourceCmd = &cobra.Command{
	Use:   "create_confluent_resource [account_id]",
	Short: "Add resource to Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.CreateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ConfluentResourceRequest{})
		if err != nil {
			log.Fatalf("failed to create_confluent_resource: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent_cloud")
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentResourceCmd)
}
