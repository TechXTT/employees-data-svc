package services

import (
	"context"

	"github.com/TechXTT/employees-data-svc/pkg/models"
	"github.com/TechXTT/employees-data-svc/pkg/pb"
	"gorm.io/gorm"
)

type DepartmentService interface {
	GetDepartments(ctx context.Context, req *pb.GetDepartmentsRequest) (*pb.GetDepartmentsResponse, error)
	GetDepartment(ctx context.Context, req *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error)

	CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error)
	UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error)
	DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error)
}

type departmentService struct {
	*gorm.DB
}

func NewDepartmentService(db *gorm.DB) DepartmentService {
	return &departmentService{db}
}

func (s *departmentService) GetDepartments(ctx context.Context, req *pb.GetDepartmentsRequest) (*pb.GetDepartmentsResponse, error) {
	var departments []*models.Department
	if err := s.DB.WithContext(ctx).Find(&departments).Error; err != nil {
		return nil, err
	}
	var departmentPBs []*pb.Department
	for _, department := range departments {
		departmentPB := &pb.Department{
			Id:   department.ID.String(),
			Name: department.Name,
		}
		departmentPBs = append(departmentPBs, departmentPB)
	}

	return &pb.GetDepartmentsResponse{
		Departments: departmentPBs,
	}, nil
}

func (s *departmentService) GetDepartment(ctx context.Context, req *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error) {
	var department models.Department
	if err := s.DB.WithContext(ctx).First(&department, req.Id).Error; err != nil {
		return nil, err
	}

	departmentPB := &pb.Department{
		Id:   department.ID.String(),
		Name: department.Name,
	}

	return &pb.GetDepartmentResponse{
		Department: departmentPB,
	}, nil
}

func (s *departmentService) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.CreateDepartmentResponse, error) {
	department := &models.Department{
		Name: req.Name,
	}

	if err := s.DB.WithContext(ctx).Create(department).Error; err != nil {
		return nil, err
	}
	return &pb.CreateDepartmentResponse{
		Department: &pb.Department{
			Id:   department.ID.String(),
			Name: department.Name,
		},
	}, nil
}

func (s *departmentService) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.UpdateDepartmentResponse, error) {
	if err := s.DB.WithContext(ctx).Model(&models.Department{}).Where("id = ?", req.Id).Updates(&models.Department{Name: req.Name}).Error; err != nil {
		return nil, err
	}
	return &pb.UpdateDepartmentResponse{
		Department: &pb.Department{
			Id:   req.Id,
			Name: req.Name,
		},
	}, nil
}

func (s *departmentService) DeleteDepartment(ctx context.Context, req *pb.DeleteDepartmentRequest) (*pb.DeleteDepartmentResponse, error) {
	if err := s.DB.WithContext(ctx).Delete(&models.Department{}, req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteDepartmentResponse{}, nil
}
