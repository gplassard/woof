package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use:   "listvulnerabilities",
	Short: "List vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/vulnerabilities")
		fmt.Println("OperationID: ListVulnerabilities")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerabilitiesCmd)
}
