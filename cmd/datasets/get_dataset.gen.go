package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDatasetCmd = &cobra.Command{
	Use:     "get-dataset [dataset_id]",
	Aliases: []string{"get"},
	Short:   "Get a single dataset by ID",
	Long: `Get a single dataset by ID
Documentation: https://docs.datadoghq.com/api/latest/datasets/#get-dataset`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseSingle
		var err error

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetDataset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {

	Cmd.AddCommand(GetDatasetCmd)
}
