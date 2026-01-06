package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertCatalogKindCmd = &cobra.Command{
	Use: "upsert-catalog-kind",

	Short: "Create or update kinds",
	Long: `Create or update kinds
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#upsert-catalog-kind`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpsertCatalogKindResponse
		var err error

		var body datadogV2.UpsertCatalogKindRequest
		err = cmdutil.UnmarshalPayload(cmd, &body)
		cmdutil.HandleError(err, "failed to read payload")

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.UpsertCatalogKind(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to upsert-catalog-kind")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {

	UpsertCatalogKindCmd.Flags().StringP("payload", "p", "", "JSON payload of the request")
	UpsertCatalogKindCmd.Flags().StringP("payload-file", "f", "", "Path to the JSON payload file")

	Cmd.AddCommand(UpsertCatalogKindCmd)
}
