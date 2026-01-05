package software_catalog

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"

	"encoding/json"
)

var UpsertCatalogEntityCmd = &cobra.Command{
	Use: "upsert-catalog-entity [payload]",

	Short: "Create or update entities",
	Long: `Create or update entities
Documentation: https://docs.datadoghq.com/api/latest/software-catalog/#upsert-catalog-entity`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.UpsertCatalogEntityResponse
		var err error

		var body datadogV2.UpsertCatalogEntityRequest
		err = json.Unmarshal([]byte(args[len(args)-1]), &body)
		cmdutil.HandleError(err, "failed to unmarshal request body")

		api := datadogV2.NewSoftwareCatalogApi(client.NewAPIClient())
		res, _, err = api.UpsertCatalogEntity(client.NewContext(apiKey, appKey, site), body)
		cmdutil.HandleError(err, "failed to upsert-catalog-entity")

		cmd.Println(cmdutil.FormatJSON(res, "software_catalog"))
	},
}

func init() {
	Cmd.AddCommand(UpsertCatalogEntityCmd)
}
