package authn_mappings

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListAuthNMappingsCmd = &cobra.Command{
	Use:   "listauthnmappings",
	Short: "List all AuthN Mappings",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/authn_mappings")
		fmt.Println("OperationID: ListAuthNMappings")
	},
}

func init() {
	Cmd.AddCommand(ListAuthNMappingsCmd)
}
