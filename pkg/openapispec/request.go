package openapispec

// Request ...
type Request struct {
	Path              string
	PackageName       string
	RouteNameSpace    string
	OrganizationName  string
	Method            NameForm
	PathName          Name
	Description       string
	AddParamsForTest  string
	RequestSchemaName Name
	RequestSchema     *Schema
	Parameters        []*Parameter
	Responses         []*Response
	SuccessResponse   *Response
}
