package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"github.com/Long-Plan/longplan-api/pkg/mapper"
	"github.com/samber/lo"
)

type categoryService struct {
	categoryRepo                port.SysCategoryRepo
	categoryTypeRepo            port.SysCategoryTypeRepo
	categoryRequirementRepo     port.SysCategoryRequirementRepo
	categoryRelationshipRepo    port.SysCategoryRelationshipRepo
	categoryCourseRepo          port.SysCategoryCourseRepo
	categoryCourseRequisiteRepo port.SysCategoryCourseRequisiteRepo
	courseDetailRepo            port.SysCourseDetailRepo
}

func NewCategoryService(
	categoryRepo port.SysCategoryRepo,
	categoryTypeRepo port.SysCategoryTypeRepo,
	categoryRequirementRepo port.SysCategoryRequirementRepo,
	categoryRelationshipRepo port.SysCategoryRelationshipRepo,
	categoryCourseRepo port.SysCategoryCourseRepo,
	categoryCourseRequisiteRepo port.SysCategoryCourseRequisiteRepo,
	courseDetailRepo port.SysCourseDetailRepo,
) domain.CategoryService {
	return &categoryService{
		categoryRepo:                categoryRepo,
		categoryTypeRepo:            categoryTypeRepo,
		categoryRequirementRepo:     categoryRequirementRepo,
		categoryRelationshipRepo:    categoryRelationshipRepo,
		categoryCourseRepo:          categoryCourseRepo,
		categoryCourseRequisiteRepo: categoryCourseRequisiteRepo,
		courseDetailRepo:            courseDetailRepo,
	}
}

func (s *categoryService) GetTypes() ([]model.SysCategoryType, error) {
	types, err := s.categoryTypeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return types, nil
}

func (s *categoryService) CreateType(types model.SysCategoryType) error {
	return s.categoryTypeRepo.Create(&types)
}

func (s *categoryService) UpdateType(types model.SysCategoryType) error {
	return s.categoryTypeRepo.Update(&types)
}

func (s *categoryService) DeleteType(typeID int) error {
	return s.categoryTypeRepo.Delete(typeID)
}

func (s *categoryService) getCategoryHelper(startCate model.SysCategory, cates []model.SysCategory) (*dto.Category, error) {
	category, err := mapper.Mapper[model.SysCategory, dto.Category](startCate)
	if err != nil {
		return nil, err
	}

	requirements, err := mapper.MapSlice[model.SysCategoryRequirement, dto.CategoryRequirement](startCate.Requirements)
	if err != nil {
		return nil, err
	}
	category.Requirements = requirements

	relationships, err := mapper.MapSlice[model.SysCategoryRelationship, dto.CategoryRelationship](startCate.Relationships)
	if err != nil {
		return nil, err
	}
	category.Relationships = relationships

	childCategories := lo.Filter(cates, func(cate model.SysCategory, _ int) bool {
		_, ok := lo.Find(relationships, func(relationship dto.CategoryRelationship) bool {
			return relationship.ChildCategoryID == cate.ID
		})
		return ok
	})

	for _, childCate := range childCategories {
		childCategory, err := s.getCategoryHelper(childCate, cates)
		if err != nil {
			return nil, err
		}
		category.ChildCategories = append(category.ChildCategories, *childCategory)
	}

	courses, err := mapper.MapSlice[model.SysCategoryCourse, dto.CategoryCourse](startCate.Courses)
	if err != nil {
		return nil, err
	}

	for i, course := range courses {
		requisites, err := mapper.MapSlice[model.SysCategoryCourseRequisite, dto.CategoryCourseRequisite](startCate.Courses[i].Requisites)
		if err != nil {
			return nil, err
		}
		course.Requisites = requisites

		courseDetail, err := s.courseDetailRepo.GetByCourseNo(*course.CourseNo)
		if err != nil {
			return nil, err
		}
		course.Detail = *courseDetail
		course.Credit = courseDetail.Credit

		courses[i] = course
	}

	category.Courses = courses
	return category, nil
}

func (s *categoryService) GetByCurriculumID(curriculumID int) (*dto.Category, error) {
	sysCategories, err := s.categoryRepo.GetByCurriculumID(curriculumID)
	if err != nil {
		return nil, err
	}

	totalCategory, ok := lo.Find(sysCategories, func(cate model.SysCategory) bool {
		return cate.TypeID == 7
	})

	if !ok {
		return nil, nil
	}

	category, err := s.getCategoryHelper(totalCategory, sysCategories)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) Create(category dto.Category) error {
	sysCategory, err := mapper.Mapper[dto.Category, model.SysCategory](category)
	if err != nil {
		return err
	}
	return s.categoryRepo.Create(sysCategory)
}

func (s *categoryService) Update(category dto.Category) error {
	sysCategory, err := mapper.Mapper[dto.Category, model.SysCategory](category)
	if err != nil {
		return err
	}
	return s.categoryRepo.Update(sysCategory)
}

func (s *categoryService) Delete(categoryID int) error {
	return s.categoryRepo.Delete(categoryID)
}
