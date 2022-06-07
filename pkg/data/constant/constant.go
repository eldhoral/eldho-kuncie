package constant

import "os"

const (
	DefaultDatetimeLayout         = "2006-01-02 15:04:05"
	DefaultDatetimeTimezoneLayout = "2006-01-02T15:04:05Z"

	DefaultDateLayout = "2006-01-02"

	DefaultTimeLayout = "15:04:05"

	YYYYmmddHHmmss = "20060102150405"

	YYYYLayout = "2006"
	MMLayout   = "01"

	TaskListDateLayout = "02 Jan 2006"
	HHMMLayout         = "15:04"

	TalentaURL     = "https://hr.talenta.co"
	TalentaIconURL = "https://talenta.s3-ap-southeast-1.amazonaws.com/assets/images/talentaIcon.png"
	BlankAvatarURL = "https://talenta.s3-ap-southeast-1.amazonaws.com/avatar/blank.jpg"

	CompanyTaskComment                = 0
	CompanyTaskCommentChangeStatus    = 1
	CompanyTaskCommentCreatedTask     = 2
	CompanyTaskCommentDeletedActivity = 3
	CompanyTaskCommentAddAttachment   = 4
	MaxFileSize                       = 10485760

	TimesheetSummaryFileName = "Timesheet_Summary_Report"
	ContentTypeExcel         = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

var (
	EmptyArray         = []int{} // For return [], compatibility with Yii2
	EmptyStringPointer = func() *string { i := ""; return &i }()
	TimesheetGroups    = []string{"project", "task", "employee"}
	AllowedFile        = []string{".pdf", ".jpg", ".png", ".xlsx", ".xls", ".jpeg", ".docx", ".doc", ".csv", ".txt", ".ppt", ".pptx"}
)

// IsMySQL check is MySQL (UTC), otherwise MariaDB (+7 GMT TAL)
// go.mekari.io/golden-path MariaDB is deprecated, use MySQL+UTC for future development
// Based on IsMySQL, we can change SQL, and timezone to match with DB
func IsMySQL() bool {
	return os.Getenv("DB_TZ") == "UTC"
}
