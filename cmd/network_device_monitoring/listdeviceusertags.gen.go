package network_device_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListDeviceUserTagsCmd = &cobra.Command{
	Use:   "listdeviceusertags",
	Short: "Get the list of tags for a device",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ndm/tags/devices/{device_id}")
		fmt.Println("OperationID: ListDeviceUserTags")
	},
}

func init() {
	Cmd.AddCommand(ListDeviceUserTagsCmd)
}
