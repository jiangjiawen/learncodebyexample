import os
import time
from concurrent.futures.process import ProcessPoolExecutor

import uvicorn
from multiprocessing import Pool
from fastapi import FastAPI


def simple_routine(sleep_for):
    print(f"PID {os.getpid()} has sleep time: {sleep_for}")
    time.sleep(sleep_for)
    return "done"

app = FastAPI()

@app.post("/test-endpoint/?")
async def test_endpoint():
    print(f"main process: {os.getpid()}")

    START_TIME = time.time()
    STOP_TIME = START_TIME + 2

    pool = ProcessPoolExecutor(max_workers=3)
    futures = [
        pool.submit(simple_routine, [1]),
        pool.submit(simple_routine, [1]),
        pool.submit(simple_routine, [10]),
    ]
    results = []
    for fut in futures:
        remains = max(STOP_TIME - time.time(), 0)
        try:
            results.append(fut.get(timeout = remains))
        except:
            results.append("not done")

    # terminate the entire pool
    pool.shutdown(wait=False)
    print("exiting at: ", int(time.time() - START_TIME))
    return "True"


if __name__=="__main__":
    uvicorn.run(app, host="127.0.0.1", port=5900)