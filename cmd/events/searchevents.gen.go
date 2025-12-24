package events

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchEventsCmd = &cobra.Command{
	Use:   "searchevents",
	Short: "Search events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/events/search")
		fmt.Println("OperationID: SearchEvents")
	},
}

func init() {
	Cmd.AddCommand(SearchEventsCmd)
}
