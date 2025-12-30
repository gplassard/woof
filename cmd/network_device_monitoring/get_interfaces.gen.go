package network_device_monitoring

import (
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"
	"ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var GetInterfacesCmd = &cobra.Command{
	Use: "get-interfaces [device_id]",

	Short: "Get the list of interfaces of the device",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewNetworkDeviceMonitoringApi(client.NewAPIClient())
		res, _, err := api.GetInterfaces(client.NewContext(apiKey, appKey, site), args[0])
		cmdutil.HandleError(err, "failed to get-interfaces")

		cmd.Println(cmdutil.FormatJSON(res, "network_device_monitoring"))
	},
}

func init() {
	Cmd.AddCommand(GetInterfacesCmd)
}
