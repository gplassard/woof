package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpdateReferenceTableCmd = &cobra.Command{
	Use:     "update-reference-table [id]",
	Aliases: []string{"update"},
	Short:   "Update reference table",
	Long: `Update reference table
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#update-reference-table`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		var body datadogV2.PatchTableRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.UpdateReferenceTable(client.NewContext(apiKey, appKey, site), args[0], body)
		cmdutil.HandleError(err, "failed to update-reference-table")

	},
}

func init() {

	UpdateReferenceTableCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpdateReferenceTableCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpdateReferenceTableCmd)
}
