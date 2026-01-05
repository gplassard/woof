package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListConfluentResourceCmd = &cobra.Command{
	Use: "list-confluent-resource [account_id]",

	Short: "List Confluent Account resources",
	Long: `List Confluent Account resources
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#list-confluent-resource`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentResourcesResponse
		var err error

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err = api.ListConfluentResource(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to list-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {

	Cmd.AddCommand(ListConfluentResourceCmd)
}
