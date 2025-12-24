package on_call_paging

import (
	"fmt"
	"github.com/spf13/cobra"
)

var AcknowledgeOnCallPageCmd = &cobra.Command{
	Use:   "acknowledgeoncallpage",
	Short: "Acknowledge On-Call Page",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/pages/{page_id}/acknowledge")
		fmt.Println("OperationID: AcknowledgeOnCallPage")
	},
}

func init() {
	Cmd.AddCommand(AcknowledgeOnCallPageCmd)
}
