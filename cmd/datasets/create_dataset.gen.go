package datasets

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateDatasetCmd = &cobra.Command{
	Use:     "create-dataset",
	Aliases: []string{"create"},
	Short:   "Create a dataset",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.CreateDataset(client.NewContext(apiKey, appKey, site), datadogV2.DatasetCreateRequest{})
		cmdutil.HandleError(err, "failed to create-dataset")

		cmd.Println(cmdutil.FormatJSON(res, "dataset"))
	},
}

func init() {
	Cmd.AddCommand(CreateDatasetCmd)
}
