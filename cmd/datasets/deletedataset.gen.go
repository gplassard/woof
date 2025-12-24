package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteDatasetCmd = &cobra.Command{
	Use:   "deletedataset",
	Short: "Delete a dataset",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/datasets/{dataset_id}")
		fmt.Println("OperationID: DeleteDataset")
	},
}

func init() {
	Cmd.AddCommand(DeleteDatasetCmd)
}
