package confluent_cloud

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateConfluentResourceCmd = &cobra.Command{
	Use:   "create-confluent-resource [account_id]",
	
	Short: "Add resource to Confluent account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.CreateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ConfluentResourceRequest{})
		cmdutil.HandleError(err, "failed to create-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentResourceCmd)
}
