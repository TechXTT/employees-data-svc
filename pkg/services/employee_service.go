package services

import (
	"context"
	"time"

	"github.com/TechXTT/employees-data-svc/pkg/models"
	"github.com/TechXTT/employees-data-svc/pkg/pb"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetEmployees(ctx context.Context, req *pb.GetEmployeesRequest) (*pb.GetEmployeesResponse, error)
	GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error)

	CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error)
	UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error)
	DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error)
}

type employeeService struct {
	*gorm.DB
}

func NewEmployeeService(db *gorm.DB) EmployeeService {
	return &employeeService{db}
}

func (s *employeeService) GetEmployees(ctx context.Context, req *pb.GetEmployeesRequest) (*pb.GetEmployeesResponse, error) {
	var employees []*models.Employee
	if err := s.DB.WithContext(ctx).Preload("Position").Preload("Position.Department").Find(&employees).Error; err != nil {
		return nil, err
	}
	var employeePBs []*pb.Employee
	for _, employee := range employees {
		employeePB := &pb.Employee{
			Id:        employee.ID.String(),
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			HireDate:  employee.HireDate.String(),
			Salary:    float32(employee.Salary),
			Position: &pb.Position{
				Id:    employee.Position.ID.String(),
				Title: employee.Position.Title,
				Department: &pb.Department{
					Id:   employee.Position.Department.ID.String(),
					Name: employee.Position.Department.Name,
				},
			},
		}
		employeePBs = append(employeePBs, employeePB)
	}

	return &pb.GetEmployeesResponse{
		Employees: employeePBs,
	}, nil
}

func (s *employeeService) GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.GetEmployeeResponse, error) {
	var employee models.Employee
	if err := s.DB.WithContext(ctx).First(&employee, req.Id).Error; err != nil {
		return nil, err
	}

	employeePB := &pb.Employee{
		Id:        employee.ID.String(),
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Email:     employee.Email,
		HireDate:  employee.HireDate.String(),
		Salary:    float32(employee.Salary),
		Position: &pb.Position{
			Id:    employee.Position.ID.String(),
			Title: employee.Position.Title,
			Department: &pb.Department{
				Id:   employee.Position.Department.ID.String(),
				Name: employee.Position.Department.Name,
			},
		},
	}

	return &pb.GetEmployeeResponse{
		Employee: employeePB,
	}, nil
}

func (s *employeeService) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {
	hireDate, err := time.Parse("2006-01-02", req.HireDate)
	if err != nil {
		return nil, err
	}

	employee := &models.Employee{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		HireDate:   hireDate,
		Salary:     float64(req.Salary),
		PositionID: req.PositionId,
	}

	if err := s.DB.WithContext(ctx).Create(employee).Error; err != nil {
		return nil, err
	}

	return &pb.CreateEmployeeResponse{
		Employee: &pb.Employee{
			Id:        employee.ID.String(),
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			HireDate:  employee.HireDate.String(),
			Salary:    float32(employee.Salary),
			Position: &pb.Position{
				Id:    employee.Position.ID.String(),
				Title: employee.Position.Title,
				Department: &pb.Department{
					Id:   employee.Position.Department.ID.String(),
					Name: employee.Position.Department.Name,
				},
			},
		},
	}, nil

}

func (s *employeeService) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	employee := &models.Employee{
		ID:         uuid.FromStringOrNil(req.Id),
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		HireDate:   time.Time{},
		Salary:     float64(req.Salary),
		PositionID: req.PositionId,
	}

	if err := s.DB.WithContext(ctx).Save(employee).Error; err != nil {
		return nil, err
	}

	return &pb.UpdateEmployeeResponse{
		Employee: &pb.Employee{
			Id:        employee.ID.String(),
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Email:     employee.Email,
			HireDate:  employee.HireDate.String(),
			Salary:    float32(employee.Salary),
			Position: &pb.Position{
				Id:    employee.Position.ID.String(),
				Title: employee.Position.Title,
				Department: &pb.Department{
					Id:   employee.Position.Department.ID.String(),
					Name: employee.Position.Department.Name,
				},
			},
		},
	}, nil
}

func (s *employeeService) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	if err := s.DB.WithContext(ctx).Delete(&models.Employee{}, req.Id).Error; err != nil {
		return nil, err
	}

	return &pb.DeleteEmployeeResponse{}, nil
}
