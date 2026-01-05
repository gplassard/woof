package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListConfluentAccountCmd = &cobra.Command{
	Use: "list-confluent-account",

	Short: "List Confluent accounts",
	Long: `List Confluent accounts
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#list-confluent-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentAccountsResponse
		var err error

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err = api.ListConfluentAccount(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-confluent-account")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-accounts"))
	},
}

func init() {

	Cmd.AddCommand(ListConfluentAccountCmd)
}
