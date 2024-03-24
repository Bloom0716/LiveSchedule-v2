import discord
from discord.ext import commands
import config


TOKEN = config.DISCORD_TOKEN

intents = discord.Intents.all()
bot = commands.Bot(command_prefix="/", intents=intents)

@bot.event
async def on_ready():
    print("Ready!")

@bot.command()
async def hello(ctx, name: str):
    await ctx.send(f"Hello! {name}")

@bot.command()
async def add(ctx, num1: int, num2: int):
    result: int = num1 + num2
    await ctx.send(f"{result}")

@bot.command()
async def ping(ctx):
    await ctx.send("Pong!")

bot.run(TOKEN)
