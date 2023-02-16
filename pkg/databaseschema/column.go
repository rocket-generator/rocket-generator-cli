package databaseschema

// Column ...
type Column struct {
	Name              Name
	DataType          Name
	Primary           bool
	DefaultValue      string
	Nullable          bool
	APIReturnable     bool
	APIUpdatable      bool
	APIType           string
	FakerType         string
	IsCommonColumn    bool
	TableName         Name
	RelationTableName Name
	IsSystemUseColumn bool
}
