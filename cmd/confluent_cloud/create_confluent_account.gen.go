package confluent_cloud

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateConfluentAccountCmd = &cobra.Command{
	Use:   "create-confluent-account",
	
	Short: "Add Confluent account",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.CreateConfluentAccount(client.NewContext(apiKey, appKey, site), datadogV2.ConfluentAccountCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create-confluent-account: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent-cloud-accounts")
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentAccountCmd)
}
