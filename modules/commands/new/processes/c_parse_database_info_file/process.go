package c_parse_database_info_file

import (
	"encoding/json"
	"fmt"
	"os"

	newCommand "github.com/rocket-generator/rocket-generator-cli/modules/commands/new/payload"
	"github.com/rocket-generator/rocket-generator-cli/pkg/databaseschema/objects"
)

type Process struct {
}

type DB struct {
	Authenticatable bool       `json:"authenticatable"`
	RequiredIndexes [][]string `json:"requiredIndexes"`
}

type DBs map[string]DB

func (process *Process) Execute(payload *newCommand.Payload) (*newCommand.Payload, error) {
	if payload.DatabaseInfoFileName == "" {
		return payload, nil
	}
	databases, err := loadDatabaseInfoFile(payload.DatabaseInfoFileName)
	if err != nil {
		return nil, err
	}
	updateDBSpec(payload, databases)

	return payload, nil
}

func loadDatabaseInfoFile(filename string) (*DBs, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var databases DBs
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&databases); err != nil {
		return nil, err
	}

	return &databases, nil
}

func updateDBSpec(payload *newCommand.Payload, databases *DBs) {
	for index, table := range payload.DatabaseSchema.Entities {
		if db, ok := (*databases)[table.Name.Default.Snake]; ok {
			payload.DatabaseSchema.Entities[index].Authenticatable = db.Authenticatable
			if db.Authenticatable {
				fmt.Println("Authenticatable table: " + table.Name.Default.Snake)
			} else {
				fmt.Println("Normal: " + table.Name.Default.Snake)
			}
			for _, requiredIndex := range db.RequiredIndexes {
				payload.DatabaseSchema.Entities[index].Indexes = append(payload.DatabaseSchema.Entities[index].Indexes, &objects.Index{
					Columns:  requiredIndex,
					IsUnique: true,
				})
			}
		}
	}
}
