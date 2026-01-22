package confluent_cloud

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateConfluentResourceCmd = &cobra.Command{
	Use: "create-confluent-resource [account_id]",

	Short: "Add resource to Confluent account",
	Long: `Add resource to Confluent account
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#create-confluent-resource`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentResourceResponse
		var err error

		var body datadogV2.ConfluentResourceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-confluent-resource")

		fmt.Println(cmdutil.FormatJSON(res, "confluent_resource"))
	},
}

func init() {

	CreateConfluentResourceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateConfluentResourceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateConfluentResourceCmd)
}
