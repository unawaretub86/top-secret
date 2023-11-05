package ports

type MessagePort interface {
	GetMessage(string, ...[]string) string
}
