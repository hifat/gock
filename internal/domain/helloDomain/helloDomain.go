package helloDomain

type HelloService interface {
	GetHello() string
}

type HelloRepository interface {
	GetHello() string
}
