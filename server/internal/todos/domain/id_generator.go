package domain

//go:generate moq -out id_generator_mock.go . IDGenerator
type IDGenerator interface {
	Generate() string
}
