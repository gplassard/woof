package confluent_cloud

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteConfluentAccountCmd = &cobra.Command{
	Use:   "delete-confluent-account [account_id]",
	
	Short: "Delete Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		_, err := api.DeleteConfluentAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-confluent-account")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteConfluentAccountCmd)
}
