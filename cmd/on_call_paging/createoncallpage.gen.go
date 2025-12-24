package on_call_paging

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOnCallPageCmd = &cobra.Command{
	Use:   "createoncallpage",
	Short: "Create On-Call Page",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/on-call/pages")
		fmt.Println("OperationID: CreateOnCallPage")
	},
}

func init() {
	Cmd.AddCommand(CreateOnCallPageCmd)
}
