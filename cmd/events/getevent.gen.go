package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetEventCmd = &cobra.Command{
	Use:   "getevent",
	Short: "Get an event",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/events/{event_id}")
		fmt.Println("OperationID: GetEvent")
	},
}

func init() {
	Cmd.AddCommand(GetEventCmd)
}
