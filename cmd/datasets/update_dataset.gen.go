package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDatasetCmd = &cobra.Command{
	Use:     "update-dataset [dataset_id]",
	Aliases: []string{"update"},
	Short:   "Edit a dataset",
	Long: `Edit a dataset
Documentation: https://docs.datadoghq.com/api/latest/datasets/#update-dataset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseSingle
		var err error

		var body datadogV2.DatasetUpdateRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err = api.UpdateDataset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {

	UpdateDatasetCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateDatasetCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateDatasetCmd)
}
