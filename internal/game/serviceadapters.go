package game

type Service interface {
	Start() error
	Stop() error
}
