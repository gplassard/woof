package datasets

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDatasetCmd = &cobra.Command{
	Use:     "get-dataset [dataset_id]",
	Aliases: []string{"get"},
	Short:   "Get a single dataset by ID",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.GetDataset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(GetDatasetCmd)
}
