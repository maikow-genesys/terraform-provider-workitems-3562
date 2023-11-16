package genesyscloud

import (
	registrar "terraform-provider-genesyscloud/genesyscloud/resource_register"
)

func SetRegistrar(l registrar.Registrar) {

	registerDataSources(l)
	registerResources(l)
	registerExporters(l)
}

func registerDataSources(l registrar.Registrar) {

	l.RegisterDataSource("genesyscloud_routing_wrapupcode", DataSourceRoutingWrapupcode())
	l.RegisterDataSource("genesyscloud_routing_queue", DataSourceRoutingQueue())
	l.RegisterDataSource("genesyscloud_flow", DataSourceFlow())
	l.RegisterDataSource("genesyscloud_location", DataSourceLocation())
	l.RegisterDataSource("genesyscloud_auth_division_home", DataSourceAuthDivisionHome())
	l.RegisterDataSource("genesyscloud_architect_datatable", DataSourceArchitectDatatable())
	l.RegisterDataSource("genesyscloud_architect_emergencygroup", DataSourceArchitectEmergencyGroup())
	l.RegisterDataSource("genesyscloud_architect_schedules", DataSourceSchedule())
	l.RegisterDataSource("genesyscloud_architect_schedulegroups", DataSourceArchitectScheduleGroups())
	l.RegisterDataSource("genesyscloud_architect_user_prompt", dataSourceUserPrompt())
	l.RegisterDataSource("genesyscloud_auth_role", DataSourceAuthRole())
	l.RegisterDataSource("genesyscloud_auth_division", dataSourceAuthDivision())
	l.RegisterDataSource("genesyscloud_auth_division_home", DataSourceAuthDivisionHome())
	l.RegisterDataSource("genesyscloud_employeeperformance_externalmetrics_definitions", dataSourceEmployeeperformanceExternalmetricsDefinition())
	l.RegisterDataSource("genesyscloud_flow", DataSourceFlow())
	l.RegisterDataSource("genesyscloud_group", DataSourceGroup())
	l.RegisterDataSource("genesyscloud_journey_action_map", dataSourceJourneyActionMap())
	l.RegisterDataSource("genesyscloud_journey_action_template", dataSourceJourneyActionTemplate())
	l.RegisterDataSource("genesyscloud_journey_outcome", dataSourceJourneyOutcome())
	l.RegisterDataSource("genesyscloud_journey_segment", dataSourceJourneySegment())
	l.RegisterDataSource("genesyscloud_knowledge_knowledgebase", dataSourceKnowledgeKnowledgebase())
	l.RegisterDataSource("genesyscloud_knowledge_category", dataSourceKnowledgeCategory())
	l.RegisterDataSource("genesyscloud_knowledge_label", dataSourceKnowledgeLabel())
	l.RegisterDataSource("genesyscloud_location", DataSourceLocation())
	l.RegisterDataSource("genesyscloud_oauth_client", dataSourceOAuthClient())
	l.RegisterDataSource("genesyscloud_organizations_me", DataSourceOrganizationsMe())
	l.RegisterDataSource("genesyscloud_quality_forms_evaluation", DataSourceQualityFormsEvaluations())
	l.RegisterDataSource("genesyscloud_quality_forms_survey", dataSourceQualityFormsSurvey())
	l.RegisterDataSource("genesyscloud_responsemanagement_library", dataSourceResponsemanagementLibrary())
	l.RegisterDataSource("genesyscloud_responsemanagement_response", dataSourceResponsemanagementResponse())
	l.RegisterDataSource("genesyscloud_responsemanagement_responseasset", dataSourceResponseManagamentResponseAsset())
	l.RegisterDataSource("genesyscloud_routing_language", dataSourceRoutingLanguage())
	l.RegisterDataSource("genesyscloud_routing_queue", DataSourceRoutingQueue())
	l.RegisterDataSource("genesyscloud_routing_settings", dataSourceRoutingSettings())
	l.RegisterDataSource("genesyscloud_routing_skill", dataSourceRoutingSkill())
	l.RegisterDataSource("genesyscloud_routing_skill_group", dataSourceRoutingSkillGroup())
	l.RegisterDataSource("genesyscloud_routing_email_domain", DataSourceRoutingEmailDomain())
	l.RegisterDataSource("genesyscloud_routing_wrapupcode", DataSourceRoutingWrapupcode())
	l.RegisterResource("genesyscloud_routing_wrapupcode", ResourceRoutingWrapupCode())

	l.RegisterDataSource("genesyscloud_user", DataSourceUser())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_edge_group", dataSourceEdgeGroup())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_extension_pool", dataSourceExtensionPool())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_linebasesettings", dataSourceLineBaseSettings())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_phonebasesettings", dataSourcePhoneBaseSettings())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_trunk", dataSourceTrunk())
	l.RegisterDataSource("genesyscloud_telephony_providers_edges_trunkbasesettings", dataSourceTrunkBaseSettings())
	l.RegisterDataSource("genesyscloud_webdeployments_deployment", dataSourceWebDeploymentsDeployment())
	l.RegisterDataSource("genesyscloud_widget_deployment", dataSourceWidgetDeployments())
}

