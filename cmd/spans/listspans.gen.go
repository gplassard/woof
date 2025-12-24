package spans

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSpansCmd = &cobra.Command{
	Use:   "listspans",
	Short: "Search spans",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/spans/events/search")
		fmt.Println("OperationID: ListSpans")
	},
}

func init() {
	Cmd.AddCommand(ListSpansCmd)
}
