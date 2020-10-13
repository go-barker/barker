package types

type User struct {
	//Telegram first name
	FirstName string `json:"FirstName,omitempty"`
	//Telegram last name
	LastName    string `json:"LastName,omitempty"`
	DisplayName string `json:"DisplayName,omitempty"`
	//Telegram user name (e.g. @foo)
	UserName string `json:"UserName,omitempty"`
	//Telegram ID
	TelegramID int64 `json:"TelegramID,omitempty"`
	//ID of bot which this user belongs to
	BotID int64 `json:"BotID,omitempty"`
}
