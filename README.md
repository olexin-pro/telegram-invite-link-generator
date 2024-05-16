# Telegram invite link generator

## Uses

At the first start, if there is no configuration file, it will be generated

Specify in the configuration file:
- Bot token: `Token`
- ID of the closed group where the bot is admin: `Group`
- Maximum number of entries per link: `MembersLimit`
- Number of requests per minute to the Telegram API: `RequestsPerMinute`

```json
{
  "Token": "BOT_TOKEN",
  "Group": "ID_GROUP",
  "RequestsPerMinute": 20,
  "MembersLimit": 1
}
```
