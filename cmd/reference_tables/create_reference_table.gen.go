package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateReferenceTableCmd = &cobra.Command{
	Use:   "create-reference-table",
	Aliases: []string{ "create", },
	Short: "Create reference table",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.CreateReferenceTable(client.NewContext(apiKey, appKey, site), datadogV2.CreateTableRequest{})
		if err != nil {
			log.Fatalf("failed to create-reference-table: %v", err)
		}

		cmdutil.PrintJSON(res, "reference_table")
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableCmd)
}
