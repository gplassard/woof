package datasets

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteDatasetCmd = &cobra.Command{
	Use:     "delete-dataset [dataset_id]",
	Aliases: []string{"delete"},
	Short:   "Delete a dataset",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		_, err := api.DeleteDataset(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-dataset")

	},
}

func init() {
	Cmd.AddCommand(DeleteDatasetCmd)
}
