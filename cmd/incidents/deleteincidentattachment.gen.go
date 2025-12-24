package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteIncidentAttachmentCmd = &cobra.Command{
	Use:   "deleteincidentattachment",
	Short: "Delete incident attachment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/incidents/{incident_id}/attachments/{attachment_id}")
		fmt.Println("OperationID: DeleteIncidentAttachment")
	},
}

func init() {
	Cmd.AddCommand(DeleteIncidentAttachmentCmd)
}
