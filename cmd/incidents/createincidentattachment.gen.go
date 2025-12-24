package incidents

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateIncidentAttachmentCmd = &cobra.Command{
	Use:   "createincidentattachment",
	Short: "Create incident attachment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/incidents/{incident_id}/attachments")
		fmt.Println("OperationID: CreateIncidentAttachment")
	},
}

func init() {
	Cmd.AddCommand(CreateIncidentAttachmentCmd)
}
