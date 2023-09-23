package parser

import (
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"strings"
)

func IsTypeInteger(dataType string) bool {
	switch strings.ToLower(dataType) {
	case "int":
	case "integer":
	case "bigint":
	case "bigserial":
	case "number":
		return true
	}

	return false
}

func IsTypeString(dataType string) bool {
	switch strings.ToLower(dataType) {
	case "text":
	case "string":
	case "varchar":
	case "char":
		return true
	}
	return false
}

func GuessFakerType(tableName string, column *objects.Column) string {
	nameMap := map[string]string{
		"longitude":     "longitude",
		"latitude":      "latitude",
		"email":         "email",
		"phone_number":  "phoneNumber",
		"phone":         "phoneNumber",
		"language_code": "languageCode",
		"country_code":  "countryCode",
		"url":           "url",
		"slug":          "slug",
		"description":   "sentence",
	}

	typeMap := map[string]string{
		"bool":     "boolean",
		"boolean":  "boolean",
		"uuid":     "uuid",
		"date":     "date",
		"datetime": "dateTime",
	}

	columnName := column.Name.Default.Snake
	fakerTypeString := column.DataType.Original
	for key, value := range nameMap {
		if strings.HasSuffix(columnName, key) {
			return value
		}
	}

	if IsTypeInteger(fakerTypeString) {
		if strings.HasSuffix(columnName, "_at") {
			return "unixTime"
		}
	}

	if IsTypeString(fakerTypeString) {
		return "word"
	}

	for key, value := range typeMap {
		if column.DataType.Default.Snake == key {
			return value
		}
	}

	return "word"
}
