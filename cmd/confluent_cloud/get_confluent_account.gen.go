package confluent_cloud

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetConfluentAccountCmd = &cobra.Command{
	Use: "get-confluent-account [account_id]",

	Short: "Get Confluent account",
	Long: `Get Confluent account
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#get-confluent-account`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentAccountResponse
		var err error

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetConfluentAccount(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-confluent-account")

		fmt.Println(cmdutil.FormatJSON(res, "confluent_account"))
	},
}

func init() {

	Cmd.AddCommand(GetConfluentAccountCmd)
}
