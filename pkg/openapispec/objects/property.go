package objects

// Property ...
type Property struct {
	Name          Name
	Type          string
	ObjectType    string
	ArrayItemType string
	ArrayItemName string
	Description   string
	Reference     string
	Required      bool
}
