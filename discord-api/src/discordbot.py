import discord
from discord.ext import commands
import config
from message import Message


intents = discord.Intents.all()
bot = commands.Bot(command_prefix="?", intents=intents)

@bot.event
async def on_ready():
    print("Ready!")


@bot.command()
async def get_id(ctx):
    if ctx.author.bot:
        return
    
    user_id = ctx.author.id
    await ctx.send(f"Your ID is: {user_id}")

@bot.command()
async def hello(ctx, name: str):
    await ctx.send(f"Hello! {name}")

@bot.command()
async def explanation(ctx):
    message = Message()
    await ctx.send(message.help())
