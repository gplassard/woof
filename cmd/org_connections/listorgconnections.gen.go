package org_connections

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListOrgConnectionsCmd = &cobra.Command{
	Use:   "listorgconnections",
	Short: "List Org Connections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/org_connections")
		fmt.Println("OperationID: ListOrgConnections")
	},
}

func init() {
	Cmd.AddCommand(ListOrgConnectionsCmd)
}
