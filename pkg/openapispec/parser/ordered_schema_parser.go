package parser

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rocket-generator/rocket-generator-cli/pkg/openapispec/objects"
	"github.com/stoewer/go-strcase"
	"gopkg.in/yaml.v3"
)

func ParseSchemaPropertyOrder(filePath string) *map[string]objects.OrderedProperties {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil
	}

	var root yaml.Node
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&root)
	if err != nil {
		fmt.Printf("Error parsing YAML: %v\n", err)
		return nil
	}

	// ドキュメントノードの最初の要素を取得
	if len(root.Content) == 0 {
		return nil
	}
	docNode := root.Content[0]

	// components.schemasを探す
	componentsNode := findChildNode(docNode, "components")
	if componentsNode == nil {
		return nil
	}

	schemasNode := findChildNode(componentsNode, "schemas")
	if schemasNode == nil {
		return nil
	}

	schemas := make(map[string]objects.OrderedProperties)

	// スキーマを処理
	for i := 0; i < len(schemasNode.Content); i += 2 {
		schemaName := schemasNode.Content[i].Value
		schemaNode := schemasNode.Content[i+1]

		propertiesNode := findChildNode(schemaNode, "properties")
		if propertiesNode == nil {
			continue
		}

		var properties []string
		for j := 0; j < len(propertiesNode.Content); j += 2 {
			keyNode := propertiesNode.Content[j]
			propName := keyNode.Value
			properties = append(properties, propName)
		}
		snakeCaseSchemaName := strcase.SnakeCase(schemaName)
		schemas[snakeCaseSchemaName] = objects.OrderedProperties{
			Name:       schemaName,
			Properties: properties,
		}
	}

	return &schemas
}

// findChildNode は指定されたキーを持つ子ノードを探す
func findChildNode(node *yaml.Node, key string) *yaml.Node {
	if node.Kind != yaml.MappingNode {
		return nil
	}

	for i := 0; i < len(node.Content); i += 2 {
		keyNode := node.Content[i]
		if keyNode.Value == key {
			return node.Content[i+1]
		}
	}
	return nil
}

/*
// デバッグ用の補助関数
func dumpNode(node *yaml.Node, indent string) {
	if node == nil {
		return
	}

	fmt.Printf("%sKind: %v, Value: %v, Line: %d, Column: %d\n",
		indent, node.Kind, node.Value, node.Line, node.Column)

	for _, child := range node.Content {
		dumpNode(child, indent+"  ")
	}
}
*/
