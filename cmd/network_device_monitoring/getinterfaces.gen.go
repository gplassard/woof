package network_device_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetInterfacesCmd = &cobra.Command{
	Use:   "getinterfaces",
	Short: "Get the list of interfaces of the device",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ndm/interfaces")
		fmt.Println("OperationID: GetInterfaces")
	},
}

func init() {
	Cmd.AddCommand(GetInterfacesCmd)
}
