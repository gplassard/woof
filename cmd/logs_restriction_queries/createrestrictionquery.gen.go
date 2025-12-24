package logs_restriction_queries

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateRestrictionQueryCmd = &cobra.Command{
	Use:   "createrestrictionquery",
	Short: "Create a restriction query",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/logs/config/restriction_queries")
		fmt.Println("OperationID: CreateRestrictionQuery")
	},
}

func init() {
	Cmd.AddCommand(CreateRestrictionQueryCmd)
}
