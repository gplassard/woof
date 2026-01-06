package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDatasetCmd = &cobra.Command{
	Use:     "create-dataset",
	Aliases: []string{"create"},
	Short:   "Create a dataset",
	Long: `Create a dataset
Documentation: https://docs.datadoghq.com/api/latest/datasets/#create-dataset`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseSingle
		var err error

		var body datadogV2.DatasetCreateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDataset(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {

	CreateDatasetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDatasetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDatasetCmd)
}
