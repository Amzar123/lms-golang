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

// func (uu *UserUsecase) Login(userDomain *Domain) string {
// 	user := uu.userRepository.GetByEmail(userDomain)

// 	if user.ID == 0 {
// 		return ""
// 	}

// 	token := uu.jwtAuth.GenerateToken(int(user.ID))

// 	return token
// }

// func (uu *UserUsecase) GetAllUser() []Domain {
// 	return uu.userRepository.GetAllUser()
// }