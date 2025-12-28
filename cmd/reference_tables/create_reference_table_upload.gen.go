package reference_tables

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var CreateReferenceTableUploadCmd = &cobra.Command{
	Use:   "create-reference-table-upload",
	Aliases: []string{ "create-upload", },
	Short: "Create reference table upload",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewReferenceTablesApi(client.NewAPIClient())
		res, _, err := api.CreateReferenceTableUpload(client.NewContext(apiKey, appKey, site), datadogV2.CreateUploadRequest{})
		cmdutil.HandleError(err, "failed to create-reference-table-upload")

		cmdutil.PrintJSON(res, "upload")
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableUploadCmd)
}
