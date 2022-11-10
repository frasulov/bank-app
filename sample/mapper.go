package sample

import (
	"database/sql"
	db "defaultProjectStructure_sqlc/db/sqlc"
)

func ToSampleDto(model db.Sample) SampleDto {
	return SampleDto{
		Id:   model.ID,
		Name: model.Name.String,
	}
}

func ToSampleModel(dto SampleDto) db.Sample {
	return db.Sample{
		ID:   dto.Id,
		Name: sql.NullString{String: dto.Name, Valid: true},
	}
}

func ToSampleDtoArray(models []db.Sample) (result []SampleDto) {
	for _, val := range models {
		result = append(result, ToSampleDto(val))
	}
	return
}

func ToSampleModelArray(dtoArray []SampleDto) (result []db.Sample) {
	for _, val := range dtoArray {
		result = append(result, ToSampleModel(val))
	}
	return
}
