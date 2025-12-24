package network_device_monitoring

import (
	"log"
	"ouaf/cmd/util"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var UpdateDeviceUserTagsCmd = &cobra.Command{
	Use:   "updatedeviceusertags [device_id]",
	Short: "Update the tags for a device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := util.GetConfig()
		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.UpdateDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0], datadogV2.ListTagsResponse{})
		if err != nil {
			log.Fatalf("failed to updatedeviceusertags: %v", err)
		}

		cmdutil.PrintJSON(res, "network_device_monitoring")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeviceUserTagsCmd)
}
