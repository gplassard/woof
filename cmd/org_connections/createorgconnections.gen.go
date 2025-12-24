package org_connections

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateOrgConnectionsCmd = &cobra.Command{
	Use:   "createorgconnections",
	Short: "Create Org Connection",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/org_connections")
		fmt.Println("OperationID: CreateOrgConnections")
	},
}

func init() {
	Cmd.AddCommand(CreateOrgConnectionsCmd)
}
