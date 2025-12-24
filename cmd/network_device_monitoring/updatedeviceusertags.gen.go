package network_device_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDeviceUserTagsCmd = &cobra.Command{
	Use:   "updatedeviceusertags",
	Short: "Update the tags for a device",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/ndm/tags/devices/{device_id}")
		fmt.Println("OperationID: UpdateDeviceUserTags")
	},
}

func init() {
	Cmd.AddCommand(UpdateDeviceUserTagsCmd)
}
