package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpdateReferenceTableCmd = &cobra.Command{
	Use:     "update-reference-table [id] [payload]",
	Aliases: []string{"update"},
	Short:   "Update reference table",
	Long: `Update reference table
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#update-reference-table`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.PatchTableRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		_, err = api.UpdateReferenceTable(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-reference-table")

	},
}

func init() {
	Cmd.AddCommand(UpdateReferenceTableCmd)
}
