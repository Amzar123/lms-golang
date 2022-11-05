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

// func (uu *UserUsecase) Login(userDomain *Domain) string {
// 	user := uu.userRepository.GetByEmail(userDomain)

// 	if user.ID == 0 {
// 		return ""
// 	}

// 	token := uu.jwtAuth.GenerateToken(int(user.ID))

// 	return token
// }

func (uu *TeacherUsecase) GetTeachers() []Domain {
	return uu.teacherRepository.GetTeachers()
}