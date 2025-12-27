package datasets

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetAllDatasetsCmd = &cobra.Command{
	Use:   "getalldatasets",
	Short: "Get all datasets",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewDatasetsApi(client.NewAPIClient())
		res, _, err := api.GetAllDatasets(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getalldatasets: %v", err)
		}

		cmdutil.PrintJSON(res, "datasets")
	},
}

func init() {
	Cmd.AddCommand(GetAllDatasetsCmd)
}
