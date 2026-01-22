package reference_tables

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var CreateReferenceTableUploadCmd = &cobra.Command{
	Use:     "create-reference-table-upload",
	Aliases: []string{"create-upload"},
	Short:   "Create reference table upload",
	Long: `Create reference table upload
Documentation: https://docs.datadoghq.com/api/latest/reference-tables/#create-reference-table-upload`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateUploadResponse
		var err error

		var body datadogV2.CreateUploadRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateReferenceTableUpload(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-reference-table-upload")

		fmt.Println(cmdutil.FormatJSON(res, "reference_table_upload"))
	},
}

func init() {

	CreateReferenceTableUploadCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateReferenceTableUploadCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateReferenceTableUploadCmd)
}
