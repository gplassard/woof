package reference_tables

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

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
		cmdutil.HandleError(err, "failed to upsert-rows")

		
	},
}

func init() {
	Cmd.AddCommand(UpsertRowsCmd)
}
