package objects

// Schema ...
type Schema struct {
	FilePath           string
	ProjectName        string
	OrganizationName   string
	PackageName        string
	PrimaryKeyDataType string
	Entities           []*Entity
}
