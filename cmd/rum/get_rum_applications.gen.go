package rum

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRUMApplicationsCmd = &cobra.Command{
	Use:     "get-rum-applications",
	Aliases: []string{"get-applications"},
	Short:   "List all the RUM applications",
	Long: `List all the RUM applications
Documentation: https://docs.datadoghq.com/api/latest/r-u-m/#get-rum-applications`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.RUMApplicationsResponse
		var err error

		api := datadogV2.NewRUMApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.GetRUMApplications(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-rum-applications")

		fmt.Println(cmdutil.FormatJSON(res, "r_u_m_application"))
	},
}

func init() {

	Cmd.AddCommand(GetRUMApplicationsCmd)
}
