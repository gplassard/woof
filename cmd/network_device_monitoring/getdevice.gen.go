package network_device_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDeviceCmd = &cobra.Command{
	Use:   "getdevice",
	Short: "Get the device details",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ndm/devices/{device_id}")
		fmt.Println("OperationID: GetDevice")
	},
}

func init() {
	Cmd.AddCommand(GetDeviceCmd)
}
