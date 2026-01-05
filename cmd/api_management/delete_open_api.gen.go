package api_management

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteOpenAPICmd = &cobra.Command{
	Use: "delete-open-api [id]",

	Short: "Delete an API",
	Long: `Delete an API
Documentation: https://docs.datadoghq.com/api/latest/api-management/#delete-open-api`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewAPIManagementApi(client.NewAPIClient())
		_, err = api.DeleteOpenAPI(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to delete-open-api")

	},
}

func init() {
	Cmd.AddCommand(DeleteOpenAPICmd)
}
