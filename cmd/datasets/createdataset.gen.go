package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CreateDatasetCmd = &cobra.Command{
	Use:   "createdataset",
	Short: "Create a dataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: POST /api/v2/datasets")
		fmt.Println("OperationID: CreateDataset")
	},
}

func init() {
	Cmd.AddCommand(CreateDatasetCmd)
}
