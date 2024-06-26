package usecase

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
)

type ExerciceTrainingUseCase struct {
	TrainingExerciceInterface entity.TrainingExerciceInterface
}

func NewExerciceTraining(
	TrainingExerciceInterface entity.TrainingExerciceInterface,
) *ExerciceTrainingUseCase {
	return &ExerciceTrainingUseCase{
		TrainingExerciceInterface: TrainingExerciceInterface,
	}
}

func (a *ExerciceTrainingUseCase) CreateNewExerciceTraining(input dto.TrainingExerciceDto) (dto.TrainingExerciceDto, error) {
	ExerciceTraining := entity.TrainingExerciceEntity{
		ID:                   input.ID,
		IdExercice:           input.IdExercice,
		IdTraining:           input.IdTraining,
		Sets:                 input.Sets,
		Reps:                 input.Reps,
		DurationSeconds:      input.DurationSeconds,
		TotalDurationSeconds: input.TotalDurationSeconds,
		Order:                input.Order,
	}
	if _, err := a.TrainingExerciceInterface.Create(&ExerciceTraining); err != nil {
		return dto.TrainingExerciceDto{}, err
	}
	dto := dto.TrainingExerciceDto{
		ID:                   ExerciceTraining.ID,
		IdExercice:           input.IdExercice,
		IdTraining:           input.IdTraining,
		Sets:                 input.Sets,
		Reps:                 input.Reps,
		DurationSeconds:      input.DurationSeconds,
		TotalDurationSeconds: input.TotalDurationSeconds,
		Order:                input.Order,
	}

	return dto, nil
}

func (a *ExerciceTrainingUseCase) FindAllExerciceTrainings(page, limit int, sort string) ([]dto.TrainingExerciceDto, error) {
	entities, err := a.TrainingExerciceInterface.FindAll(page, limit, sort)
	if err != nil {
		return nil, err
	}
	var ExerciceTraining []dto.TrainingExerciceDto
	for _, entity := range entities {
		dto := dto.TrainingExerciceDto{
			ID:                   entity.ID,
			IdExercice:           entity.IdExercice,
			IdTraining:           entity.IdTraining,
			Sets:                 entity.Sets,
			Reps:                 entity.Reps,
			DurationSeconds:      entity.DurationSeconds,
			TotalDurationSeconds: entity.TotalDurationSeconds,
			Order:                entity.Order,
		}
		ExerciceTraining = append(ExerciceTraining, dto)
	}

	return ExerciceTraining, nil
}

func (a *ExerciceTrainingUseCase) FindExerciceTrainingByID(id string) (dto.TrainingExerciceDto, error) {
	entity, err := a.TrainingExerciceInterface.FindByID(id)
	if err != nil {
		return dto.TrainingExerciceDto{}, err
	}
	dto := dto.TrainingExerciceDto{
		ID:                   entity.ID,
		IdExercice:           entity.IdExercice,
		IdTraining:           entity.IdTraining,
		Sets:                 entity.Sets,
		Reps:                 entity.Reps,
		DurationSeconds:      entity.DurationSeconds,
		TotalDurationSeconds: entity.TotalDurationSeconds,
		Order:                entity.Order,
	}
	return dto, nil
}
