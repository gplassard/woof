package ci_visibility_tests

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCIAppTestEventsCmd = &cobra.Command{
	Use: "list-ci-app-test-events",

	Short: "Get a list of tests events",
	Long: `Get a list of tests events
Documentation: https://docs.datadoghq.com/api/latest/c-i-visibility-tests/#list-ci-app-test-events`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.CIAppTestEventsResponse
		var err error

		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListCIAppTestEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-ci-app-test-events")

		fmt.Println(cmdutil.FormatJSON(res, "c_i_app_test_event"))
	},
}

func init() {

	Cmd.AddCommand(ListCIAppTestEventsCmd)
}
