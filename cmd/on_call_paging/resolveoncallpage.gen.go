package on_call_paging

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ResolveOnCallPageCmd = &cobra.Command{
	Use:   "resolveoncallpage",
	Short: "Resolve On-Call Page",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/pages/{page_id}/resolve")
		fmt.Println("OperationID: ResolveOnCallPage")
	},
}

func init() {
	Cmd.AddCommand(ResolveOnCallPageCmd)
}
