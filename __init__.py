import time
import re
from models import TwitterSession, Tweet, User

class Twitter:
    class Upload:
        def __init__(self, session):
            self.session = session
       
    def __init__(self, session):
        self.session = session

    def get_tweet(self, id):
        return Tweet(self.session, id)
    
    def get_user(self, id):
        return User(self.session, id)

    def get_user_id(self, username: str):
        url = (
            'https://api.twitter.com/graphql/-xfUfZsnR_zqjFd-IfrN5A/UserByScreenName?variables={"screen_name":"'
            + username + '","withHighlightedLabel":true}'
        )
        response = self.session.get(url)
        try:
            data = response.json()
            user_id = data["data"]["user"]["rest_id"]
            return user_id
        except:
            return False

    def create_tweet(self, text: str, reply_to=None, media_data=None):
        media_id = None

        if media_data:
            up = self.Upload(self.session)
            up.import_data(media_data, 'tweet_image')
            media_id = up.upload()
            time.sleep(30)

        payload = {
            "status": text,
            "in_reply_to_status_id": reply_to,
            "batch_mode": "subsequent",
            "media_ids": media_id
        }

        response = self.session.post('https://api.twitter.com/1.1/statuses/update.json', data=payload).json()
        try:
            tweet_id = response["id_str"]
            return tweet_id
        except:
            print(response)


    def send_message(self, text: str, conversation_id: str, request_id: str):
        payload = {
            "text": text,
            "conversation_id": conversation_id,
            "request_id": request_id
        }
        response = self.session.post('https://twitter.com/i/api/1.1/dm/new.json', data=payload).json()
        return response
    
    def get_inbox(self):
        response = self.session.get('https://twitter.com/i/api/1.1/dm/inbox_initial_state.json').json()
        return response
    
    def get_untrusted_messages(self, max_id: str):
        response = self.session.get(f'https://twitter.com/i/api/1.1/dm/inbox_timeline/untrusted.json?max_id={max_id}')
        return response.json()
    
    def update_username(self, username: str):
        payload = {"screen_name": username}
        self.session.post('https://twitter.com/i/api/1.1/account/settings.json', data=payload)

    

