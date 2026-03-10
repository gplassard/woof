package status_pages

import (
	"fmt"
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	"github.com/google/uuid"

	"github.com/spf13/cobra"
)

var CreateDegradationCmd = &cobra.Command{
	Use: "create-degradation [page_id]",

	Short: "Create degradation",
	Long: `Create degradation
Documentation: https://docs.datadoghq.com/api/latest/status-pages/#create-degradation`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.Degradation
		var err error

		var body datadogV2.CreateDegradationRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewStatusPagesApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.CreateDegradation(client.NewContext(apiKey, appKey, site), uuid.MustParse(args[0]), body)
		cmdutil.HandleError(err, "failed to create-degradation")

		fmt.Println(cmdutil.FormatJSON(res, "degradations"))
	},
}

func init() {

	CreateDegradationCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	CreateDegradationCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(CreateDegradationCmd)
}
