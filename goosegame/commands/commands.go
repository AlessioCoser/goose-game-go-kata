package commands

type Command interface {
	CanHandle(string)bool
	Execute(string)string
}

func All() []Command {
	return []Command{
		&BaseCmd{},
	}
}
