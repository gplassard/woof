package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTableCmd = &cobra.Command{
	Use:   "get-table [id]",
	
	Short: "Get table",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.GetTable(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to get-table: %v", err)
		}

		cmdutil.PrintJSON(res, "reference_table")
	},
}

func init() {
	Cmd.AddCommand(GetTableCmd)
}
