package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateDatasetCmd = &cobra.Command{
	Use:     "create-dataset [payload]",
	Aliases: []string{"create"},
	Short:   "Create a dataset",
	Long: `Create a dataset
Documentation: https://docs.datadoghq.com/api/latest/datasets/#create-dataset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseSingle
		var err error

		var body datadogV2.DatasetCreateRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err = api.CreateDataset(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(CreateDatasetCmd)
}
