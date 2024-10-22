package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"github.com/Long-Plan/longplan-api/pkg/mapper"
)

type studentCurriculumService struct {
	studentCurriculumRepo               port.StudentCurriculumRepo
	studentCurriculumCourseRepo         port.StudentCurriculumCourseRepo
	studentCurriculumQuestionAnswerRepo port.StudentCurriculumQuestionAnswerRepo
	sysCategoryCourseRepo               port.SysCategoryCourseRepo
}

func NewStudentCurriculumService(
	studentCurriculumRepo port.StudentCurriculumRepo,
	studentCurriculumCourseRepo port.StudentCurriculumCourseRepo,
	studentCurriculumQuestionAnswerRepo port.StudentCurriculumQuestionAnswerRepo,
	sysCategoryCourseRepo port.SysCategoryCourseRepo,
) domain.StudentCurriculumService {
	return &studentCurriculumService{
		studentCurriculumRepo:               studentCurriculumRepo,
		studentCurriculumCourseRepo:         studentCurriculumCourseRepo,
		studentCurriculumQuestionAnswerRepo: studentCurriculumQuestionAnswerRepo,
		sysCategoryCourseRepo:               sysCategoryCourseRepo,
	}
}

func (s *studentCurriculumService) GetByStudentCode(studentCode int, majorId int) ([]dto.StudentCurriculum, error) {
	studentCurricula, err := s.studentCurriculumRepo.GetByStudentCode(studentCode, majorId)
	if err != nil {
		return nil, err
	}
	studentCurriculaDto, err := mapper.MapSlice[model.StudentCurriculum, dto.StudentCurriculum](studentCurricula)
	if err != nil {
		return nil, err
	}

	for i, studentCurriculum := range studentCurriculaDto {
		studentCurriculumCourses, err := s.studentCurriculumCourseRepo.GetByStudentCurriculumID(studentCurriculum.ID)
		if err != nil {
			return nil, err
		}

		courses, err := mapper.MapSlice[model.StudentCurriculumCourse, dto.StudentCurriculumCourse](studentCurriculumCourses)
		if err != nil {
			return nil, err
		}
		studentCurriculaDto[i].Courses = courses

		studentCurriculumQuestionAnswers, err := s.studentCurriculumQuestionAnswerRepo.GetByStudentCurriculumID(studentCurriculum.ID)
		if err != nil {
			return nil, err
		}

		answers, err := mapper.MapSlice[model.StudentCurriculumQuestionAnswer, dto.StudentCurriculumQuestionAnswer](studentCurriculumQuestionAnswers)
		if err != nil {
			return nil, err
		}
		studentCurriculaDto[i].Answers = answers
	}
	return studentCurriculaDto, nil
}

func (s *studentCurriculumService) GetByStudentCurriculumID(studentCurriculumID int) (*dto.StudentCurriculum, error) {
	studentCurriculum, err := s.studentCurriculumRepo.GetByStudentCurriculumID(studentCurriculumID)
	if err != nil {
		return nil, err
	}

	studentCurriculumDto, err := mapper.Mapper[model.StudentCurriculum, dto.StudentCurriculum](*studentCurriculum)
	if err != nil {
		return nil, err
	}

	studentCurriculumCourses, err := s.studentCurriculumCourseRepo.GetByStudentCurriculumID(studentCurriculum.ID)
	if err != nil {
		return nil, err
	}

	courses, err := mapper.MapSlice[model.StudentCurriculumCourse, dto.StudentCurriculumCourse](studentCurriculumCourses)
	if err != nil {
		return nil, err
	}
	studentCurriculumDto.Courses = courses

	studentCurriculumQuestionAnswers, err := s.studentCurriculumQuestionAnswerRepo.GetByStudentCurriculumID(studentCurriculum.ID)
	if err != nil {
		return nil, err
	}

	answers, err := mapper.MapSlice[model.StudentCurriculumQuestionAnswer, dto.StudentCurriculumQuestionAnswer](studentCurriculumQuestionAnswers)
	if err != nil {
		return nil, err
	}
	studentCurriculumDto.Answers = answers

	return studentCurriculumDto, nil
}

func (s *studentCurriculumService) Create(studentCurriculum dto.StudentCurriculumCreate) (*int, error) {
	courses, err := s.sysCategoryCourseRepo.GetByCurriculumID(studentCurriculum.CurriculumID)
	if err != nil {
		return nil, err
	}

	studentCurriculumModel := model.StudentCurriculum{
		Name:         studentCurriculum.Name,
		StudentCode:  studentCurriculum.StudentCode,
		CurriculumID: studentCurriculum.CurriculumID,
		IsSystem:     studentCurriculum.IsSystem,
	}

	err = s.studentCurriculumRepo.Create(&studentCurriculumModel)
	if err != nil {
		return nil, err
	}

	for _, answer := range studentCurriculum.Answers {
		studentCurriculumQuestionAnswer := model.StudentCurriculumQuestionAnswer{
			StudentCurriculumID: studentCurriculumModel.ID,
			QuestionID:          answer.QuestionID,
			ChoiceID:            answer.ChoiceID,
		}
		if err := s.studentCurriculumQuestionAnswerRepo.Create(&studentCurriculumQuestionAnswer); err != nil {
			return nil, err
		}
	}

	for _, course := range courses {
		if course.Year != nil && course.Semester != nil {
			studentCurriculumCourse := model.StudentCurriculumCourse{
				StudentCurriculumID: studentCurriculumModel.ID,
				Year:                *course.Year,
				Semester:            *course.Semester,
				CourseNo:            course.CourseNo,
				CategoryID:          course.CategoryID,
			}
			if err := s.studentCurriculumCourseRepo.Create(&studentCurriculumCourse); err != nil {
				return nil, err
			}
		}
	}
	return &studentCurriculumModel.ID, nil
}

func (s *studentCurriculumService) Update(studentCurriculum model.StudentCurriculum) error {
	return s.studentCurriculumRepo.Update(&studentCurriculum)
}

func (s *studentCurriculumService) Delete(studentCurriculumID int) error {
	return s.studentCurriculumRepo.Delete(studentCurriculumID)
}

func (s *studentCurriculumService) UpdateCourses(studentCurriculumID int, courses []model.StudentCurriculumCourse) error {
	return s.studentCurriculumCourseRepo.Updates(courses)
}

func (s *studentCurriculumService) UpdateQuestionAnswers(studentCurriculumID int, questions []dto.StudentCurriculumQuestionAnswer) error {
	answers, err := mapper.MapSlice[dto.StudentCurriculumQuestionAnswer, model.StudentCurriculumQuestionAnswer](questions)
	if err != nil {
		return err
	}
	for i, answer := range answers {
		answer.StudentCurriculumID = studentCurriculumID
		answers[i] = answer
	}
	return s.studentCurriculumQuestionAnswerRepo.Updates(answers)
}
