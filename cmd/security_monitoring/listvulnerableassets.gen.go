package security_monitoring

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListVulnerableAssetsCmd = &cobra.Command{
	Use:   "listvulnerableassets",
	Short: "List vulnerable assets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/security/vulnerable-assets")
		fmt.Println("OperationID: ListVulnerableAssets")
	},
}

func init() {
	Cmd.AddCommand(ListVulnerableAssetsCmd)
}
