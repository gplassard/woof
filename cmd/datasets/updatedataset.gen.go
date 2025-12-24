package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

var UpdateDatasetCmd = &cobra.Command{
	Use:   "updatedataset",
	Short: "Edit a dataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: PUT /api/v2/datasets/{dataset_id}")
		fmt.Println("OperationID: UpdateDataset")
	},
}

func init() {
	Cmd.AddCommand(UpdateDatasetCmd)
}
