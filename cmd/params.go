package cmd

import (
	"os"
	"path"
	"runtime"
)

func initParams() map[string]string {
	params = make(map[string]string)

	params["APP_ENV"] = os.Getenv("APP_ENV")
	params["DD_AGENT_ADDR"] = os.Getenv("DD_AGENT_ADDR")

	params["kafka-brokers"] = os.Getenv("ULMS_KAFKA_BROKERS")
	params["kafka-mixpanel-topic"] = "tnt.mixpanel"
	params["kafka-mail-topic"] = "tnt.mail"

	params["ulms-kafka-brokers"] = os.Getenv("ULMS_KAFKA_BROKERS")
	params["flock-error-report-channel"] = os.Getenv("FLOCK_ERROR_REPORT_CHANNEL")
	params["app-version"] = os.Getenv("APP_VERSION")
	params["app-name"] = os.Getenv("APP_NAME")

	params["talenta-url"] = os.Getenv("TALENTA_URL")
	params["tnt-api-key"] = os.Getenv("TNT_API_KEY")

	// Talenta Core public API
	params["user-role-validation-path"] = "/public_api/user-service/get-user-role-validation"
	params["inbox-timesheet-summary-path"] = "/public_api/user-service/inbox-timesheet-summary"
	params["get-employee-list"] = "/public_api/user-service/get-employee-list"
	params["employee-by-approval-line"] = "/public_api/user-service/get-employee-by-approval-line"
	params["employee-resign-by-company-id"] = "/public_api/user-service/employee-resign-by-company-id"
	params["get-organization-by-id"] = "/public_api/user-service/organization"
	params["job-structure-by-parent-id-and-company-id"] = "/public_api/user-service/job-structure"
	params["employee-resign-by-employee-id"] = "/public_api/user-service/employee-resign-by-employee-id"
	params["get-file-storage-by-owner-id"] = "/public_api/user-service/file-storage-by-owner"
	params["get-branch-by-id-and-company-id"] = "/public_api/user-service/branch"
	params["get-company-by-id"] = "/public_api/user-service/get-company-by-id"
	params["get-filter-branch-and-organization-and-name"] = "/public_api/user-service/get-filter-branch-and-organization-and-name"
	params["is-resign"] = "/public_api/user-service/is-resign"

	_, b, _, _ := runtime.Caller(0)
	appDir := path.Join(path.Dir(b), "..")
	params["app-dir"] = appDir
	params["template-dir"] = appDir + "/internal/template"

	// For staging
	if !isProd() {
		if params["talenta-url"] == "" {
			//params["talenta-url"] = "https://talenta-staging-core-static-web.talentadev.com/" // Set to default web staging
			params["talenta-url"] = "https://talenta-staging-core-scaling-sso.talentadev.com" // Set to default web staging
		}
	}

	return params
}
