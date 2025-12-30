package reference_tables

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateReferenceTableCmd = &cobra.Command{
	Use:     "create-reference-table",
	Aliases: []string{"create"},
	Short:   "Create reference table",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.CreateReferenceTable(client.NewContext(apiKey, appKey, site), datadogV2.CreateTableRequest{})
		cmdutil.HandleError(err, "failed to create-reference-table")

		cmd.Println(cmdutil.FormatJSON(res, "reference_table"))
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableCmd)
}
