package sample

import (
	"database/sql"
	db "defaultProjectStructure_sqlc/db/sqlc"
)

type SampleRepository struct {
	*db.Queries
	*sql.DB
}

func NewSampleRepository(connection *sql.DB) *SampleRepository {
	return &SampleRepository{
		Queries: db.New(connection),
		DB:      connection,
	}
}

// write your custom repository functions

// example
func (sr *SampleRepository) CustomGetSampleById() {
	// logic
}
