package authn_mappings

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAuthNMappingCmd = &cobra.Command{
	Use:   "getauthnmapping",
	Short: "Get an AuthN Mapping by UUID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/authn_mappings/{authn_mapping_id}")
		fmt.Println("OperationID: GetAuthNMapping")
	},
}

func init() {
	Cmd.AddCommand(GetAuthNMappingCmd)
}
