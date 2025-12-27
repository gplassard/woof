package datasets

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetDatasetCmd = &cobra.Command{
	Use:   "get_dataset [dataset_id]",
	Short: "Get a single dataset by ID",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.GetDataset(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get_dataset: %v", err)
		}

		cmdutil.PrintJSON(res, "datasets")
	},
}

func init() {
	Cmd.AddCommand(GetDatasetCmd)
}