func registerResources(l registrar.Registrar) {

	l.RegisterResource("genesyscloud_routing_wrapupcode", ResourceRoutingWrapupCode())
	l.RegisterResource("genesyscloud_routing_queue", ResourceRoutingQueue())
	l.RegisterResource("genesyscloud_flow", ResourceFlow())
	l.RegisterResource("genesyscloud_location", ResourceLocation())
	l.RegisterResource("genesyscloud_architect_datatable", ResourceArchitectDatatable())
	l.RegisterResource("genesyscloud_architect_datatable_row", ResourceArchitectDatatableRow())
	l.RegisterResource("genesyscloud_architect_emergencygroup", ResourceArchitectEmergencyGroup())
	l.RegisterResource("genesyscloud_flow", ResourceFlow())
	l.RegisterResource("genesyscloud_architect_schedules", ResourceArchitectSchedules())
	l.RegisterResource("genesyscloud_architect_schedulegroups", ResourceArchitectScheduleGroups())
	l.RegisterResource("genesyscloud_architect_user_prompt", ResourceArchitectUserPrompt())
	l.RegisterResource("genesyscloud_auth_role", ResourceAuthRole())
	l.RegisterResource("genesyscloud_auth_division", ResourceAuthDivision())
	l.RegisterResource("genesyscloud_employeeperformance_externalmetrics_definitions", ResourceEmployeeperformanceExternalmetricsDefinition())
	l.RegisterResource("genesyscloud_group", ResourceGroup())
	l.RegisterResource("genesyscloud_group_roles", ResourceGroupRoles())
	l.RegisterResource("genesyscloud_idp_adfs", ResourceIdpAdfs())
	l.RegisterResource("genesyscloud_idp_generic", ResourceIdpGeneric())
	l.RegisterResource("genesyscloud_idp_gsuite", ResourceIdpGsuite())
	l.RegisterResource("genesyscloud_idp_okta", ResourceIdpOkta())
	l.RegisterResource("genesyscloud_idp_onelogin", ResourceIdpOnelogin())
	l.RegisterResource("genesyscloud_idp_ping", ResourceIdpPing())
	l.RegisterResource("genesyscloud_idp_salesforce", ResourceIdpSalesforce())
	l.RegisterResource("genesyscloud_journey_action_map", ResourceJourneyActionMap())
	l.RegisterResource("genesyscloud_journey_action_template", ResourceJourneyActionTemplate())
	l.RegisterResource("genesyscloud_journey_outcome", ResourceJourneyOutcome())
	l.RegisterResource("genesyscloud_journey_segment", ResourceJourneySegment())
	l.RegisterResource("genesyscloud_knowledge_knowledgebase", ResourceKnowledgeKnowledgebase())
	l.RegisterResource("genesyscloud_knowledge_document", ResourceKnowledgeDocument())
	l.RegisterResource("genesyscloud_knowledge_v1_document", ResourceKnowledgeDocumentV1())
	l.RegisterResource("genesyscloud_knowledge_document_variation", ResourceKnowledgeDocumentVariation())
	l.RegisterResource("genesyscloud_knowledge_category", ResourceKnowledgeCategory())
	l.RegisterResource("genesyscloud_knowledge_v1_category", ResourceKnowledgeCategoryV1())
	l.RegisterResource("genesyscloud_knowledge_label", ResourceKnowledgeLabel())
	l.RegisterResource("genesyscloud_location", ResourceLocation())
	l.RegisterResource("genesyscloud_oauth_client", ResourceOAuthClient())

	l.RegisterResource("genesyscloud_orgauthorization_pairing", resourceOrgauthorizationPairing())
	l.RegisterResource("genesyscloud_quality_forms_evaluation", ResourceEvaluationForm())
	l.RegisterResource("genesyscloud_quality_forms_survey", ResourceSurveyForm())
	l.RegisterResource("genesyscloud_responsemanagement_library", ResourceResponsemanagementLibrary())
	l.RegisterResource("genesyscloud_responsemanagement_response", ResourceResponsemanagementResponse())
	l.RegisterResource("genesyscloud_responsemanagement_responseasset", resourceResponseManagamentResponseAsset())
	l.RegisterResource("genesyscloud_routing_email_domain", ResourceRoutingEmailDomain())
	l.RegisterResource("genesyscloud_routing_email_route", ResourceRoutingEmailRoute())
	l.RegisterResource("genesyscloud_routing_language", ResourceRoutingLanguage())
	l.RegisterResource("genesyscloud_routing_queue", ResourceRoutingQueue())
	l.RegisterResource("genesyscloud_routing_skill", ResourceRoutingSkill())
	l.RegisterResource("genesyscloud_routing_skill_group", ResourceRoutingSkillGroup())
	l.RegisterResource("genesyscloud_routing_settings", ResourceRoutingSettings())
	l.RegisterResource("genesyscloud_routing_utilization", ResourceRoutingUtilization())
	l.RegisterResource("genesyscloud_routing_wrapupcode", ResourceRoutingWrapupCode())
	l.RegisterResource("genesyscloud_telephony_providers_edges_edge_group", ResourceEdgeGroup())
	l.RegisterResource("genesyscloud_telephony_providers_edges_extension_pool", ResourceTelephonyExtensionPool())
	l.RegisterResource("genesyscloud_telephony_providers_edges_phonebasesettings", ResourcePhoneBaseSettings())
	l.RegisterResource("genesyscloud_telephony_providers_edges_trunkbasesettings", ResourceTrunkBaseSettings())
	l.RegisterResource("genesyscloud_telephony_providers_edges_trunk", ResourceTrunk())
	l.RegisterResource("genesyscloud_user", ResourceUser())
	l.RegisterResource("genesyscloud_user_roles", ResourceUserRoles())
	l.RegisterResource("genesyscloud_webdeployments_deployment", ResourceWebDeployment())
	l.RegisterResource("genesyscloud_widget_deployment", ResourceWidgetDeployment())

}

