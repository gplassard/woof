package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateEventCmd = &cobra.Command{
	Use:   "createevent",
	Short: "Post an event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/events")
		fmt.Println("OperationID: CreateEvent")
	},
}

func init() {
	Cmd.AddCommand(CreateEventCmd)
}
