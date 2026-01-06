package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteTableCmd = &cobra.Command{
	Use: "delete-table [id]",

	Short: "Delete table",
	Long: `Delete table
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#delete-table`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteTable(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-table")

	},
}

func init() {

	Cmd.AddCommand(DeleteTableCmd)
}
