package confluent_cloud

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateConfluentResourceCmd = &cobra.Command{
	Use:   "update-confluent-resource [account_id] [resource_id]",
	
	Short: "Update resource in Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.UpdateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1], datadogV2.ConfluentResourceRequest{})
		cmdutil.HandleError(err, "failed to update-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {
	Cmd.AddCommand(UpdateConfluentResourceCmd)
}
