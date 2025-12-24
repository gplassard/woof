package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ListIncidentAttachmentsCmd = &cobra.Command{
	Use:   "listincidentattachments",
	Short: "List incident attachments",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/incidents/{incident_id}/attachments")
		fmt.Println("OperationID: ListIncidentAttachments")
	},
}

func init() {
	Cmd.AddCommand(ListIncidentAttachmentsCmd)
}
