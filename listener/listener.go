package listener

type Listener interface {
	PreRun() error

	PostRun() error

	GetPriority() int
}
