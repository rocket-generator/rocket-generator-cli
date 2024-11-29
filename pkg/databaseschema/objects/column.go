package objects

// Column ...
type Column struct {
	Name              Name
	DataType          Name
	DataSize          int
	ObjectType        string
	CodeType          string
	Primary           bool
	DefaultValue      string
	Nullable          bool
	APIReturnable     bool
	APIUpdatable      bool
	APIType           string
	FakerType         string
	IsCommonColumn    bool
	IsForeignKey      bool
	TableName         Name
	RelationTableName Name
	IsSystemUseColumn bool
	Note              string
}
