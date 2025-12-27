package network_device_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDeviceUserTagsCmd = &cobra.Command{
	Use:   "listdeviceusertags [device_id]",
	Short: "Get the list of tags for a device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListDeviceUserTags(client.NewContext(apiKey, appKey, site), args[0])
		if err != nil {
			log.Fatalf("failed to listdeviceusertags: %v", err)
		}

		cmdutil.PrintJSON(res, "network_device_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListDeviceUserTagsCmd)
}
