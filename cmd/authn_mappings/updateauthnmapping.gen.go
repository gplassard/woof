package authn_mappings

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateAuthNMappingCmd = &cobra.Command{
	Use:   "updateauthnmapping",
	Short: "Edit an AuthN Mapping",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/authn_mappings/{authn_mapping_id}")
		fmt.Println("OperationID: UpdateAuthNMapping")
	},
}

func init() {
	Cmd.AddCommand(UpdateAuthNMappingCmd)
}
