import redis
import pickle
from queue import Queue
from threading import Thread
import random
import time

channelname = 'test'
r=redis.Redis(host='localhost', port=6379, db=0)
p=r.publish()
p.psubscribe(channelname)
p.parse_response()


def producer(out_q):
    while True:
        msgdata = p.get_message()
        if msgdata is not None:
            getDict = pickle.loads(msgdata['data'])
            out_q.put(getDict)

def consumer(in_q):
    store = []
    while True:
        data = in_q.get()
        if data not in store:
            print(data)
            time.sleep(random.randint(1,7))
            store.append(data)

q = Queue()
t1 = Thread(target=consumer, args=(q,))
t2 = Thread(target=producer, args=(q,))
t1.start()
t2.start()