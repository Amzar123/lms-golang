package students

import "mini-project/app/middlewares"

type StudentUsecase struct {
	studentRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewStudentUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &StudentUsecase{
		studentRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *StudentUsecase) Create(studentDomain *Domain) Domain {
	return uu.studentRepository.Create(studentDomain)
}

func (uu *StudentUsecase) Login(studentDomain *Domain) string {
	student := uu.studentRepository.GetByEmail(studentDomain)

	if student.NISN == "" {
		return " ";
	}

	token := uu.jwtAuth.GenerateToken(2)

	return token
}

// func (uu *StudentUsecase) GetModules() []Domain {
// 	return uu.studentRepository.GetModules()
// }