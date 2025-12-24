package container_images

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListContainerImagesCmd = &cobra.Command{
	Use:   "listcontainerimages",
	Short: "Get all Container Images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/container_images")
		fmt.Println("OperationID: ListContainerImages")
	},
}

func init() {
	Cmd.AddCommand(ListContainerImagesCmd)
}
