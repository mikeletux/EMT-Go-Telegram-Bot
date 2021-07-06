## EMT-Go-Telegram-Bot
Small telegram bot for getting information regarding EMT bus lines from Madrid City. 

After compiled, the following environment vars must be set before starting the bot:
  - **TELEGRAM_API_TOKEN**: Token given by [botfather](https://t.me/botfather) 
  - **TELEGRAM_ALLOWED_USERS**: list of users allowed to use this bot separated by commands (i.e: "user1,user2,user3")
  - **EMT_ENDPOINT**: Mobility Labs EMT Madrid endpoint (*https://openapi.emtmadrid.es/v2*)
  - **EMT_XCLIENTID**: XClientID provided by Mobility Labs EMT Madrid
  - **EMT_PASSKEY**: PassKey provided by Mobility Labs EMT Madrid

/Miguel Sama (miguelsamamerino@gmail.com) 2021