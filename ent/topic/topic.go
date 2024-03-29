// Code generated by ent, DO NOT EDIT.

package topic

const (
	// Label holds the string label denoting the topic type in the database.
	Label = "topic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSubName holds the string denoting the sub_name field in the database.
	FieldSubName = "sub_name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldGithubURL holds the string denoting the github_url field in the database.
	FieldGithubURL = "github_url"
	// FieldPlatformID holds the string denoting the platform_id field in the database.
	FieldPlatformID = "platform_id"
	// EdgePlatform holds the string denoting the platform edge name in mutations.
	EdgePlatform = "platform"
	// EdgeRepos holds the string denoting the repos edge name in mutations.
	EdgeRepos = "repos"
	// Table holds the table name of the topic in the database.
	Table = "topic"
	// PlatformTable is the table that holds the platform relation/edge.
	PlatformTable = "topic"
	// PlatformInverseTable is the table name for the Platform entity.
	// It exists in this package in order to avoid circular dependency with the "platform" package.
	PlatformInverseTable = "platform"
	// PlatformColumn is the table column denoting the platform relation/edge.
	PlatformColumn = "platform_id"
	// ReposTable is the table that holds the repos relation/edge.
	ReposTable = "repo"
	// ReposInverseTable is the table name for the Repo entity.
	// It exists in this package in order to avoid circular dependency with the "repo" package.
	ReposInverseTable = "repo"
	// ReposColumn is the table column denoting the repos relation/edge.
	ReposColumn = "topic_id"
)

// Columns holds all SQL columns for topic fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSubName,
	FieldDescription,
	FieldURL,
	FieldGithubURL,
	FieldPlatformID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
