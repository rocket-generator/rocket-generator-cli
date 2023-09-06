package objects

// Relation ...
type Relation struct {
	Name             Name
	Entity           *Entity
	Column           *Column
	RelationType     string
	MultipleEntities bool
}
