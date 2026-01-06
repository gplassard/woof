package case_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCaseCommentCmd = &cobra.Command{
	Use: "delete-case-comment [case_id] [cell_id]",

	Short: "Delete case comment",
	Long: `Delete case comment
Documentation: https://docs.datadoghq.com/api/latest/case-management/#delete-case-comment`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCaseManagementApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteCaseComment(client.NewContext(apiKey, appKey, site), args[0], args[1])
		cmdutil.HandleError(err, "failed to delete-case-comment")

	},
}

func init() {

	Cmd.AddCommand(DeleteCaseCommentCmd)
}
