package port

import "github.com/Long-Plan/longplan-api/internal/core/model"

type SysCourseDetailRepo interface {
	GetAll() ([]model.SysCourseDetail, error)
	GetByCourseNo(courseNo string) (*model.SysCourseDetail, error)
	Create(courseDetail *model.SysCourseDetail) error
	Update(courseDetail *model.SysCourseDetail) error
	Delete(courseNo string) error
}
