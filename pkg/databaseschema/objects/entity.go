package objects

// Entity ...
type Entity struct {
	Name               Name
	Columns            []*Column
	Relations          []*Relation
	PrimaryKey         *Column
	PrimaryKeyName     string
	PrimaryKeyDataType string
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
	Authenticatable    bool
	RequiredIndexes    [][]string
}
