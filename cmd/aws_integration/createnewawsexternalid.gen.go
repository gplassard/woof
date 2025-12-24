package aws_integration

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateNewAWSExternalIDCmd = &cobra.Command{
	Use:   "createnewawsexternalid",
	Short: "Generate a new external ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/integration/aws/generate_new_external_id")
		fmt.Println("OperationID: CreateNewAWSExternalID")
	},
}

func init() {
	Cmd.AddCommand(CreateNewAWSExternalIDCmd)
}
