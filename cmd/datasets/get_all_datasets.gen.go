package datasets

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetAllDatasetsCmd = &cobra.Command{
	Use:     "get-all-datasets",
	Aliases: []string{"get-all"},
	Short:   "Get all datasets",
	Long: `Get all datasets
Documentation: https://docs.datadoghq.com/api/latest/datasets/#get-all-datasets`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.DatasetResponseMulti
		var err error

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetAllDatasets(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-all-datasets")

		fmt.Println(cmdutil.FormatJSON(res, "all_dataset"))
	},
}

func init() {

	Cmd.AddCommand(GetAllDatasetsCmd)
}
