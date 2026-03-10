package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var ListComponentsCmd = &cobra.Command{
	Use: "list-components [page_id]",

	Short: "List components",
	Long: `List components
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#list-components`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.StatusPagesComponentArray
		var err error

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListComponents(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]))
		cmdutil.HandleError(err, "failed to list-components")

		fmt.Println(cmdutil.FormatJSON(res, "components"))
	},
}

func init() {

	Cmd.AddCommand(ListComponentsCmd)
}
