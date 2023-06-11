package services

import (
	"context"

	"github.com/TechXTT/employees-data-svc/pkg/models"
	"github.com/TechXTT/employees-data-svc/pkg/pb"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type PositionService interface {
	GetPositions(ctx context.Context, req *pb.GetPositionsRequest) (*pb.GetPositionsResponse, error)
	GetPosition(ctx context.Context, req *pb.GetPositionRequest) (*pb.GetPositionResponse, error)
	GetPositionByTitle(ctx context.Context, req *pb.GetPositionByTitleRequest) (*pb.GetPositionByTitleResponse, error)

	CreatePosition(ctx context.Context, req *pb.CreatePositionRequest) (*pb.CreatePositionResponse, error)
	UpdatePosition(ctx context.Context, req *pb.UpdatePositionRequest) (*pb.UpdatePositionResponse, error)
	DeletePosition(ctx context.Context, req *pb.DeletePositionRequest) (*pb.DeletePositionResponse, error)
}

type positionService struct {
	*gorm.DB
}

func NewPositionService(db *gorm.DB) PositionService {
	return &positionService{db}
}

func (s *positionService) GetPositions(ctx context.Context, req *pb.GetPositionsRequest) (*pb.GetPositionsResponse, error) {
	var positions []*models.Position
	if err := s.DB.WithContext(ctx).Preload("Department").Find(&positions).Error; err != nil {
		return nil, err
	}
	var positionPBs []*pb.Position
	for _, position := range positions {
		positionPB := &pb.Position{
			Id:    position.ID.String(),
			Title: position.Title,
			Department: &pb.Department{
				Id:   position.Department.ID.String(),
				Name: position.Department.Name,
			},
		}
		positionPBs = append(positionPBs, positionPB)
	}

	return &pb.GetPositionsResponse{
		Positions: positionPBs,
	}, nil
}

func (s *positionService) GetPosition(ctx context.Context, req *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	var position models.Position
	if err := s.DB.WithContext(ctx).Where("id = ?", req.Id).First(&position).Error; err != nil {
		return nil, err
	}

	positionPB := &pb.Position{
		Id:    position.ID.String(),
		Title: position.Title,
		Department: &pb.Department{
			Id:   position.Department.ID.String(),
			Name: position.Department.Name,
		},
	}

	return &pb.GetPositionResponse{
		Position: positionPB,
	}, nil
}

func (s *positionService) GetPositionByTitle(ctx context.Context, req *pb.GetPositionByTitleRequest) (*pb.GetPositionByTitleResponse, error) {
	var position models.Position
	if err := s.DB.WithContext(ctx).Where("title = ?", req.Title).First(&position).Error; err != nil {
		return nil, err
	}

	positionPB := &pb.Position{
		Id:    position.ID.String(),
		Title: position.Title,
		Department: &pb.Department{
			Id:   position.Department.ID.String(),
			Name: position.Department.Name,
		},
	}

	return &pb.GetPositionByTitleResponse{
		Position: positionPB,
	}, nil
}

func (s *positionService) CreatePosition(ctx context.Context, req *pb.CreatePositionRequest) (*pb.CreatePositionResponse, error) {
	position := &models.Position{
		Title:        req.Title,
		DepartmentID: req.DepartmentId,
	}
	if err := s.DB.WithContext(ctx).Create(position).Error; err != nil {
		return nil, err
	}

	return &pb.CreatePositionResponse{
		Position: &pb.Position{
			Id:    position.ID.String(),
			Title: position.Title,
			Department: &pb.Department{
				Id:   position.Department.ID.String(),
				Name: position.Department.Name,
			},
		},
	}, nil
}

func (s *positionService) UpdatePosition(ctx context.Context, req *pb.UpdatePositionRequest) (*pb.UpdatePositionResponse, error) {
	position := &models.Position{
		ID:           uuid.FromStringOrNil(req.Id),
		Title:        req.Title,
		DepartmentID: req.DepartmentId,
	}
	if err := s.DB.WithContext(ctx).Save(position).Error; err != nil {
		return nil, err
	}

	return &pb.UpdatePositionResponse{
		Position: &pb.Position{
			Id:    position.ID.String(),
			Title: position.Title,
			Department: &pb.Department{
				Id:   position.Department.ID.String(),
				Name: position.Department.Name,
			},
		},
	}, nil
}

func (s *positionService) DeletePosition(ctx context.Context, req *pb.DeletePositionRequest) (*pb.DeletePositionResponse, error) {
	if err := s.DB.WithContext(ctx).Where("id = ?", req.Id).Delete(&models.Position{}).Error; err != nil {
		return nil, err
	}

	return &pb.DeletePositionResponse{}, nil
}
