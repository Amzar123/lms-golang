package teachers

import "mini-project/app/middlewares"

type TeacherUsecase struct {
	teacherRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewTeacherUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &TeacherUsecase{
		teacherRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *TeacherUsecase) CreateTeacher(teacherDomain *Domain) Domain {
	return uu.teacherRepository.CreateTeacher(teacherDomain)
}

func (uu *TeacherUsecase) GetByID(id string) Domain {
	return uu.teacherRepository.GetByID(id)
}

func (nu *TeacherUsecase) Update(id string, teacherDomain *Domain) Domain {
	return nu.teacherRepository.Update(id, teacherDomain)
}

func (nu *TeacherUsecase) Delete(id string) bool {
	return nu.teacherRepository.Delete(id)
}

func (uu *TeacherUsecase) GetTeachers() []Domain {
	return uu.teacherRepository.GetTeachers()
}