func registerExporters(l registrar.Registrar) {
	l.RegisterExporter("genesyscloud_architect_datatable", ArchitectDatatableExporter())
	l.RegisterExporter("genesyscloud_architect_datatable_row", ArchitectDatatableRowExporter())
	l.RegisterExporter("genesyscloud_architect_emergencygroup", ArchitectEmergencyGroupExporter())
	l.RegisterExporter("genesyscloud_architect_schedules", ArchitectSchedulesExporter())
	l.RegisterExporter("genesyscloud_architect_schedulegroups", ArchitectScheduleGroupsExporter())
	l.RegisterExporter("genesyscloud_architect_user_prompt", ArchitectUserPromptExporter())
	l.RegisterExporter("genesyscloud_auth_division", AuthDivisionExporter())
	l.RegisterExporter("genesyscloud_auth_role", AuthRoleExporter())
	l.RegisterExporter("genesyscloud_employeeperformance_externalmetrics_definitions", EmployeeperformanceExternalmetricsDefinitionExporter())
	l.RegisterExporter("genesyscloud_flow", FlowExporter())
	l.RegisterExporter("genesyscloud_group", GroupExporter())
	l.RegisterExporter("genesyscloud_group_roles", GroupRolesExporter())
	l.RegisterExporter("genesyscloud_idp_adfs", IdpAdfsExporter())
	l.RegisterExporter("genesyscloud_idp_generic", IdpGenericExporter())
	l.RegisterExporter("genesyscloud_idp_gsuite", IdpGsuiteExporter())
	l.RegisterExporter("genesyscloud_idp_okta", IdpOktaExporter())
	l.RegisterExporter("genesyscloud_idp_onelogin", IdpOneloginExporter())
	l.RegisterExporter("genesyscloud_idp_ping", IdpPingExporter())
	l.RegisterExporter("genesyscloud_idp_salesforce", IdpSalesforceExporter())
	l.RegisterExporter("genesyscloud_journey_action_map", JourneyActionMapExporter())
	l.RegisterExporter("genesyscloud_journey_action_template", JourneyActionTemplateExporter())
	l.RegisterExporter("genesyscloud_journey_outcome", JourneyOutcomeExporter())
	l.RegisterExporter("genesyscloud_journey_segment", JourneySegmentExporter())
	l.RegisterExporter("genesyscloud_knowledge_knowledgebase", KnowledgeKnowledgebaseExporter())
	l.RegisterExporter("genesyscloud_knowledge_document", KnowledgeDocumentExporter())
	l.RegisterExporter("genesyscloud_knowledge_category", KnowledgeCategoryExporter())
	l.RegisterExporter("genesyscloud_location", LocationExporter())
	l.RegisterExporter("genesyscloud_oauth_client", OauthClientExporter())
	l.RegisterExporter("genesyscloud_quality_forms_evaluation", EvaluationFormExporter())
	l.RegisterExporter("genesyscloud_quality_forms_survey", SurveyFormExporter())
	l.RegisterExporter("genesyscloud_responsemanagement_library", ResponsemanagementLibraryExporter())
	l.RegisterExporter("genesyscloud_responsemanagement_response", ResponsemanagementResponseExporter())
	l.RegisterExporter("genesyscloud_routing_email_domain", RoutingEmailDomainExporter())
	l.RegisterExporter("genesyscloud_routing_email_route", RoutingEmailRouteExporter())
	l.RegisterExporter("genesyscloud_routing_language", RoutingLanguageExporter())
	l.RegisterExporter("genesyscloud_routing_queue", RoutingQueueExporter())
	l.RegisterExporter("genesyscloud_routing_settings", RoutingSettingsExporter())
	l.RegisterExporter("genesyscloud_routing_skill", RoutingSkillExporter())
	l.RegisterExporter("genesyscloud_routing_skill_group", ResourceSkillGroupExporter())
	l.RegisterExporter("genesyscloud_routing_utilization", RoutingUtilizationExporter())
	l.RegisterExporter("genesyscloud_routing_wrapupcode", RoutingWrapupCodeExporter())
	l.RegisterExporter("genesyscloud_telephony_providers_edges_edge_group", EdgeGroupExporter())
	l.RegisterExporter("genesyscloud_telephony_providers_edges_extension_pool", TelephonyExtensionPoolExporter())
	l.RegisterExporter("genesyscloud_telephony_providers_edges_phonebasesettings", PhoneBaseSettingsExporter())
	l.RegisterExporter("genesyscloud_telephony_providers_edges_trunkbasesettings", TrunkBaseSettingsExporter())
	l.RegisterExporter("genesyscloud_telephony_providers_edges_trunk", TrunkExporter())
	l.RegisterExporter("genesyscloud_user", UserExporter())
	l.RegisterExporter("genesyscloud_user_roles", UserRolesExporter())
	l.RegisterExporter("genesyscloud_webdeployments_deployment", WebDeploymentExporter())
	l.RegisterExporter("genesyscloud_widget_deployment", WidgetDeploymentExporter())

	l.RegisterExporter("genesyscloud_knowledge_v1_document", KnowledgeDocumentExporterV1())
	l.RegisterExporter("genesyscloud_knowledge_document_variation", KnowledgeDocumentVariationExporter())
	l.RegisterExporter("genesyscloud_knowledge_label", KnowledgeLabelExporter())
	l.RegisterExporter("genesyscloud_knowledge_v1_category", KnowledgeCategoryExporterV1())

}
