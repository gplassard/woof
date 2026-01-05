package network_device_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDevicesCmd = &cobra.Command{
	Use: "list-devices",

	Short: "Get the list of devices",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.ListDevices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-devices")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(ListDevicesCmd)
}
