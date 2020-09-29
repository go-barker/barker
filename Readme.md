# Barker - Telegram bot broadcasting platform

## API

`POST /bot { Title, Token }` - create a bot

`PUT /bot/:ID { Title, Token }` - update a bot 

`GET /bot/:ID` - get a bot

`POST /bot/:BotID/campaign {Title string, Message string, Active bool}` create a campaign

`GET /bot/:BotID/campaign/:ID {Title string, Message string, Active bool}` get a campaign

`PUT /bot/:BotID/user {...}` - create or update a user

`POST /bot/:BotID/campaign/:CampaignID/delivery` - create a delivery