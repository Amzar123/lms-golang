package modules

import "mini-project/app/middlewares"

type ModuleUsecase struct {
	moduleRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewModuleUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &ModuleUsecase{
		moduleRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *ModuleUsecase) GetModules() []Domain {
	return uu.moduleRepository.GetModules()
}

func (uu *ModuleUsecase) Create(moduleDomain *Domain) Domain {
	return uu.moduleRepository.Create(moduleDomain)
}

func (uu *ModuleUsecase) GetByID(id string) Domain {
	return uu.moduleRepository.GetByID(id)
}

func (nu *ModuleUsecase) Update(id string, moduleDomain *Domain) Domain {
	return nu.moduleRepository.Update(id, moduleDomain)
}

func (nu *ModuleUsecase) Delete(id string) bool {
	return nu.moduleRepository.Delete(id)
}