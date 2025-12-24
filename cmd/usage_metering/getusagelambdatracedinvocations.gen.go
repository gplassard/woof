package usage_metering

import (
	"fmt"
	"github.com/spf13/cobra"
)

var GetUsageLambdaTracedInvocationsCmd = &cobra.Command{
	Use:   "getusagelambdatracedinvocations",
	Short: "Get hourly usage for Lambda traced invocations",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Endpoint: GET /api/v2/usage/lambda_traced_invocations")
		fmt.Println("OperationID: GetUsageLambdaTracedInvocations")
	},
}

func init() {
	Cmd.AddCommand(GetUsageLambdaTracedInvocationsCmd)
}
