package ci_visibility_tests

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCIAppTestEventsCmd = &cobra.Command{
	Use:   "listciapptestevents",
	Short: "Get a list of tests events",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewCIVisibilityTestsApi(client.NewAPIClient())
		res, _, err := api.ListCIAppTestEvents(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to listciapptestevents: %v", err)
		}

		cmdutil.PrintJSON(res, "ci_visibility_tests")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppTestEventsCmd)
}
