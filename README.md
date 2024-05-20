# Telegram invite link generator

## Configuration

At the first start, if there is no configuration file, it will be generated

Specify in the configuration file `conf.json`:
- `Token` Bot token
- `Group` ID of the closed group where the bot is admin
- `MembersLimit` Maximum number of entries per link 
- `RequestsPerMinute` Number of requests per minute to the Telegram API 
- `Prefix` Text before the link to save to file 
- `Postfix` Text after the link to save to file 

```json
{
  "Token": "BOT_TOKEN",
  "Group": "ID_GROUP",
  "RequestsPerMinute": 20,
  "MembersLimit": 1,
  "Prefix": "TEXT",
  "Postfix": "TEXT"
}
```

## Uses
Download the [latest release](https://github.com/olexin-pro/telegram-invite-link-generator/releases)

Fill in the configuration file

Run the .exe file