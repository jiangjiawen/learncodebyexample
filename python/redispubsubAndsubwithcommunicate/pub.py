import redis
import random
import pickle
import time

r = redis.Redis(host='localhost', port=6379, db=0)
channelname = 'test'
while True:
    for i in range(10):
        NUM_RandomInt = int(random.randint(0,10))
        time.sleep(random.randint(4,6))
        print({i: NUM_RandomInt})
        messageData = pickle.dumps({i: NUM_RandomInt})
        r.publish(channelname, messageData)