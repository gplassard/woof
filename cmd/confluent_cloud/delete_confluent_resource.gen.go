package confluent_cloud

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteConfluentResourceCmd = &cobra.Command{
	Use: "delete-confluent-resource [account_id] [resource_id]",

	Short: "Delete resource from Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		_, err := api.DeleteConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-confluent-resource")

	},
}

func init() {
	Cmd.AddCommand(DeleteConfluentResourceCmd)
}
