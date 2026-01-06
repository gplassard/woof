package case_management_type

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteCaseTypeCmd = &cobra.Command{
	Use: "delete-case-type [case_type_id]",

	Short: "Delete a case type",
	Long: `Delete a case type
Documentation: https://docs.datadoghq.com/api/latest/case-management-type/#delete-case-type`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewCaseManagementTypeApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteCaseType(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-case-type")

	},
}

func init() {

	Cmd.AddCommand(DeleteCaseTypeCmd)
}
