package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpsertRowsCmd = &cobra.Command{
	Use:   "upsert-rows [id]",
	Short: "Upsert rows",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err := api.UpsertRows(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BatchUpsertRowsRequestArray{})
		if err != nil {
			log.Fatalf("failed to upsert-rows: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UpsertRowsCmd)
}
