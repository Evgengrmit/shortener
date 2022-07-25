package link

type Link struct {
	Data string
}

type LinkStorage interface {
	Add(original string) (string, error)
	Get(short string) (string, error)
}
