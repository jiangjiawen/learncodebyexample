# 这个是东西一直来，我自己这里准备好来接住。
from multiprocessing import Process, Queue, Manager, Pool
from time import sleep
import random
import queue
 
def get(q, d, doingordone):
    while True:
        info = q.get()
        if info not in doingordone:
            doingordone.append(info)
            print("using time {}".format(info))
            for k,v in d.items():
                if v==True:
                    doSourcehere = k
                    break
            d[doSourcehere]=False
            print("using source {}".format(doSourcehere))
            sleep(info)
            d[doSourcehere]=True
            print("end")
        else:
            continue

 
def put(q):
	while True:
		q.put(int(random.random()*10))
	print('put is done')
 
def main():
    print('main task start')
    d = Manager().dict({1:True,2:True})
    qDoingorDone = Manager().list([])
    qin = Queue()
    p1 = Process(target=put, args=(qin, ))
    p2 = Process(target=get, args=(qin, d, qDoingorDone))
    p3 = Process(target=get, args=(qin, d, qDoingorDone))

    # 最好加个delay 好像有同步竞争问题
    p1.start()
    p2.start()
    sleep(1)
    p3.start()
    p1.join()
    p2.join()
    p3.join()

    print('main task done')
 
if __name__ == '__main__':
	main()