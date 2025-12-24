package org_connections

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateOrgConnectionsCmd = &cobra.Command{
	Use:   "updateorgconnections",
	Short: "Update Org Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/org_connections/{connection_id}")
		fmt.Println("OperationID: UpdateOrgConnections")
	},
}

func init() {
	Cmd.AddCommand(UpdateOrgConnectionsCmd)
}
