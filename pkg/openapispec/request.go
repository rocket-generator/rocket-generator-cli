package openapispec

// Request ...
type Request struct {
	Path                      string
	PackageName               string
	RouteNameSpace            string
	PathFrameworkPresentation string
	Method                    NameForm
	PathNameForm              NameForm
	Description               string
	AddParamsForTest          string
	RequestSchemaName         Name
	Parameters                []*Parameter
	Responses                 []*Response
}
