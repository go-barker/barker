# Barker - Telegram bot broadcasting platform

## API

`POST /bot { Title, Token }` - create a bot

`PUT /bot/:ID { Title, Token }` - update a bot 

`GET /bot/:ID` - get a bot

`POST /bot/:BotID/campaign {Title string, Message string, Active bool}` create a campaign

`GET /bot/:BotID/campaign/:CampaignID {Title string, Message string, Active bool}` get a campaign

`PUT /bot/:BotID/user {	TelegramID int64, FirstName string, LastName string, DisplayName string, UserName string }` - create or update a user

`GET /bot/:BotID/user/:UserID` - get a user

`POST /bot/:BotID/campaign/:CampaignID/delivery` - create a delivery