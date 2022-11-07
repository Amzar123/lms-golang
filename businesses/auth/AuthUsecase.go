package auth

import "mini-project/app/middlewares"

type AuthUsecase struct {
	teacherRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewAuthUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &AuthUsecase{
		teacherRepository: ur,
		jwtAuth:        jwtAuth,
	}
}


func (uu *AuthUsecase) Login(authDomain *Domain) string {
	// student := uu.studentRepository.GetByEmail(authDomain)
	teacher := uu.teacherRepository.GetByEmail(authDomain)

	if teacher.Email == "" {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(int(2))

	return token
}