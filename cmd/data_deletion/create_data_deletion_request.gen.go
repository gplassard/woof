package data_deletion

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDataDeletionRequestCmd = &cobra.Command{
	Use:     "create-data-deletion-request [product]",
	Aliases: []string{"create-request"},
	Short:   "Creates a data deletion request",
	Long: `Creates a data deletion request
Documentation: https://docs.datadoghq.com/api/latest/data-deletion/#create-data-deletion-request`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateDataDeletionResponseBody
		var err error

		var body datadogV2.CreateDataDeletionRequestBody
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDataDeletionApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDataDeletionRequest(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to create-data-deletion-request")

		cmd.Println(cmdutil.FormatJSON(res, "data_deletion"))
	},
}

func init() {

	CreateDataDeletionRequestCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDataDeletionRequestCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDataDeletionRequestCmd)
}
