package rum

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetRUMApplicationsCmd = &cobra.Command{
	Use:     "get-rum-applications",
	Aliases: []string{"get-applications"},
	Short:   "List all the RUM applications",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRUMApi(client.NewAPIClient())
		res, _, err := api.GetRUMApplications(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to get-rum-applications")

		cmd.Println(cmdutil.FormatJSON(res, "rum_application"))
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationsCmd)
}
