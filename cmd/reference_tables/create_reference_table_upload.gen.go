package reference_tables

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var CreateReferenceTableUploadCmd = &cobra.Command{
	Use:     "create-reference-table-upload [payload]",
	Aliases: []string{"create-upload"},
	Short:   "Create reference table upload",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CreateUploadResponse
		var err error

		var body datadogV2.CreateUploadRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err = api.CreateReferenceTableUpload(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to create-reference-table-upload")

		cmd.Println(cmdutil.FormatJSON(res, "upload"))
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableUploadCmd)
}
