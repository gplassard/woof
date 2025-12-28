package confluent_cloud

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateConfluentAccountCmd = &cobra.Command{
	Use:   "update_confluent_account [account_id]",
	Short: "Update Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.UpdateConfluentAccount(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ConfluentAccountUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to update_confluent_account: %v", err)
		}

		cmdutil.PrintJSON(res, "confluent-cloud-accounts")
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentAccountCmd)
}
