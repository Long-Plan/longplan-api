package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"github.com/Long-Plan/longplan-api/pkg/mapper"
)

type curriculumService struct {
	sysCurriculumRepo               port.SysCurriculumRepo
	sysCurriculumQuestionRepo       port.SysCurriculumQuestionRepo
	sysCurriculumQuestionChoiceRepo port.SysCurriculumQuestionChoiceRepo
}

func NewCurriculumService(
	curriculumRepo port.SysCurriculumRepo,
	curriculumQuestionRepo port.SysCurriculumQuestionRepo,
	curriculumQuestionChoiceRepo port.SysCurriculumQuestionChoiceRepo,
) domain.CurriculumService {
	return &curriculumService{
		sysCurriculumRepo:               curriculumRepo,
		sysCurriculumQuestionRepo:       curriculumQuestionRepo,
		sysCurriculumQuestionChoiceRepo: curriculumQuestionChoiceRepo,
	}
}

func (s *curriculumService) GetAll() ([]dto.Curriculum, error) {
	var curricula []dto.Curriculum
	syscurricula, err := s.sysCurriculumRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, sysCurriculum := range syscurricula {
		curriculum, err := mapper.Mapper[model.SysCurriculum, dto.Curriculum](sysCurriculum)
		if err != nil {
			return nil, err
		}

		questions, err := mapper.MapSlice[model.SysCurriculumQuestion, dto.CurriculumQuestion](sysCurriculum.Questions)
		if err != nil {
			return nil, err
		}

		for i, question := range questions {
			choices, err := mapper.MapSlice[model.SysCurriculumQuestionChoice, dto.CurriculumQuestionChoice](sysCurriculum.Questions[i].Choices)
			if err != nil {
				return nil, err
			}
			question.Choices = choices
			questions[i] = question
		}

		curriculum.Questions = questions
		curricula = append(curricula, *curriculum)
	}

	return curricula, nil
}

func (s *curriculumService) GetAllByMajorID(majorID int) ([]dto.Curriculum, error) {
	var curricula []dto.Curriculum
	syscurricula, err := s.sysCurriculumRepo.GetAllByMajorID(majorID)
	if err != nil {
		return nil, err
	}

	for _, sysCurriculum := range syscurricula {
		curriculum, err := mapper.Mapper[model.SysCurriculum, dto.Curriculum](sysCurriculum)
		if err != nil {
			return nil, err
		}

		questions, err := mapper.MapSlice[model.SysCurriculumQuestion, dto.CurriculumQuestion](sysCurriculum.Questions)
		if err != nil {
			return nil, err
		}

		for i, question := range questions {
			choices, err := mapper.MapSlice[model.SysCurriculumQuestionChoice, dto.CurriculumQuestionChoice](sysCurriculum.Questions[i].Choices)
			if err != nil {
				return nil, err
			}
			question.Choices = choices
			questions[i] = question
		}

		curriculum.Questions = questions
		curricula = append(curricula, *curriculum)
	}
	return curricula, nil
}

func (s *curriculumService) GetByID(curriculumID int) (*dto.Curriculum, error) {
	sysCurriculum, err := s.sysCurriculumRepo.GetByID(curriculumID)
	if err != nil {
		return nil, err
	}

	curriculum, err := mapper.Mapper[model.SysCurriculum, dto.Curriculum](*sysCurriculum)
	if err != nil {
		return nil, err
	}

	questions, err := mapper.MapSlice[model.SysCurriculumQuestion, dto.CurriculumQuestion](sysCurriculum.Questions)
	if err != nil {
		return nil, err
	}

	for i, question := range questions {
		choices, err := mapper.MapSlice[model.SysCurriculumQuestionChoice, dto.CurriculumQuestionChoice](sysCurriculum.Questions[i].Choices)
		if err != nil {
			return nil, err
		}
		question.Choices = choices
		questions[i] = question
	}

	curriculum.Questions = questions

	return curriculum, nil
}

func (s *curriculumService) Create(curriculum dto.Curriculum) error {
	sysCurriculum, err := mapper.Mapper[dto.Curriculum, model.SysCurriculum](curriculum)
	if err != nil {
		return err
	}
	return s.sysCurriculumRepo.Create(sysCurriculum)
}

func (s *curriculumService) Update(curriculum dto.Curriculum) error {
	sysCurriculum, err := mapper.Mapper[dto.Curriculum, model.SysCurriculum](curriculum)
	if err != nil {
		return err
	}
	return s.sysCurriculumRepo.Update(sysCurriculum)
}

func (s *curriculumService) Delete(curriculumID int) error {
	return s.sysCurriculumRepo.Delete(curriculumID)
}
