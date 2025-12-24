package rum

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetRUMApplicationsCmd = &cobra.Command{
	Use:   "getrumapplications",
	Short: "List all the RUM applications",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.GetRUMApplications(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to getrumapplications: %v", err)
		}

		cmdutil.PrintJSON(res, "rum")
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationsCmd)
}
