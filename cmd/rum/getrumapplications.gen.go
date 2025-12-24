package rum

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetRUMApplicationsCmd = &cobra.Command{
	Use:   "getrumapplications",
	Short: "List all the RUM applications",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/rum/applications")
		fmt.Println("OperationID: GetRUMApplications")
	},
}

func init() {
	Cmd.AddCommand(GetRUMApplicationsCmd)
}
