package rum

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRUMApplicationsCmd = &cobra.Command{
	Use:   "get-r-u-m-applications",
	Short: "List all the RUM applications",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.GetRUMApplications(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to get-r-u-m-applications: %v", err)
		}

		cmdutil.PrintJSON(res, "rum_application")
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationsCmd)
}
