package reference_tables

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetTableCmd = &cobra.Command{
	Use:   "gettable [id]",
	Short: "Get table",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.GetTable(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to gettable: %v", err)
		}

		cmdutil.PrintJSON(res, "reference_tables")
	},
}

func init() {
	Cmd.AddCommand(GetTableCmd)
}
