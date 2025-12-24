package spans

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListSpansGetCmd = &cobra.Command{
	Use:   "listspansget",
	Short: "Get a list of spans",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/spans/events")
		fmt.Println("OperationID: ListSpansGet")
	},
}

func init() {
	Cmd.AddCommand(ListSpansGetCmd)
}
