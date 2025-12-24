package datasets

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDatasetCmd = &cobra.Command{
	Use:   "updatedataset [dataset_id]",
	Short: "Edit a dataset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.UpdateDataset(client.NewContext(apiKey, appKey, site), args[0], datadogV2.DatasetUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatedataset: %v", err)
		}

		cmdutil.PrintJSON(res, "datasets")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatasetCmd)
}
