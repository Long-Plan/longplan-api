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
	sysCategoryCourseRepo           port.SysCategoryCourseRepo
	sysCourseDetailRepo             port.SysCourseDetailRepo
}

func NewCurriculumService(
	curriculumRepo port.SysCurriculumRepo,
	curriculumQuestionRepo port.SysCurriculumQuestionRepo,
	curriculumQuestionChoiceRepo port.SysCurriculumQuestionChoiceRepo,
	categoryCourseRepo port.SysCategoryCourseRepo,
	courseDetail port.SysCourseDetailRepo,
) domain.CurriculumService {
	return &curriculumService{
		sysCurriculumRepo:               curriculumRepo,
		sysCurriculumQuestionRepo:       curriculumQuestionRepo,
		sysCurriculumQuestionChoiceRepo: curriculumQuestionChoiceRepo,
		sysCategoryCourseRepo:           categoryCourseRepo,
		sysCourseDetailRepo:             courseDetail,
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
	sysCurricula, err := s.sysCurriculumRepo.GetAllByMajorID(majorID)
	if err != nil {
		return nil, err
	}

	for _, sysCurriculum := range sysCurricula {
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

func (s *curriculumService) GetCoursesByCurriculumID(curriculumID int) ([]dto.CategoryCourse, error) {
	sysCategoryCourses, err := s.sysCategoryCourseRepo.GetByCurriculumID(curriculumID)
	if err != nil {
		return nil, err
	}

	categoryCourses, err := mapper.MapSlice[model.SysCategoryCourse, dto.CategoryCourse](sysCategoryCourses)
	if err != nil {
		return nil, err
	}

	for i, categoryCourse := range categoryCourses {
		requisites, err := mapper.MapSlice[model.SysCategoryCourseRequisite, dto.CategoryCourseRequisite](sysCategoryCourses[i].Requisites)
		if err != nil {
			return nil, err
		}
		categoryCourses[i].Requisites = requisites

		if categoryCourse.CourseNo != nil {
			courseDetail, err := s.sysCourseDetailRepo.GetByCourseNo(*categoryCourse.CourseNo)
			if err != nil {
				return nil, err
			}
			categoryCourses[i].Detail = *courseDetail
		}
		categoryCourses[i].Credit = sysCategoryCourses[i].Credit
	}

	return categoryCourses, nil
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
