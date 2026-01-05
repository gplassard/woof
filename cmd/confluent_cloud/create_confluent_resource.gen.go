package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateConfluentResourceCmd = &cobra.Command{
	Use: "create-confluent-resource [account_id] [payload]",

	Short: "Add resource to Confluent account",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var body datadogV2.ConfluentResourceRequest
		err := json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err := api.CreateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {
	Cmd.AddCommand(CreateConfluentResourceCmd)
}
