import requests

from discordbot import bot
import config


def main() -> None:
    TOKEN = config.DISCORD_TOKEN
    bot.run(TOKEN)


if __name__ == "__main__":
    main()
