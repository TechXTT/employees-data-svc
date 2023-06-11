package services

type Services interface {
	EmployeeService
	PositionService
	DepartmentService
}

type services struct {
	EmployeeService
	PositionService
	DepartmentService
}

func NewServices(employeeService EmployeeService, positionService PositionService, departmentService DepartmentService) Services {
	return &services{
		employeeService,
		positionService,
		departmentService,
	}
}
