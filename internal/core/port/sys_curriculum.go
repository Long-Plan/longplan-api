package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCurriculumRepo interface {
	GetAll() ([]model.SysCurriculum, error)
	GetAllByMajorID(majorID int) ([]model.SysCurriculum, error)
	GetByID(curriculumId int) (*model.SysCurriculum, error)
	Create(curriculum *model.SysCurriculum) error
	Update(curriculum *model.SysCurriculum) error
	Delete(curriculumId int) error
}
