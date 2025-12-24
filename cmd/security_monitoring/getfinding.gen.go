package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetFindingCmd = &cobra.Command{
	Use:   "getfinding",
	Short: "Get a finding",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/posture_management/findings/{finding_id}")
		fmt.Println("OperationID: GetFinding")
	},
}

func init() {
	Cmd.AddCommand(GetFindingCmd)
}
