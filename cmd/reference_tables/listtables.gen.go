package reference_tables

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListTablesCmd = &cobra.Command{
	Use:   "listtables",
	Short: "List tables",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.ListTables(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listtables: %v", err)
		}

		cmdutil.PrintJSON(res, "reference_tables")
	},
}

func init() {
	Cmd.AddCommand(ListTablesCmd)
}
