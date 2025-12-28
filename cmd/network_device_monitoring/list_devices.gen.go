package network_device_monitoring

import (
	"log"
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Get the list of devices",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListDevices(client.NewContext(apiKey, appKey, site))
		if err != nil {
			log.Fatalf("failed to list-devices: %v", err)
		}

		cmdutil.PrintJSON(res, "network_device_monitoring")
	},
}

func init() {
	Cmd.AddCommand(ListDevicesCmd)
}
