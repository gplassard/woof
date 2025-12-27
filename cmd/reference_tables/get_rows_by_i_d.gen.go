package reference_tables

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	"strings"
	"github.com/spf13/cobra"
	
)

var GetRowsByIDCmd = &cobra.Command{
	Use:   "get_rows_by_i_d [id] [row_id]",
	Short: "Get rows by id",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.GetRowsByID(client.NewContext(apiKey, appKey, site), args[0], strings.Split(args[1], ","))
		if err != nil {
			log.Fatalf("failed to get_rows_by_i_d: %v", err)
		}

		cmdutil.PrintJSON(res, "reference_tables")
	},
}

func init() {
	Cmd.AddCommand(GetRowsByIDCmd)
}
