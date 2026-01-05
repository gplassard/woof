package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetConfluentResourceCmd = &cobra.Command{
	Use: "get-confluent-resource [account_id] [resource_id]",

	Short: "Get resource from Confluent account",
	Long: `Get resource from Confluent account
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#get-confluent-resource`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentResourceResponse
		var err error

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err = api.GetConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to get-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {

	Cmd.AddCommand(GetConfluentResourceCmd)
}
