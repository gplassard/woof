package security_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var DeleteSecurityFilterCmd = &cobra.Command{
	Use:   "delete-security-filter [security_filter_id]",
	
	Short: "Delete a security filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		_, err := api.DeleteSecurityFilter(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to delete-security-filter")

		
	},
}

func init() {
	Cmd.AddCommand(DeleteSecurityFilterCmd)
}
