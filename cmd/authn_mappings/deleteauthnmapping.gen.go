package authn_mappings

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteAuthNMappingCmd = &cobra.Command{
	Use:   "deleteauthnmapping",
	Short: "Delete an AuthN Mapping",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/authn_mappings/{authn_mapping_id}")
		fmt.Println("OperationID: DeleteAuthNMapping")
	},
}

func init() {
	Cmd.AddCommand(DeleteAuthNMappingCmd)
}
