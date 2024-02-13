package objects

// Request ...
type Request struct {
	Path              string
	GroupRelativePath string
	PackageName       string
	RouteNameSpace    string
	OrganizationName  string
	Method            NameForm
	PathName          Name
	Description       string
	AddParamsForTest  string
	Services          []string
	RequireAuth       bool
	RequiredRoles     []string
	RequestSchemaName Name
	RequestSchema     *Schema
	Parameters        []*Parameter
	Responses         []*Response
	SuccessResponse   *Response
	RequestType       string
	TargetModel       string
	HasStatusResponse bool
}
