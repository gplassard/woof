package network_device_monitoring

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListDevicesCmd = &cobra.Command{
	Use: "list-devices",

	Short: "Get the list of devices",
	Long: `Get the list of devices
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#list-devices`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListDevicesResponse
		var err error

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListDevices(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-devices")

		fmt.Println(cmdutil.FormatJSON(res, "device"))
	},
}

func init() {

	Cmd.AddCommand(ListDevicesCmd)
}
