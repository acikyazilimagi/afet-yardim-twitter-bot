from models import *
from __init__ import *
import requests
from  constants import *


session = TwitterSession(user_agent, cookie, csrf_token)
client = Twitter(session)
tweet = client.get_tweet('1622652066635055135') # get tweet id
tweet.retweet()