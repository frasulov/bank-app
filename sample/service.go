package sample

import (
	"context"
	my_errors "defaultProjectStructure_sqlc/errors"
	"github.com/gofiber/fiber/v2"
)

type SampleService struct {
	repository *SampleRepository
	//store      *db.Store
}

func GetNewSampleService(repository *SampleRepository) *SampleService {
	return &SampleService{
		repository: repository,
		//store:      store,
	}
}

// FindById example method
func (ss *SampleService) FindById(id int) (*SampleDto, error) {
	sample, err := ss.repository.GetSampleById(context.Background(), int64(id))
	if err != nil {
		return nil, my_errors.NewHttpError(fiber.StatusNotFound, my_errors.NewResponseByKey("not_found", "en"))
	}
	dto := ToSampleDto(sample)
	return &dto, nil
}
