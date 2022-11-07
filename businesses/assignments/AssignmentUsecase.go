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

func (uu *AssignmentUsecase) GetByID(id string) Domain {
	return uu.assignmentRepository.GetByID(id)
}

func (nu *AssignmentUsecase) Update(id string, assignmentDomain *Domain) Domain {
	return nu.assignmentRepository.Update(id, assignmentDomain)
}

func (nu *AssignmentUsecase) Delete(id string) bool {
	return nu.assignmentRepository.Delete(id)
}