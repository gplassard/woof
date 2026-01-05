package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateConfluentAccountCmd = &cobra.Command{
	Use: "create-confluent-account",

	Short: "Add Confluent account",
	Long: `Add Confluent account
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#create-confluent-account`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentAccountResponse
		var err error

		var body datadogV2.ConfluentAccountCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		res, _, err = api.CreateConfluentAccount(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-confluent-account")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-accounts"))
	},
}

func init() {

	CreateConfluentAccountCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateConfluentAccountCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateConfluentAccountCmd)
}
