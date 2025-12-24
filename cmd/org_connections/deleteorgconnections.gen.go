package org_connections

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteOrgConnectionsCmd = &cobra.Command{
	Use:   "deleteorgconnections",
	Short: "Delete Org Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/org_connections/{connection_id}")
		fmt.Println("OperationID: DeleteOrgConnections")
	},
}

func init() {
	Cmd.AddCommand(DeleteOrgConnectionsCmd)
}
