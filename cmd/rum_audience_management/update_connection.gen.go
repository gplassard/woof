package rum_audience_management

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateConnectionCmd = &cobra.Command{
	Use:   "update-connection [entity]",
	
	Short: "Update connection",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewRumAudienceManagementApi(client.NewAPIClient())
		_, err := api.UpdateConnection(client.NewContext(apiKey, appKey, site), args[0], datadogV2.UpdateConnectionRequest{})
		cmdutil.HandleError(err, "failed to update-connection")

		
	},
}

func init() {
	Cmd.AddCommand(UpdateConnectionCmd)
}
