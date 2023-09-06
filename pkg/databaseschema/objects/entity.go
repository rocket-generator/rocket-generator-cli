package objects

// Entity ...
type Entity struct {
	Name               Name
	PrimaryKeyDataType string
	Columns            []*Column
	Relations          []*Relation
	PrimaryKey         *Column
	Description        string
	HasDecimal         bool
	HasJSON            bool
	PackageName        string
	UseSoftDelete      bool
	DateTime           string
	Year               string
	Month              string
	Day                string
	Time               string
	OrganizationName   string
}
