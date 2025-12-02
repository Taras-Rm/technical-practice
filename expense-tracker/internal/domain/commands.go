package domain

type CLICommand string

const (
	COMMAND_ADD     CLICommand = "add"
	COMMAND_LIST    CLICommand = "list"
	COMMAND_SUMMARY CLICommand = "summary"
	COMMAND_DELETE  CLICommand = "delete"
)
