package status_pages

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var DeleteComponentCmd = &cobra.Command{
	Use: "delete-component [page_id] [component_id]",

	Short: "Delete component",
	Long: `Delete component
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#delete-component`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteComponent(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), uuid.MustParse(args[1]))
		cmdutil.HandleError(err, "failed to delete-component")

	},
}

func init() {

	Cmd.AddCommand(DeleteComponentCmd)
}
