package types

type User struct {
	//Telegram first name
	FirstName string
	//Telegram last name
	LastName    string
	DisplayName string
	//Telegram user name (e.g. @foo)
	UserName string
	//Telegram ID
	TelegramID int64
	//ID of bot which this user belongs to
	BotID int64
}
