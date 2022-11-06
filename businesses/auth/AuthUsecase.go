package auth

import "mini-project/app/middlewares"

type AuthUsecase struct {
	// studentRepository Repository
	// teacherRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

// func (uu *AuthUsecase) Login(authDomain *Domain) string {
// 	student := uu.studentRepository.GetByEmail(authDomain)
// 	teacher := uu.teacherRepository.GetByEmail(authDomain)

// 	if user.ID == 0 {
// 		return ""
// 	}

// 	token := uu.jwtAuth.GenerateToken(int(user.ID))

// 	return token
// }
