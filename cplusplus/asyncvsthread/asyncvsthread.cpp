// https://eli.thegreenplace.net/2016/the-promises-and-challenges-of-stdasync-task-based-parallelism-in-c11/
#include <iostream>
#include <algorithm>
#include <numeric>
#include <vector>
#include <thread>
#include <future>
#include <functional>
#include <ctime>

using namespace std;

void accumulate_block_worker(int* data, size_t count, int* result){
    *result = accumulate(data, data + count, 0);
}

void use_worker_in_std_thread(){
    vector<int> v{1,2,3,4,5,6,7,8};
    int result;
    thread worker(accumulate_block_worker, v.data(), v.size(), &result);
    worker.join();
    cout << "result is: " << result << endl;
}

int accumulate_block_worker_ret(int* data, size_t count){
    return accumulate(data, data+count, 0);
}

void use_worker_in_async(){
    vector<int> v{1,2,3,4,5,6,7,8};
    future<int> fut = async(
        launch::async, accumulate_block_worker_ret, v.data(), v.size());
    cout << "result is: " << fut.get() <<endl;
}

vector<thread> launch_split_wokers_with_std_thread(vector<int>& v, vector<int>* results){
    vector<thread> threads;
    threads.emplace_back(accumulate_block_worker, v.data(), v.size()/2, &((*results)[0]));
    threads.emplace_back(accumulate_block_worker, v.data() + v.size()/2, v.size()/2, &((*results)[1]));
    return threads;
}

vector<future<int>> launch_split_workers_with_std_async(vector<int>& v){
    vector<future<int>> futures;
    futures.push_back(async(launch::async, accumulate_block_worker_ret, v.data(), v.size()/2));
    futures.push_back(async(launch::async, accumulate_block_worker_ret, v.data()+v.size()/2, v.size()/2));
    return futures;
}

int main(){
    use_worker_in_std_thread();
    use_worker_in_async();

    clock_t start = clock();
    vector<int> v{1,2,3,4,5,6,7,8};
    vector<int> results(2,0);
    vector<thread> threads = launch_split_wokers_with_std_thread(v, &results);
    for(auto& t: threads) {
        t.join();
    }
    cout<<"result is :"<< results[0]<< results[1] <<endl;
    clock_t end = clock();
    cout << "time is: " << double(end - start)/ CLOCKS_PER_SEC << "ms" << endl;

     clock_t start2 = clock();
    vector<int> v2{1,2,3,4,5,6,7,8};
    
    vector<future<int>> futures=launch_split_workers_with_std_async(v2);
    cout<<"results is :" << futures[0].get() << "and" << futures[1].get() << endl;
    clock_t end2 = clock();
    cout << "time is: " << double(end2 - start2)/ CLOCKS_PER_SEC << "ms" << endl;

    return 0;
}