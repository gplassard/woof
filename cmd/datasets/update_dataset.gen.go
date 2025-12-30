package datasets

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateDatasetCmd = &cobra.Command{
	Use:     "update-dataset [dataset_id]",
	Aliases: []string{"update"},
	Short:   "Edit a dataset",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.UpdateDataset(client.NewContext(apiKey, appKey, site), args[0], datadogV2.DatasetUpdateRequest{})
		cmdutil.HandleError(err, "failed to update-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDatasetCmd)
}
