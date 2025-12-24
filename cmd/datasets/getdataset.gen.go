package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetDatasetCmd = &cobra.Command{
	Use:   "getdataset",
	Short: "Get a single dataset by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/datasets/{dataset_id}")
		fmt.Println("OperationID: GetDataset")
	},
}

func init() {
	Cmd.AddCommand(GetDatasetCmd)
}
