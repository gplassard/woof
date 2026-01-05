package software_catalog

import (
	"github.com/gplassard/ouaf/pkg/client"
	"github.com/gplassard/ouaf/pkg/cmdutil"
	"github.com/gplassard/ouaf/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListCatalogKindCmd = &cobra.Command{
	Use: "list-catalog-kind",

	Short: "Get a list of entity kinds",

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.ListCatalogKind(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-catalog-kind")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {
	Cmd.AddCommand(ListCatalogKindCmd)
}
