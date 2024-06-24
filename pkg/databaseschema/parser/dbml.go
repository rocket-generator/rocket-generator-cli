package parser

import (
	"errors"
	"fmt"
	"github.com/duythinht/dbml-go/core"
	"github.com/duythinht/dbml-go/parser"
	"github.com/duythinht/dbml-go/scanner"
	"github.com/rocket-generator/rocket-generator-cli/pkg/data_mapper"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
	"os"
	"strings"
	"time"
)

type RelationKey struct {
	FromTableName  string
	FromColumnName string
	ToTableName    string
	ToColumnName   string
}

func ParseTables(dbmlObject *core.DBML, organizationName string, typeMapper *data_mapper.Mapper) []*objects.Entity {
	var entities []*objects.Entity
	for index, entity := range dbmlObject.Tables {
		entityName := entity.Name
		entityObject := objects.Entity{
			Name:               generateName(entityName),
			PrimaryKeyDataType: "int64",
			PrimaryKeyName:     "id",
			HasDecimal:         false,
			HasJSON:            false,
			UseSoftDelete:      false,
			DateTime:           time.Now().Format("20060201150405"),
			Year:               time.Now().Format("2006"),
			Month:              time.Now().Format("01"),
			Day:                time.Now().Format("02"),
			Time:               time.Now().Format("150405"),
			OrganizationName:   organizationName,
			Authenticatable:    false,
			RequiredIndexes:    [][]string{},
			Index:              index + 1,
			IndexString4Digit:  fmt.Sprintf("%04d", index+1),
			Hash12Digit:        fmt.Sprintf("%x%04d", time.Now().Format("20060102"), index+1),
		}
		for _, column := range entity.Columns {
			primary := false
			nullable := true
			name := column.Name
			dataType, dataSize := convertDataTypeToDataTypeAndSize(strings.ToLower(column.Type))
			defaultValue := column.Settings.Default
			if name == "created_at" || name == "updated_at" {
				continue
			}
			nullable = column.Settings.Null
			primary = column.Settings.PK
			if primary {
				if dataType != "uuid" && dataType != "text" {
					dataType = "bigserial"
				} else {
					defaultValue = "uuid_generate_v4()"
					entityObject.PrimaryKeyDataType = "uuid"
				}
			}

			columnObject := &objects.Column{
				TableName:    entityObject.Name,
				Name:         generateName(name),
				DataType:     generateName(dataType),
				DataSize:     dataSize,
				ObjectType:   data_mapper.MapString(typeMapper, "database", dataType),
				Primary:      primary,
				Nullable:     nullable,
				DefaultValue: defaultValue,
				Note:         column.Settings.Note,
			}
			columnObject.APIReturnable = checkAPIReturnable(columnObject)
			columnObject.APIUpdatable = checkAPIUpdatable(columnObject)
			columnObject.APIType = getAPIType(columnObject)
			if primary {
				columnObject.IsCommonColumn = true
			} else {
				columnObject.IsCommonColumn = false
			}
			entityObject.Columns = append(entityObject.Columns, columnObject)
			if primary {
				entityObject.PrimaryKey = columnObject
				entityObject.PrimaryKeyName = columnObject.Name.Original
			}
			if strings.HasPrefix(dataType, "decimal") || strings.HasPrefix(dataType, "numeric") {
				entityObject.HasDecimal = true
			}
			if strings.HasPrefix(dataType, "json") {
				entityObject.HasJSON = true
			}
			if name == "deleted_at" {
				columnObject.IsSystemUseColumn = true
				entityObject.UseSoftDelete = true
			} else {
				columnObject.IsSystemUseColumn = false
			}

			columnObject.FakerType = GuessFakerType(entityObject.Name.Original, columnObject)
		}
		entities = append(entities, &entityObject)
	}

	return entities
}

