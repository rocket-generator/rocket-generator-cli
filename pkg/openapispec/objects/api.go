package objects

// API ...
type API struct {
	FilePath         string
	ProjectName      string
	APINameSpace     string
	PackageName      string
	OrganizationName string
	BasePath         string
	RouteNameSpace   string
	Requests         []*Request
	RequestGroups    *[]*RequestGroup
	Schemas          map[string]*Schema
}
