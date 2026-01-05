package network_device_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetInterfacesCmd = &cobra.Command{
	Use: "get-interfaces [device_id]",

	Short: "Get the list of interfaces of the device",
	Long: `Get the list of interfaces of the device
Documentation: https://docs.datadoghq.com/api/latest/network-device-monitoring/#get-interfaces`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.GetInterfacesResponse
		var err error

		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err = api.GetInterfaces(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-interfaces")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(GetInterfacesCmd)
}
