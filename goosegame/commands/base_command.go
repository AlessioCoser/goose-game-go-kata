package commands

type BaseCmd struct {}

func (c *BaseCmd) CanHandle(command string) bool {
	return true
}

func (c *BaseCmd) Execute(command string) string {
	return "Command not found"
}
