package main

import (
	"fmt"
	"log"
	"net"

	"github.com/TechXTT/employees-data-svc/pkg/config"
	"github.com/TechXTT/employees-data-svc/pkg/db"
	"github.com/TechXTT/employees-data-svc/pkg/pb"
	"github.com/TechXTT/employees-data-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	DB := db.Init(c.DatabaseURL)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	fmt.Println("Data service is running on port", c.Port)

	es := services.NewEmployeeService(DB)
	ds := services.NewDepartmentService(DB)
	ps := services.NewPositionService(DB)
	s := services.NewServices(es, ps, ds)

	grpcServer := grpc.NewServer()

	pb.RegisterDataServiceServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
