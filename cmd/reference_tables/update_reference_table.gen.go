package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateReferenceTableCmd = &cobra.Command{
	Use:   "update-reference-table [id]",
	Aliases: []string{ "update", },
	Short: "Update reference table",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err := api.UpdateReferenceTable(client.NewContext(apiKey, appKey, site), args[0], datadogV2.PatchTableRequest{})
		if err != nil {
			log.Fatalf("failed to update-reference-table: %v", err)
		}

		
	},
}

func init() {
	Cmd.AddCommand(UpdateReferenceTableCmd)
}
