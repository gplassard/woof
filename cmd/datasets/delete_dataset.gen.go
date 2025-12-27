package datasets

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteDatasetCmd = &cobra.Command{
	Use:   "delete_dataset [dataset_id]",
	Short: "Delete a dataset",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		_, err := api.DeleteDataset(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to delete_dataset: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteDatasetCmd)
}
