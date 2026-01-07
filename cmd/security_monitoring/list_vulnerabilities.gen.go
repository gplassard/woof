package security_monitoring

import (
	"github.com/gplassard/woof/pkg/client"
	"github.com/gplassard/woof/pkg/cmdutil"
	"github.com/gplassard/woof/pkg/config"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadogV2"

	"github.com/spf13/cobra"
)

var ListVulnerabilitiesCmd = &cobra.Command{
	Use: "list-vulnerabilities",

	Short: "List vulnerabilities",
	Long: `List vulnerabilities
Documentation: https://docs.datadoghq.com/api/latest/security-monitoring/#list-vulnerabilities`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey, appKey, site := config.GetConfig()
		var res datadogV2.ListVulnerabilitiesResponse
		var err error

		optionalParams := datadogV2.NewListVulnerabilitiesOptionalParameters()

		if cmd.Flags().Changed("page-token") {
			val, _ := cmd.Flags().GetString("page-token")
			optionalParams.WithPageToken(val)
		}

		if cmd.Flags().Changed("page-number") {
			val, _ := cmd.Flags().GetInt64("page-number")
			optionalParams.WithPageNumber(val)
		}

		if cmd.Flags().Changed("filter-type") {
			val, _ := cmd.Flags().GetString("filter-type")
			optionalParams.WithFilterType(val)
		}

		if cmd.Flags().Changed("filter-cvss.base.score-`$op`") {
			val, _ := cmd.Flags().GetFloat64("filter-cvss.base.score-`$op`")
			optionalParams.WithFilterCvssBaseScoreOp(val)
		}

		if cmd.Flags().Changed("filter-cvss.base.severity") {
			val, _ := cmd.Flags().GetString("filter-cvss.base.severity")
			optionalParams.WithFilterCvssBaseSeverity(val)
		}

		if cmd.Flags().Changed("filter-cvss.base.vector") {
			val, _ := cmd.Flags().GetString("filter-cvss.base.vector")
			optionalParams.WithFilterCvssBaseVector(val)
		}

		if cmd.Flags().Changed("filter-cvss.datadog.score-`$op`") {
			val, _ := cmd.Flags().GetFloat64("filter-cvss.datadog.score-`$op`")
			optionalParams.WithFilterCvssDatadogScoreOp(val)
		}

		if cmd.Flags().Changed("filter-cvss.datadog.severity") {
			val, _ := cmd.Flags().GetString("filter-cvss.datadog.severity")
			optionalParams.WithFilterCvssDatadogSeverity(val)
		}

		if cmd.Flags().Changed("filter-cvss.datadog.vector") {
			val, _ := cmd.Flags().GetString("filter-cvss.datadog.vector")
			optionalParams.WithFilterCvssDatadogVector(val)
		}

		if cmd.Flags().Changed("filter-status") {
			val, _ := cmd.Flags().GetString("filter-status")
			optionalParams.WithFilterStatus(val)
		}

		if cmd.Flags().Changed("filter-tool") {
			val, _ := cmd.Flags().GetString("filter-tool")
			optionalParams.WithFilterTool(val)
		}

		if cmd.Flags().Changed("filter-library.name") {
			val, _ := cmd.Flags().GetString("filter-library.name")
			optionalParams.WithFilterLibraryName(val)
		}

		if cmd.Flags().Changed("filter-library.version") {
			val, _ := cmd.Flags().GetString("filter-library.version")
			optionalParams.WithFilterLibraryVersion(val)
		}

		if cmd.Flags().Changed("filter-advisory.id") {
			val, _ := cmd.Flags().GetString("filter-advisory.id")
			optionalParams.WithFilterAdvisoryId(val)
		}

		if cmd.Flags().Changed("filter-risks.exploitation-probability") {
			val, _ := cmd.Flags().GetString("filter-risks.exploitation-probability")
			optionalParams.WithFilterRisksExploitationProbability(val)
		}

		if cmd.Flags().Changed("filter-risks.poc-exploit-available") {
			val, _ := cmd.Flags().GetString("filter-risks.poc-exploit-available")
			optionalParams.WithFilterRisksPocExploitAvailable(val)
		}

		if cmd.Flags().Changed("filter-risks.exploit-available") {
			val, _ := cmd.Flags().GetString("filter-risks.exploit-available")
			optionalParams.WithFilterRisksExploitAvailable(val)
		}

		if cmd.Flags().Changed("filter-risks.epss.score-`$op`") {
			val, _ := cmd.Flags().GetFloat64("filter-risks.epss.score-`$op`")
			optionalParams.WithFilterRisksEpssScoreOp(val)
		}

		if cmd.Flags().Changed("filter-risks.epss.severity") {
			val, _ := cmd.Flags().GetString("filter-risks.epss.severity")
			optionalParams.WithFilterRisksEpssSeverity(val)
		}

		if cmd.Flags().Changed("filter-language") {
			val, _ := cmd.Flags().GetString("filter-language")
			optionalParams.WithFilterLanguage(val)
		}

		if cmd.Flags().Changed("filter-ecosystem") {
			val, _ := cmd.Flags().GetString("filter-ecosystem")
			optionalParams.WithFilterEcosystem(val)
		}

		if cmd.Flags().Changed("filter-code-location.location") {
			val, _ := cmd.Flags().GetString("filter-code-location.location")
			optionalParams.WithFilterCodeLocationLocation(val)
		}

		if cmd.Flags().Changed("filter-code-location.file-path") {
			val, _ := cmd.Flags().GetString("filter-code-location.file-path")
			optionalParams.WithFilterCodeLocationFilePath(val)
		}

		if cmd.Flags().Changed("filter-code-location.method") {
			val, _ := cmd.Flags().GetString("filter-code-location.method")
			optionalParams.WithFilterCodeLocationMethod(val)
		}

		if cmd.Flags().Changed("filter-fix-available") {
			val, _ := cmd.Flags().GetString("filter-fix-available")
			optionalParams.WithFilterFixAvailable(val)
		}

		if cmd.Flags().Changed("filter-repo-digests") {
			val, _ := cmd.Flags().GetString("filter-repo-digests")
			optionalParams.WithFilterRepoDigests(val)
		}

		if cmd.Flags().Changed("filter-origin") {
			val, _ := cmd.Flags().GetString("filter-origin")
			optionalParams.WithFilterOrigin(val)
		}

		if cmd.Flags().Changed("filter-running-kernel") {
			val, _ := cmd.Flags().GetString("filter-running-kernel")
			optionalParams.WithFilterRunningKernel(val)
		}

		if cmd.Flags().Changed("filter-asset.name") {
			val, _ := cmd.Flags().GetString("filter-asset.name")
			optionalParams.WithFilterAssetName(val)
		}

		if cmd.Flags().Changed("filter-asset.type") {
			val, _ := cmd.Flags().GetString("filter-asset.type")
			optionalParams.WithFilterAssetType(val)
		}

		if cmd.Flags().Changed("filter-asset.version.first") {
			val, _ := cmd.Flags().GetString("filter-asset.version.first")
			optionalParams.WithFilterAssetVersionFirst(val)
		}

		if cmd.Flags().Changed("filter-asset.version.last") {
			val, _ := cmd.Flags().GetString("filter-asset.version.last")
			optionalParams.WithFilterAssetVersionLast(val)
		}

		if cmd.Flags().Changed("filter-asset.repository-url") {
			val, _ := cmd.Flags().GetString("filter-asset.repository-url")
			optionalParams.WithFilterAssetRepositoryUrl(val)
		}

		if cmd.Flags().Changed("filter-asset.risks.in-production") {
			val, _ := cmd.Flags().GetString("filter-asset.risks.in-production")
			optionalParams.WithFilterAssetRisksInProduction(val)
		}

		if cmd.Flags().Changed("filter-asset.risks.under-attack") {
			val, _ := cmd.Flags().GetString("filter-asset.risks.under-attack")
			optionalParams.WithFilterAssetRisksUnderAttack(val)
		}

		if cmd.Flags().Changed("filter-asset.risks.is-publicly-accessible") {
			val, _ := cmd.Flags().GetString("filter-asset.risks.is-publicly-accessible")
			optionalParams.WithFilterAssetRisksIsPubliclyAccessible(val)
		}

		if cmd.Flags().Changed("filter-asset.risks.has-privileged-access") {
			val, _ := cmd.Flags().GetString("filter-asset.risks.has-privileged-access")
			optionalParams.WithFilterAssetRisksHasPrivilegedAccess(val)
		}

		if cmd.Flags().Changed("filter-asset.risks.has-access-to-sensitive-data") {
			val, _ := cmd.Flags().GetString("filter-asset.risks.has-access-to-sensitive-data")
			optionalParams.WithFilterAssetRisksHasAccessToSensitiveData(val)
		}

		if cmd.Flags().Changed("filter-asset.environments") {
			val, _ := cmd.Flags().GetString("filter-asset.environments")
			optionalParams.WithFilterAssetEnvironments(val)
		}

		if cmd.Flags().Changed("filter-asset.teams") {
			val, _ := cmd.Flags().GetString("filter-asset.teams")
			optionalParams.WithFilterAssetTeams(val)
		}

		if cmd.Flags().Changed("filter-asset.arch") {
			val, _ := cmd.Flags().GetString("filter-asset.arch")
			optionalParams.WithFilterAssetArch(val)
		}

		if cmd.Flags().Changed("filter-asset.operating-system.name") {
			val, _ := cmd.Flags().GetString("filter-asset.operating-system.name")
			optionalParams.WithFilterAssetOperatingSystemName(val)
		}

		if cmd.Flags().Changed("filter-asset.operating-system.version") {
			val, _ := cmd.Flags().GetString("filter-asset.operating-system.version")
			optionalParams.WithFilterAssetOperatingSystemVersion(val)
		}

		api := datadogV2.NewSecurityMonitoringApi(client.NewAPIClient())
		//nolint:staticcheck // SA1019: deprecated
		res, _, err = api.ListVulnerabilities(client.NewContext(apiKey, appKey, site), *optionalParams)
		cmdutil.HandleError(err, "failed to list-vulnerabilities")

		cmd.Println(cmdutil.FormatJSON(res, "vulnerabilities"))
	},
}

