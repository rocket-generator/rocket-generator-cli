package objects

type APIInfo struct {
	Path           string              `json:"path"`
	Method         string              `json:"method"`
	Type           string              `json:"type"`
	SubType        string              `json:"subType"`
	Group          string              `json:"group"`
	RequireAuth    bool                `json:"requireAuth"`
	RequiredRoles  []string            `json:"requiredRoles"`
	TargetModel    string              `json:"targetModel"`
	AncestorModels []AncestorJsonModel `json:"ancestorModels"`
}

type AncestorJsonModel struct {
	Name      string `json:"name"`
	Parameter string `json:"parameter"`
	Column    string `json:"column"`
}

type APIInfos []APIInfo
