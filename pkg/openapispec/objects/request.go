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
	RequestSubType    string
	TargetModel       *Name
	AncestorModels    []AncestorModel
	HasStatusResponse bool
	RelatedServices   []Name
}

// AncestorModel ...
type AncestorModel struct {
	Name      Name
	Parameter Name
	Column    Name
}
