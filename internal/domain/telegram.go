package domain

type Message struct {
	ID     string
	ChatID string
	Text   string
}

type Chat struct {
	ID      string
	Name    string
	Message []*Message
}
