package go_pex

// PermissionTag is the tag to use in structs to specify the permissions of each field
const PermissionTag = "pex"

// Actions
const (
	// ActionWrite is used when the action is writing
	ActionWrite = 0
	// ActionRead is used when the action is reading
	ActionRead = 1
)

// Permissions
const (
	// PermissionNone means it hasn't any permission
	PermissionNone = 0
	// PermissionRead means it has reading permissions
	PermissionRead = 1
	// PermissionWrite means it has write permissions
	PermissionWrite = 2
	// PermissionReadWrite means it has read and write permissions
	PermissionReadWrite = 3
)