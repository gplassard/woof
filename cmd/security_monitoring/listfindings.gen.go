package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListFindingsCmd = &cobra.Command{
	Use:   "listfindings",
	Short: "List findings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/posture_management/findings")
		fmt.Println("OperationID: ListFindings")
	},
}

func init() {
	Cmd.AddCommand(ListFindingsCmd)
}
