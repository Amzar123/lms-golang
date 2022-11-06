package assignments

import "mini-project/app/middlewares"

type AssignmentUsecase struct {
	assignmentRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewAssignmentUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &AssignmentUsecase{
		assignmentRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *AssignmentUsecase) GetAssignments() []Domain {
	return uu.assignmentRepository.GetAssignments()
}

func (uu *AssignmentUsecase) Create(assignmentDomain *Domain) Domain {
	return uu.assignmentRepository.Create(assignmentDomain)
}