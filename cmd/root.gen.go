package cmd

import (
	"os"
	"ouaf/cmd/util"
	"ouaf/cmd/action_connection"
	"ouaf/cmd/actions_datastores"
	"ouaf/cmd/agentless_scanning"
	"ouaf/cmd/api_management"
	"ouaf/cmd/apm"
	"ouaf/cmd/apm_retention_filters"
	"ouaf/cmd/app_builder"
	"ouaf/cmd/application_security"
	"ouaf/cmd/audit"
	"ouaf/cmd/authn_mappings"
	"ouaf/cmd/aws_integration"
	"ouaf/cmd/aws_logs_integration"
	"ouaf/cmd/case_management"
	"ouaf/cmd/case_management_attribute"
	"ouaf/cmd/case_management_type"
	"ouaf/cmd/ci_visibility_pipelines"
	"ouaf/cmd/ci_visibility_tests"
	"ouaf/cmd/cloud_cost_management"
	"ouaf/cmd/cloud_network_monitoring"
	"ouaf/cmd/cloudflare_integration"
	"ouaf/cmd/confluent_cloud"
	"ouaf/cmd/container_images"
	"ouaf/cmd/containers"
	"ouaf/cmd/csm_agents"
	"ouaf/cmd/csm_coverage_analysis"
	"ouaf/cmd/csm_threats"
	"ouaf/cmd/dashboard_lists"
	"ouaf/cmd/data_deletion"
	"ouaf/cmd/datasets"
	"ouaf/cmd/deployment_gates"
	"ouaf/cmd/domain_allowlist"
	"ouaf/cmd/dora_metrics"
	"ouaf/cmd/downtimes"
	"ouaf/cmd/error_tracking"
	"ouaf/cmd/events"
	"ouaf/cmd/fastly_integration"
	"ouaf/cmd/fleet_automation"
	"ouaf/cmd/gcp_integration"
	"ouaf/cmd/incident_services"
	"ouaf/cmd/incident_teams"
	"ouaf/cmd/incidents"
	"ouaf/cmd/ip_allowlist"
	"ouaf/cmd/key_management"
	"ouaf/cmd/logs"
	"ouaf/cmd/logs_archives"
	"ouaf/cmd/logs_custom_destinations"
	"ouaf/cmd/logs_metrics"
	"ouaf/cmd/logs_restriction_queries"
	"ouaf/cmd/metrics"
	"ouaf/cmd/microsoft_teams_integration"
	"ouaf/cmd/monitors"
	"ouaf/cmd/network_device_monitoring"
	"ouaf/cmd/observability_pipelines"
	"ouaf/cmd/okta_integration"
	"ouaf/cmd/on_call"
	"ouaf/cmd/on_call_paging"
	"ouaf/cmd/opsgenie_integration"
	"ouaf/cmd/org_connections"
	"ouaf/cmd/organizations"
	"ouaf/cmd/powerpack"
	"ouaf/cmd/processes"
	"ouaf/cmd/reference_tables"
	"ouaf/cmd/restriction_policies"
	"ouaf/cmd/roles"
	"ouaf/cmd/rum"
	"ouaf/cmd/rum_audience_management"
	"ouaf/cmd/rum_metrics"
	"ouaf/cmd/rum_retention_filters"
	"ouaf/cmd/security_monitoring"
	"ouaf/cmd/sensitive_data_scanner"
	"ouaf/cmd/service_accounts"
	"ouaf/cmd/service_definition"
	"ouaf/cmd/service_level_objectives"
	"ouaf/cmd/service_scorecards"
	"ouaf/cmd/software_catalog"
	"ouaf/cmd/spa"
	"ouaf/cmd/spans"
	"ouaf/cmd/spans_metrics"
	"ouaf/cmd/static_analysis"
	"ouaf/cmd/synthetics"
	"ouaf/cmd/teams"
	"ouaf/cmd/test_optimization"
	"ouaf/cmd/usage_metering"
	"ouaf/cmd/users"
	"ouaf/cmd/workflow_automation"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ouaf",
	Short: "ouaf is a Datadog CLI",
	Long:  `ouaf is a Datadog CLI built with Cobra.`,
}

func Execute() {
    defaultSite := os.Getenv("DD_SITE")
	if defaultSite == "" {
		defaultSite = "datadoghq.com"
	}

	rootCmd.PersistentFlags().StringVar(&util.ApiKey, "api-key", os.Getenv("DD_API_KEY"), "Datadog API Key")
	rootCmd.PersistentFlags().StringVar(&util.AppKey, "app-key", os.Getenv("DD_APP_KEY"), "Datadog App Key")
	rootCmd.PersistentFlags().StringVar(&util.Site, "site", defaultSite, "Datadog Site")

	rootCmd.AddCommand(
		action_connection.Cmd,
		actions_datastores.Cmd,
		agentless_scanning.Cmd,
		api_management.Cmd,
		apm.Cmd,
		apm_retention_filters.Cmd,
		app_builder.Cmd,
		application_security.Cmd,
		audit.Cmd,
		authn_mappings.Cmd,
		aws_integration.Cmd,
		aws_logs_integration.Cmd,
		case_management.Cmd,
		case_management_attribute.Cmd,
		case_management_type.Cmd,
		ci_visibility_pipelines.Cmd,
		ci_visibility_tests.Cmd,
		cloud_cost_management.Cmd,
		cloud_network_monitoring.Cmd,
		cloudflare_integration.Cmd,
		confluent_cloud.Cmd,
		container_images.Cmd,
		containers.Cmd,
		csm_agents.Cmd,
		csm_coverage_analysis.Cmd,
		csm_threats.Cmd,
		dashboard_lists.Cmd,
		data_deletion.Cmd,
		datasets.Cmd,
		deployment_gates.Cmd,
		domain_allowlist.Cmd,
		dora_metrics.Cmd,
		downtimes.Cmd,
		error_tracking.Cmd,
		events.Cmd,
		fastly_integration.Cmd,
		fleet_automation.Cmd,
		gcp_integration.Cmd,
		incident_services.Cmd,
		incident_teams.Cmd,
		incidents.Cmd,
		ip_allowlist.Cmd,
		key_management.Cmd,
		logs.Cmd,
		logs_archives.Cmd,
		logs_custom_destinations.Cmd,
		logs_metrics.Cmd,
		logs_restriction_queries.Cmd,
		metrics.Cmd,
		microsoft_teams_integration.Cmd,
		monitors.Cmd,
		network_device_monitoring.Cmd,
		observability_pipelines.Cmd,
		okta_integration.Cmd,
		on_call.Cmd,
		on_call_paging.Cmd,
		opsgenie_integration.Cmd,
		org_connections.Cmd,
		organizations.Cmd,
		powerpack.Cmd,
		processes.Cmd,
		reference_tables.Cmd,
		restriction_policies.Cmd,
		roles.Cmd,
		rum.Cmd,
		rum_audience_management.Cmd,
		rum_metrics.Cmd,
		rum_retention_filters.Cmd,
		security_monitoring.Cmd,
		sensitive_data_scanner.Cmd,
		service_accounts.Cmd,
		service_definition.Cmd,
		service_level_objectives.Cmd,
		service_scorecards.Cmd,
		software_catalog.Cmd,
		spa.Cmd,
		spans.Cmd,
		spans_metrics.Cmd,
		static_analysis.Cmd,
		synthetics.Cmd,
		teams.Cmd,
		test_optimization.Cmd,
		usage_metering.Cmd,
		users.Cmd,
		workflow_automation.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
