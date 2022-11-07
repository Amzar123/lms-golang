package auth

type Domain struct {
	Email     string
	Password  string
}

type Usecase interface {
	Login(authDomain *Domain) string
}

type Repository interface {
	GetByEmail(authDomain *Domain) Domain
}