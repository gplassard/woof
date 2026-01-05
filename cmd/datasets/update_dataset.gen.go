package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateDatasetCmd = &cobra.Command{
	Use:     "update-dataset [dataset_id] [payload]",
	Aliases: []string{"update"},
	Short:   "Edit a dataset",
	Long: `Edit a dataset
Documentation: https://docs.datadoghq.com/api/latest/datasets/#update-dataset`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseSingle
		var err error

		var body datadogV2.DatasetUpdateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err = api.UpdateDataset(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDatasetCmd)
}
