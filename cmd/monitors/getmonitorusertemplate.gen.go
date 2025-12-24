package monitors

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var GetMonitorUserTemplateCmd = &cobra.Command{
	Use:   "getmonitorusertemplate [template_id]",
	Short: "Get a monitor user template",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewMonitorsApi(client.NewAPIClient())
		res, _, err := api.GetMonitorUserTemplate(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to getmonitorusertemplate: %v", err)
		}

		cmdutil.PrintJSON(res, "monitors")
	},
}

func init() {
	Cmd.AddCommand(GetMonitorUserTemplateCmd)
}
