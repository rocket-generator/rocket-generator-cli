package b_parse_service_definition_file

import (
	"encoding/json"
	createServiceCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/create/service"
	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/stoewer/go-strcase"
	"os"
)

type Process struct {
}

// Service は各サービスを表す構造体です。
type Service struct {
	Name          string        `json:"name"`
	APIEndpoints  []APIEndpoint `json:"api_endpoints"`
	RelatedModels []string      `json:"related_models"`
}

// APIEndpoint はAPIのエンドポイントを表す構造体です。
type APIEndpoint struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}

// Services は複数のServiceを格納するためのスライスです。
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
			argument := createServiceCommand.Arguments{
				ServiceName:       service.Name,
				RelatedModelNames: service.RelatedModels,
				ProjectPath:       payload.ProjectPath,
				Debug:             payload.Debug,
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
		if len(name) > 7 && name[len(name)-7:] == "Service" {
			services[index].Name = service.Name[:len(service.Name)-7]
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
