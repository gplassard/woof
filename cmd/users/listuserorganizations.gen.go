package users

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListUserOrganizationsCmd = &cobra.Command{
	Use:   "listuserorganizations",
	Short: "Get a user organization",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/users/{user_id}/orgs")
		fmt.Println("OperationID: ListUserOrganizations")
	},
}

func init() {
	Cmd.AddCommand(ListUserOrganizationsCmd)
}
