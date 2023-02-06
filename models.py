import requests
import time

# overrrided class 
class Session(requests.Session):
    def get(self, *args, **kwargs):
        for _ in range(10):
            try:
                return super(Session, self).get(*args, **kwargs)
            except Exception as e:
                time.sleep(1)

    def post(self, *args, **kwargs):
        for _ in range(10):
            try:
                return super(Session, self).post(*args, **kwargs)
            except Exception as e:
                time.sleep(1)

def TwitterSession(user_agent, cookie, csrf_token, proxy=None):
    session = Session()
    session.headers.update(
        {
            'cookie': cookie,
            'authorization': 'Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA',
            'user-agent': user_agent,
            'origin': 'https://twitter.com',
            'x-csrf-token': csrf_token
        }
    )
    if proxy:
        session.proxies.update(
            {
                'http': f'http://{proxy}',
                'https': f'http://{proxy}'
            }
        )
    return session

class Tweet:
    def __init__(self, session, id=None):
        self.session = session
        self.id = id

    def get_details(self):
        try:
            response = self.session.get(f'https://api.twitter.com/2/timeline/conversation/{self.id}.json').json()
            return response
        except:
            return None

    
    def delete(self):
        payload = {"id": self.id}
        self.session.post('https://api.twitter.com/1.1/statuses/destroy.json', data=payload).json()

    def retweet(self):
        payload = {"id": self.id}
        self.session.post('https://twitter.com/i/api/1.1/statuses/retweet.json', data=payload).json()

    def undo_retweet(self):
        payload = {"id": self.id}
        self.session.post('https://twitter.com/i/api/1.1/statuses/unretweet.json', data=payload).json()

    def like(self):
        payload = {"id": self.id}
        self.session.post('https://twitter.com/i/api/1.1/favorites/create.json', data=payload).json()
    
    def undo_like(self):
        payload = {"id": self.id}
        self.session.post('https://twitter.com/i/api/1.1/favorites/destroy.json', data=payload).json()
    
class User:
    def __init__(self, session, id):
        self.session = session
        self.id = id

    def get_details(self):
        response = self.session.get(f'https://api.twitter.com/1.1/users/lookup.json?user_id={self.id}').json()
        return response

    def get_last_tweets(self, count):
        tweet_ids = []
        response = self.session.get(f'https://twitter.com/i/api/graphql/WZT7sCTrLvSOaWOXLDsWbQ/UserTweets?variables={{"userId":"{self.id}","count":{count},"includePromotedContent":true,"withQuickPromoteEligibilityTweetFields":true,"withSuperFollowsUserFields":true,"withDownvotePerspective":false,"withReactionsMetadata":false,"withReactionsPerspective":false,"withSuperFollowsTweetFields":true,"withVoice":true,"withV2Timeline":true,"__fs_dont_mention_me_view_api_enabled":false,"__fs_interactive_text_enabled":false,"__fs_responsive_web_uc_gql_enabled":false}}')
        try:
            response_json = response.json()
        except:
            print(response.text)

        try:
            entries = response_json["data"]["user"]["result"]["timeline_v2"]["timeline"]["instructions"][0]["entries"]
        except:
            entries = response_json["data"]["user"]["result"]["timeline_v2"]["timeline"]["instructions"][1]["entries"]

        for entry in entries:
            try:
                legacy = entry["content"]["itemContent"]["tweet_results"]["result"]["legacy"]
                id_str = legacy["id_str"]
                tweet_ids.append(id_str)
            except:
                try:
                    for item in entry["content"]["items"]:
                        try:
                            id_str = item["item"]["itemContent"]["tweet_results"]["result"]["rest_id"]
                            tweet_ids.append(id_str)
                        except:
                            pass
                except:
                    pass
        return tweet_ids