package types

type Bot struct {
	ID    int64
	Title string `binding:"required"`
	Token string `binding:"required"`
}
