package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteRowsCmd = &cobra.Command{
	Use:   "deleterows [id]",
	Short: "Delete rows",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err := api.DeleteRows(client.NewContext(apiKey, appKey, site), args[0], datadogV2.BatchDeleteRowsRequestArray{})
		if err != nil {
			log.Fatalf("failed to deleterows: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(DeleteRowsCmd)
}
