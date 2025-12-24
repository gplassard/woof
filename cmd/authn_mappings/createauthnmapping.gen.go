package authn_mappings

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateAuthNMappingCmd = &cobra.Command{
	Use:   "createauthnmapping",
	Short: "Create an AuthN Mapping",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/authn_mappings")
		fmt.Println("OperationID: CreateAuthNMapping")
	},
}

func init() {
	Cmd.AddCommand(CreateAuthNMappingCmd)
}
