package r_u_m

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var DeleteRUMApplicationCmd = &cobra.Command{
	Use: "delete-rum-application [id]",

	Short: "Delete a RUM application",
	Long: `Delete a RUM application
Documentation: https://docs.datadoghq.com/api/latest/r-u-m/#delete-rum-application`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		var err error

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		_, err = api.DeleteRUMApplication(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-rum-application")

	},
}

func init() {

	Cmd.AddCommand(DeleteRUMApplicationCmd)
}
