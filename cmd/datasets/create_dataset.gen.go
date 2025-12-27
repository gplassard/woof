package datasets

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateDatasetCmd = &cobra.Command{
	Use:   "create_dataset",
	Short: "Create a dataset",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.CreateDataset(client.NewContext(apiKey, appKey, site), datadogV2.DatasetCreateRequest{})
		if err != nil {
			log.Fatalf("failed to create_dataset: %v", err)
		}

		cmdutil.PrintJSON(res, "datasets")
	},
}

func init() {
	Cmd.AddCommand(CreateDatasetCmd)
}
