package metrics

import (
	"fmt"
	"github.com/spf13/cobra"
)

var DeleteTagConfigurationCmd = &cobra.Command{
	Use:   "deletetagconfiguration",
	Short: "Delete a tag configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: DELETE /api/v2/metrics/{metric_name}/tags")
		fmt.Println("OperationID: DeleteTagConfiguration")
	},
}

func init() {
	Cmd.AddCommand(DeleteTagConfigurationCmd)
}
