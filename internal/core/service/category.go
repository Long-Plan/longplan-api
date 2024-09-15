package service

import (
	"github.com/Long-Plan/longplan-api/internal/core/domain"
	"github.com/Long-Plan/longplan-api/internal/core/dto"
	"github.com/Long-Plan/longplan-api/internal/core/model"
	"github.com/Long-Plan/longplan-api/internal/core/port"
	"github.com/Long-Plan/longplan-api/pkg/mapper"
)

type categoryService struct {
	categoryRepo                port.SysCategoryRepo
	categoryTypeRepo            port.SysCategoryTypeRepo
	categoryRequirementRepo     port.SysCategoryRequirementRepo
	categoryRelationshipRepo    port.SysCategoryRelationshipRepo
	categoryCourseRepo          port.SysCategoryCourseRepo
	categoryCourseRequisiteRepo port.SysCategoryCourseRequisiteRepo
}

func NewCategoryService(
	categoryRepo port.SysCategoryRepo,
	categoryTypeRepo port.SysCategoryTypeRepo,
	categoryRequirementRepo port.SysCategoryRequirementRepo,
	categoryRelationshipRepo port.SysCategoryRelationshipRepo,
	categoryCourseRepo port.SysCategoryCourseRepo,
	categoryCourseRequisiteRepo port.SysCategoryCourseRequisiteRepo,
) domain.CategoryService {
	return &categoryService{
		categoryRepo:                categoryRepo,
		categoryTypeRepo:            categoryTypeRepo,
		categoryRequirementRepo:     categoryRequirementRepo,
		categoryRelationshipRepo:    categoryRelationshipRepo,
		categoryCourseRepo:          categoryCourseRepo,
		categoryCourseRequisiteRepo: categoryCourseRequisiteRepo,
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

func (s *categoryService) GetByCurriculumID(curriculumID int) ([]dto.Category, error) {
	var categories []dto.Category
	sysCategories, err := s.categoryRepo.GetByCurriculumID(curriculumID)
	if err != nil {
		return nil, err
	}
	for _, sysCategory := range sysCategories {
		category, err := mapper.Mapper[model.SysCategory, dto.Category](sysCategory)
		if err != nil {
			return nil, err
		}

		requirements, err := mapper.MapSlice[model.SysCategoryRequirement, dto.CategoryRequirement](sysCategory.Requirements)
		if err != nil {
			return nil, err
		}
		category.Requirements = requirements

		relationships, err := mapper.MapSlice[model.SysCategoryRelationship, dto.CategoryRelationship](sysCategory.Relationships)
		if err != nil {
			return nil, err
		}
		category.Relationships = relationships

		courses, err := mapper.MapSlice[model.SysCategoryCourse, dto.CategoryCourse](sysCategory.Courses)
		if err != nil {
			return nil, err
		}

		for i, course := range courses {
			requisites, err := mapper.MapSlice[model.SysCategoryCourseRequisite, dto.CategoryCourseRequisite](sysCategory.Courses[i].Requisites)
			if err != nil {
				return nil, err
			}

			course.Requisites = requisites
			courses[i] = course
		}

		category.Courses = courses
		categories = append(categories, *category)
	}
	return categories, nil
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
