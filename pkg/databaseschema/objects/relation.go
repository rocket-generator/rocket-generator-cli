package objects

// Relation ...
type Relation struct {
	Name             Name
	ForeignKey       Name
	OwnerKey         Name
	Entity           *Entity
	Column           *Column
	RelationType     string
	MultipleEntities bool
}
