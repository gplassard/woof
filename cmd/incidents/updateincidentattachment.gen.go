package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateIncidentAttachmentCmd = &cobra.Command{
	Use:   "updateincidentattachment",
	Short: "Update incident attachment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PATCH /api/v2/incidents/{incident_id}/attachments/{attachment_id}")
		fmt.Println("OperationID: UpdateIncidentAttachment")
	},
}

func init() {
	Cmd.AddCommand(UpdateIncidentAttachmentCmd)
}
