package ci_visibility_tests

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCIAppTestEventsCmd = &cobra.Command{
	Use:   "list-ci-app-test-events",
	
	Short: "Get a list of tests events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.ListCIAppTestEvents(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-ci-app-test-events")

		cmdutil.PrintJSON(res, "citest")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppTestEventsCmd)
}
