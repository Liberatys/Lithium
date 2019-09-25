package logging

type Formater interface {
	Format(string) string
	Validate(string) bool
}

type Default struct {
}
