package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
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

func (s *studentCurriculumService) GetByStudentCode(studentCode int) ([]model.StudentCurriculum, error) {
	return s.studentCurriculumRepo.GetByStudentCode(studentCode)
}

func (s *studentCurriculumService) GetByStudentCurriculumID(studentCurriculumID int) (*model.StudentCurriculum, error) {
	return s.studentCurriculumRepo.GetByStudentCurriculumID(studentCurriculumID)
}

func (s *studentCurriculumService) Create(studentCurriculum model.StudentCurriculum) error {
	courses, err := s.sysCategoryCourseRepo.GetByCurriculumID(studentCurriculum.CurriculumID)
	if err != nil {
		return err
	}

	for _, course := range courses {
		if course.Year != nil && course.Semester != nil {
			studentCurriculumCourse := model.StudentCurriculumCourse{
				StudentCurriculumID: studentCurriculum.ID,
				Year:                *course.Year,
				Semester:            *course.Semester,
				CourseNo:            course.CourseNo,
				CategoryID:          course.CategoryID,
			}
			if err := s.studentCurriculumCourseRepo.Create(&studentCurriculumCourse); err != nil {
				return err
			}
		}
	}
	return s.studentCurriculumRepo.Create(&studentCurriculum)
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

func (s *studentCurriculumService) UpdateQuestionAnswers(studentCurriculumID int, questions []model.StudentCurriculumQuestionAnswer) error {
	return s.studentCurriculumQuestionAnswerRepo.Updates(questions)
}
