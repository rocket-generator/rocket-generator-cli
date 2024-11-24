package b_parse_service_definition_file

import (
	"encoding/json"
	"os"
	"strings"

	createServiceCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/service"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/stoewer/go-strcase"
)

type Process struct {
}

type Model struct {
	Name        string
	IsCRUDModel bool
}

type Service struct {
	Name          string        `json:"name"`
	APIEndpoints  []APIEndpoint `json:"apiEndpoints"`
	IsAuthService bool          `json:"isAuthService"`
	RelatedModels []Model       `json:"models"`
}

type APIEndpoint struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

type Services []Service

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	// Ignore if there is no service file
	if payload.ServiceFileName == "" {
		return payload, nil
	}
	services, err := loadServiceFile(payload.ServiceFileName)
	if err != nil {
		return nil, err
	}

	if services != nil {
		setServiceToAPISpec(payload, services)
		for _, service := range *services {

			var relatedModelNames []string
			var relatedModelWithCRUDNames []string
			for _, model := range service.RelatedModels {
				if model.IsCRUDModel {
					relatedModelWithCRUDNames = append(relatedModelWithCRUDNames, model.Name)
				} else {
					relatedModelNames = append(relatedModelNames, model.Name)
				}

			}

			argument := createServiceCommand.Arguments{
				Type:                      "service",
				Name:                      service.Name,
				RelatedModelNames:         relatedModelNames,
				RelatedModelWithCRUDNames: relatedModelWithCRUDNames,
				RelatedResponse:           nil,
				IsAuthService:             service.IsAuthService,
				ProjectPath:               payload.ProjectPath,
				Debug:                     payload.Debug,
			}
			command := createServiceCommand.Command{}
			err := command.Execute(argument)
			if err != nil {
				return nil, err
			}
		}
	}
	return payload, nil
}

func loadServiceFile(filename string) (*Services, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var services Services
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&services); err != nil {
		return nil, err
	}

	for index, service := range services {
		name := strcase.LowerCamelCase(service.Name)

		if strings.HasSuffix(name, "Service") {
			services[index].Name = strings.TrimSuffix(name, "Service")
		} else {
			services[index].Name = service.Name
		}

	}
	return &services, nil
}

func setServiceToAPISpec(payload *newCommand.Payload, services *Services) {
	for _, service := range *services {
		for _, targetRequest := range service.APIEndpoints {
			method := strcase.LowerCamelCase(targetRequest.Method)
			for _, request := range payload.OpenAPISpec.Requests {
				if request.Path == targetRequest.Path && request.Method.Snake == method {
					request.Services = append(request.Services, service.Name)
				}
			}
		}
	}
}
