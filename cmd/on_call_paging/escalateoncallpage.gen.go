package on_call_paging

import (
	"fmt"
	"github.com/spf13/cobra"
)

var EscalateOnCallPageCmd = &cobra.Command{
	Use:   "escalateoncallpage",
	Short: "Escalate On-Call Page",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/pages/{page_id}/escalate")
		fmt.Println("OperationID: EscalateOnCallPage")
	},
}

func init() {
	Cmd.AddCommand(EscalateOnCallPageCmd)
}
