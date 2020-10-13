package types

type Bot struct {
	ID    int64  `json:"ID,omitempty"`
	Title string `binding:"required" json:"Title,omitempty"`
	Token string `binding:"required" json:"Token,omitempty"`
}
