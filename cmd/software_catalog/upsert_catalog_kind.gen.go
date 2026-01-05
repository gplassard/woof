package software_catalog

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var UpsertCatalogKindCmd = &cobra.Command{
	Use: "upsert-catalog-kind",

	Short: "Create or update kinds",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.UpsertCatalogKind(client.NewContext(apiKey, appKey, site), datadogV2.UpsertCatalogKindRequest{})
		cmdutil.HandleError(err, "failed to upsert-catalog-kind")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogKindCmd)
}
