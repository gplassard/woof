package containers

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListContainersCmd = &cobra.Command{
	Use:   "listcontainers",
	Short: "Get All Containers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/containers")
		fmt.Println("OperationID: ListContainers")
	},
}

func init() {
	Cmd.AddCommand(ListContainersCmd)
}
