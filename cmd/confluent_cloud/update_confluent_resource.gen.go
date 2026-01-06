package confluent_cloud

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateConfluentResourceCmd = &cobra.Command{
	Use: "update-confluent-resource [account_id] [resource_id]",

	Short: "Update resource in Confluent account",
	Long: `Update resource in Confluent account
Documentation: https://docs.datadoghq.com/api/latest/confluent-cloud/#update-confluent-resource`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ConfluentResourceResponse
		var err error

		var body datadogV2.ConfluentResourceRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewConfluentCloudApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpdateConfluentResource(client.NewContext(apiKey, appKey, site), args[0], args[1], body)
		cmdutil.HandleError(err, "failed to update-confluent-resource")

		cmd.Println(cmdutil.FormatJSON(res, "confluent-cloud-resources"))
	},
}

func init() {

	UpdateConfluentResourceCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateConfluentResourceCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateConfluentResourceCmd)
}
