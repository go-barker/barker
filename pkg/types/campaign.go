package types

type Campaign struct {
	ID      int64  `json:"ID,omitempty"`
	BotID   int64  `json:"BotID,omitempty"`
	Title   string `binding:"required" json:"Title,omitempty"`
	Message string `binding:"required" json:"Message,omitempty"`
	Active  bool   `json:"Active,omitempty"`
}
