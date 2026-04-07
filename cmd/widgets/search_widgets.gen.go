package widgets

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var SearchWidgetsCmd = &cobra.Command{
	Use:     "search-widgets [experience_type]",
	Aliases: []string{"search"},
	Short:   "Search widgets",
	Long: `Search widgets
Documentation: https://docs.datadoghq.com/api/latest/widgets/#search-widgets`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.WidgetListResponse
		var err error

		api := datadogV2.NewWidgetsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.SearchWidgets(client.NewContext(apiKey, appKey, site), datadogV2.WidgetExperienceType(args[0]))
		cmdutil.HandleError(err, "failed to search-widgets")

		fmt.Println(cmdutil.FormatJSON(res, "widgets"))
	},
}

func init() {

	Cmd.AddCommand(SearchWidgetsCmd)
}