func addToRef(
	fromTableName string,
	fromColumnName string,
	toTableName string,
	toColumnName string,
	referenceType core.RelationshipType,
	resultHash *map[RelationKey]objects.Relation,
	data *objects.Schema,
) {
	leftTableIndex := findEntityIndex(fromTableName, data)
	rightTableIndex := findEntityIndex(toTableName, data)
	if leftTableIndex == -1 || rightTableIndex == -1 {
		return
	}
	leftTable := data.Entities[leftTableIndex]
	rightTable := data.Entities[rightTableIndex]
	leftColumnIndex := findRelationColumnIndex(fromColumnName, leftTable)
	rightColumnIndex := findRelationColumnIndex(toColumnName, rightTable)

	if leftColumnIndex == -1 || rightColumnIndex == -1 {
		return
	}

	leftRelation := objects.Relation{
		Name:             generateName(rightTable.Name.Original),
		Entity:           rightTable,
		Column:           rightTable.Columns[rightColumnIndex],
		MultipleEntities: false,
		ForeignKey:       generateName(fromColumnName),
		OwnerKey:         generateName(toColumnName),
	}
	rightRelation := objects.Relation{
		Name:             generateName(leftTable.Name.Original),
		Entity:           leftTable,
		Column:           leftTable.Columns[leftColumnIndex],
		MultipleEntities: false,
		ForeignKey:       generateName(toColumnName),
		OwnerKey:         generateName(fromColumnName),
	}

	switch referenceType {
	case core.OneToOne:
		leftRelation.RelationType = "hasOne"
		rightRelation.RelationType = "belongsTo"
		break
	case core.OneToMany:
		leftRelation.RelationType = "belongsTo"
		rightRelation.RelationType = "hasMany"
		rightRelation.MultipleEntities = true
		break
	case core.ManyToOne:
		leftRelation.RelationType = "hasMany"
		leftRelation.MultipleEntities = true
		rightRelation.RelationType = "belongsTo"
		break
	}

	(*resultHash)[RelationKey{
		FromTableName:  fromTableName,
		FromColumnName: fromColumnName,
		ToTableName:    toTableName,
		ToColumnName:   toColumnName,
	}] = leftRelation
	(*resultHash)[RelationKey{
		FromTableName:  toTableName,
		FromColumnName: toColumnName,
		ToTableName:    fromTableName,
		ToColumnName:   fromColumnName,
	}] = rightRelation
}

func ParseRefs(dbmlObject *core.DBML, data *objects.Schema) {
	relationEntries := make(map[RelationKey]objects.Relation)

	// Parse Inline Reference
	for _, entity := range dbmlObject.Tables {
		for _, column := range entity.Columns {
			if column.Settings.Ref.Type != core.None && column.Settings.Ref.To != "" {
				refTableName, refColumnName, err := parseRelation(column.Settings.Ref.To)
				if err == nil {
					continue
				}
				addToRef(entity.Name, column.Name, *refTableName, *refColumnName, column.Settings.Ref.Type, &relationEntries, data)
			}
		}
	}
	for _, relationGroup := range dbmlObject.Refs {
		for _, relation := range relationGroup.Relationships {
			fromTableName, fromColumnName, err := parseRelation(relation.From)
			if err != nil {
				continue
			}
			toTableName, toColumnName, err := parseRelation(relation.To)
			if err != nil {
				continue
			}
			addToRef(*fromTableName, *fromColumnName, *toTableName, *toColumnName, relation.Type, &relationEntries, data)
		}
	}

	for key, relation := range relationEntries {
		leftTableIndex := findEntityIndex(key.FromTableName, data)
		rightTableIndex := findEntityIndex(key.ToTableName, data)
		if leftTableIndex == -1 || rightTableIndex == -1 {
			continue
		}
		leftTable := data.Entities[leftTableIndex]
		rightTable := data.Entities[rightTableIndex]
		leftTable.Relations = append(leftTable.Relations, &relation)
		rightTable.Relations = append(rightTable.Relations, &relation)
	}
}

// ParseDBML ...
func ParseDBML(filePath string, projectName string, organizationName string, typeMapper *data_mapper.Mapper) (*objects.Schema, error) {
	data := objects.Schema{
		FilePath:           filePath,
		ProjectName:        projectName,
		OrganizationName:   organizationName,
		PrimaryKeyDataType: "int64",
	}

	fileHandler, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	s := scanner.NewScanner(fileHandler)
	dbmlParser := parser.NewParser(s)
	dbmlObject, err := dbmlParser.Parse()
	if err != nil {
		return nil, err
	}

	data.Entities = ParseTables(dbmlObject, organizationName, typeMapper)
	ParseRefs(dbmlObject, &data)

	return &data, nil
}

func parseRelation(relation string) (*string, *string, error) {
	parts := strings.Split(relation, ".")
	if len(parts) != 2 {
		return nil, nil, errors.New("invalid format: expected 'tablename.columnname'")
	}
	table := parts[0]
	column := parts[1]
	if table == "" || column == "" {
		return nil, nil, errors.New("invalid format: both tablename and columnname should be non-empty")
	}

	return &table, &column, nil
}
