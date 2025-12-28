package confluent_cloud

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetConfluentAccountCmd = &cobra.Command{
	Use:   "get-confluent-account [account_id]",
	
	Short: "Get Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.GetConfluentAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-confluent-account")

		cmdutil.PrintJSON(res, "confluent-cloud-accounts")
	},
}

func init() {
	Cmd.AddCommand(GetConfluentAccountCmd)
}
