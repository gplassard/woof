package network_device_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDevicesCmd = &cobra.Command{
	Use:   "listdevices",
	Short: "Get the list of devices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ndm/devices")
		fmt.Println("OperationID: ListDevices")
	},
}

func init() {
	Cmd.AddCommand(ListDevicesCmd)
}
