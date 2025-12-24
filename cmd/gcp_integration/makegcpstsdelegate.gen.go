package gcp_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var MakeGCPSTSDelegateCmd = &cobra.Command{
	Use:   "makegcpstsdelegate",
	Short: "Create a Datadog GCP principal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/gcp/sts_delegate")
		fmt.Println("OperationID: MakeGCPSTSDelegate")
	},
}

func init() {
	Cmd.AddCommand(MakeGCPSTSDelegateCmd)
}
