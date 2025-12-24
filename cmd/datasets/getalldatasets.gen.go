package datasets

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetAllDatasetsCmd = &cobra.Command{
	Use:   "getalldatasets",
	Short: "Get all datasets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/datasets")
		fmt.Println("OperationID: GetAllDatasets")
	},
}

func init() {
	Cmd.AddCommand(GetAllDatasetsCmd)
}
