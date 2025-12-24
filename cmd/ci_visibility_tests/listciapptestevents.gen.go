package ci_visibility_tests

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListCIAppTestEventsCmd = &cobra.Command{
	Use:   "listciapptestevents",
	Short: "Get a list of tests events",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/ci/tests/events")
		fmt.Println("OperationID: ListCIAppTestEvents")
	},
}

func init() {
	Cmd.AddCommand(ListCIAppTestEventsCmd)
}
