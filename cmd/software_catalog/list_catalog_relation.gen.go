package software_catalog

import (
	"ouaf/pkg/config"
	"ouaf/pkg/client"
	"ouaf/pkg/cmdutil"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"
	
	
	
	"github.com/spf13/cobra"
	
)

var ListCatalogRelationCmd = &cobra.Command{
	Use:   "list-catalog-relation",
	
	Short: "Get a list of entity relations",
	
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err := api.ListCatalogRelation(client.NewContext(apiKey, appKey, site))
		cmdutil.HandleError(err, "failed to list-catalog-relation")

		cmdutil.PrintJSON(res, "software_catalog")
	},
}

func init() {
	Cmd.AddCommand(ListCatalogRelationCmd)
}
