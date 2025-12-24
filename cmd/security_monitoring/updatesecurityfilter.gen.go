package security_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateSecurityFilterCmd = &cobra.Command{
	Use:   "updatesecurityfilter [security_filter_id]",
	Short: "Update a security filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateSecurityFilter(client.NewContext(apiKey, appKey, site), args[0], datadogV2.SecurityFilterUpdateRequest{})
		if err != nil {
			log.Fatalf("failed to updatesecurityfilter: %v", err)
		}

		cmdutil.PrintJSON(res, "security_monitoring")
	},
}

func init() {
	Cmd.AddCommand(UpdateSecurityFilterCmd)
}
