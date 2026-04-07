package integrations

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListIntegrationsCmd = &cobra.Command{
	Use:     "list-integrations",
	Aliases: []string{"list"},
	Short:   "List integrations",
	Long: `List integrations
Documentation: https://docs.datadoghq.com/api/latest/integrations/#list-integrations`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListIntegrationsResponse
		var err error

		api := datadogV2.NewIntegrationsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListIntegrations(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-integrations")

		fmt.Println(cmdutil.FormatJSON(res, "integration"))
	},
}

func init() {

	Cmd.AddCommand(ListIntegrationsCmd)
}