func init() {

	ListVulnerabilitiesCmd.Flags().String("page-token", "", "Its value must come from the 'links' section of the response of the first request. Do not manually edit it.")

	ListVulnerabilitiesCmd.Flags().Int64("page-number", 0, "The page number to be retrieved. It should be equal or greater than '1'")

	ListVulnerabilitiesCmd.Flags().String("filter-type", "", "Filter by vulnerability type.")

	ListVulnerabilitiesCmd.Flags().Float64("filter-cvss.base.score-`$op`", 0, "Filter by vulnerability base (i.e. from the original advisory) severity score.")

	ListVulnerabilitiesCmd.Flags().String("filter-cvss.base.severity", "", "Filter by vulnerability base severity.")

	ListVulnerabilitiesCmd.Flags().String("filter-cvss.base.vector", "", "Filter by vulnerability base CVSS vector.")

	ListVulnerabilitiesCmd.Flags().Float64("filter-cvss.datadog.score-`$op`", 0, "Filter by vulnerability Datadog severity score.")

	ListVulnerabilitiesCmd.Flags().String("filter-cvss.datadog.severity", "", "Filter by vulnerability Datadog severity.")

	ListVulnerabilitiesCmd.Flags().String("filter-cvss.datadog.vector", "", "Filter by vulnerability Datadog CVSS vector.")

	ListVulnerabilitiesCmd.Flags().String("filter-status", "", "Filter by the status of the vulnerability.")

	ListVulnerabilitiesCmd.Flags().String("filter-tool", "", "Filter by the tool of the vulnerability.")

	ListVulnerabilitiesCmd.Flags().String("filter-library.name", "", "Filter by library name.")

	ListVulnerabilitiesCmd.Flags().String("filter-library.version", "", "Filter by library version.")

	ListVulnerabilitiesCmd.Flags().String("filter-advisory.id", "", "Filter by advisory ID.")

	ListVulnerabilitiesCmd.Flags().String("filter-risks.exploitation-probability", "", "Filter by exploitation probability.")

	ListVulnerabilitiesCmd.Flags().String("filter-risks.poc-exploit-available", "", "Filter by POC exploit availability.")

	ListVulnerabilitiesCmd.Flags().String("filter-risks.exploit-available", "", "Filter by public exploit availability.")

	ListVulnerabilitiesCmd.Flags().Float64("filter-risks.epss.score-`$op`", 0, "Filter by vulnerability [EPSS](https://www.first.org/epss/) severity score.")

	ListVulnerabilitiesCmd.Flags().String("filter-risks.epss.severity", "", "Filter by vulnerability [EPSS](https://www.first.org/epss/) severity.")

	ListVulnerabilitiesCmd.Flags().String("filter-language", "", "Filter by language.")

	ListVulnerabilitiesCmd.Flags().String("filter-ecosystem", "", "Filter by ecosystem.")

	ListVulnerabilitiesCmd.Flags().String("filter-code-location.location", "", "Filter by vulnerability location.")

	ListVulnerabilitiesCmd.Flags().String("filter-code-location.file-path", "", "Filter by vulnerability file path.")

	ListVulnerabilitiesCmd.Flags().String("filter-code-location.method", "", "Filter by method.")

	ListVulnerabilitiesCmd.Flags().String("filter-fix-available", "", "Filter by fix availability.")

	ListVulnerabilitiesCmd.Flags().String("filter-repo-digests", "", "Filter by vulnerability 'repo_digest' (when the vulnerability is related to 'Image' asset).")

	ListVulnerabilitiesCmd.Flags().String("filter-origin", "", "Filter by origin.")

	ListVulnerabilitiesCmd.Flags().String("filter-running-kernel", "", "Filter for whether the vulnerability affects a running kernel (for vulnerabilities related to a 'Host' asset).")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.name", "", "Filter by asset name. This field supports the usage of wildcards (*).")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.type", "", "Filter by asset type.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.version.first", "", "Filter by the first version of the asset this vulnerability has been detected on.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.version.last", "", "Filter by the last version of the asset this vulnerability has been detected on.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.repository-url", "", "Filter by the repository url associated to the asset.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.risks.in-production", "", "Filter whether the asset is in production or not.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.risks.under-attack", "", "Filter whether the asset is under attack or not.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.risks.is-publicly-accessible", "", "Filter whether the asset is publicly accessible or not.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.risks.has-privileged-access", "", "Filter whether the asset is publicly accessible or not.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.risks.has-access-to-sensitive-data", "", "Filter whether the asset  has access to sensitive data or not.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.environments", "", "Filter by asset environments.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.teams", "", "Filter by asset teams.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.arch", "", "Filter by asset architecture.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.operating-system.name", "", "Filter by asset operating system name.")

	ListVulnerabilitiesCmd.Flags().String("filter-asset.operating-system.version", "", "Filter by asset operating system version.")

	Cmd.AddCommand(ListVulnerabilitiesCmd)
}
