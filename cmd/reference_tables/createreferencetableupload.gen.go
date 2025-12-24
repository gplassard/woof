package reference_tables

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateReferenceTableUploadCmd = &cobra.Command{
	Use:   "createreferencetableupload",
	Short: "Create reference table upload",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/reference-tables/uploads")
		fmt.Println("OperationID: CreateReferenceTableUpload")
	},
}

func init() {
	Cmd.AddCommand(CreateReferenceTableUploadCmd)
}
