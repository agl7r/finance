package console

type Command interface {
	Names() []string
	Execute(args []string) error
}
