package network_device_monitoring

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDeviceUserTagsCmd = &cobra.Command{
	Use:   "update-device-user-tags [device_id]",
	
	Short: "Update the tags for a device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ListTagsResponse{})
		cmdutil.HandleError(err, "failed to update-device-user-tags")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(UpdateDeviceUserTagsCmd)
}
