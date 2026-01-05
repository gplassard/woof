package network_device_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetDeviceCmd = &cobra.Command{
	Use: "get-device [device_id]",

	Short: "Get the device details",
	Long: `Get the device details
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#get-device`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetDeviceResponse
		var err error

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetDevice(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-device")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {

	Cmd.AddCommand(GetDeviceCmd)
}
