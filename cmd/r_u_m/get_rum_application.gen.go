package r_u_m

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRUMApplicationCmd = &cobra.Command{
	Use: "get-rum-application [id]",

	Short: "Get a RUM application",
	Long: `Get a RUM application
Documentation: https://docs.datadoghq.com/api/latest/r-u-m/#get-rum-application`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMApplicationResponse
		var err error

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRUMApplication(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-rum-application")

		fmt.Println(cmdutil.FormatJSON(res, "r_u_m_application"))
	},
}

func init() {

	Cmd.AddCommand(GetRUMApplicationCmd)
}
