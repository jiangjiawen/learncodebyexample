# 这个是东西准备好了，给多进程
import multiprocessing
import time
import random

def func(msg, timenum, d):
    print(msg, "using time {}".format(timenum))
    for k,v in d.items():
        if v==True:
            doSourcehere = k
            break
    d[doSourcehere]=False
    print("using source {}".format(doSourcehere))
    time.sleep(timenum)
    d[doSourcehere]=True
    print("end")


if __name__ == '__main__':
    pool = multiprocessing.Pool(processes=2)
    d = multiprocessing.Manager().dict({1:True,2:True})
    for i in range(9):
        msg = "hello {0}".format(i)
        pool.apply_async(func,(msg,int(random.random()*10), d))

    print("start.....")
    start_time = time.time()
    pool.close()
    pool.join()
    print("all done")
    print("--- %s seconds ---" % (time.time() - start_time))