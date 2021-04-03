//https://juejin.cn/post/6945319439772434469

const PENDING = "pending";
const FULFILLED = "fulfilled";
const REJECTED = "rejected";

class MyPromise {
  constructor(executor) {
    try {
      executor(this.resolve, this.reject);
    } catch (error) {
      this.reject(error);
    }
  }

  status = PENDING;
  value = null;
  reason = null;

  onFulfilledCallbacks = [];
  onRejectedCallbacks = [];

  resolve = (value) => {
    if (this.status == PENDING) {
      this.status = FULFILLED;
      this.value = value;
      while (this.onFulfilledCallbacks.length) {
        this.onFulfilledCallbacks.shift()(value);
      }
    }
  };

  reject = (reason) => {
    if (this.status === PENDING) {
      this.status = REJECTED;
      this.reason = reason;
      while (this.onFulfilledCallbacks.length) {
        this.onFulfilledCallbacks.shift()(reason);
      }
    }
  };

  then(onFulfilled, onRejected) {
    onFulfilled =
      typeof onFulfilled === "function" ? onFulfilled : (value) => value;
    onRejected =
      typeof onRejected === "function" ? onRejected : (value) => value;
    const promise2 = new MyPromise((resolve, reject) => {
      if (this.status === FULFILLED) {
        queueMicrotask(() => {
          try {
            const x = onFulfilled(this.value);
            resolvePromise(promise2, x, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      } else if (this.status === REJECTED) {
        queueMicrotask(() => {
          try {
            const x = onRejected(this.reason);
            resolvePromise(promise2, x, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      } else if (this.status === PENDING) {
        this.onFulfilledCallbacks.push(() => {
          queueMicrotask(() => {
            try {
              const x = onFulfilled(this.value);
              resolvePromise(promise2, x, resolve, reject);
            } catch (error) {
              reject(error);
            }
          });
        });
        this.onRejectedCallbacks.push(() => {
          queueMicrotask(() => {
            try {
              const x = onRejected(this.reason);
              resolvePromise(promise2, x, resolve, reject);
            } catch (error) {
              reject(error);
            }
          });
        });
      }
    });
    return promise2;
  }

  static resolve (parameter){
      if(parameter instanceof MyPromise){
          return parameter;
      }
      return new MyPromise((resolve) => {
          resole(parameter)
      });
  }
  static reject (reason){
      return new MyPromise((resolve,reject)=>{
          reject(reason);
      })
  }
}

function resolvePromise(promise2, x, resolve, reject) {
  if (promise2 === x) {
    return reject(
      new TypeError("Chaining cycle detected for promise #<Promise>")
    );
  }
  if (x instanceof MyPromise) {
    x.then(resolve, reject);
  } else {
    resolve(x);
  }
}

module.exports = MyPromise;

//test
// const promise = new MyPromise((resolve,reject)=>{
//     // setTimeout(() => {
//         resolve('success')
//     // },2000)
// })

// promise.then(value =>{
//     console.log('resole',value)
// }, reason=>{
//     console.log('reject',reason)
// })

// promise.then(value =>{
//     console.log(1)
//     console.log('resolve', value)
// })
// promise.then(value =>{
//     console.log(2)
//     console.log('resolve', value)
// })
// promise.then(value =>{
//     console.log(3)
//     console.log('resolve', value)
// })

// function other(){
//     return new MyPromise((resolve,reject)=>{
//         resolve('other')
//     })
// }

// promise.then(value =>{
//     console.log(1)
//     console.log('resolve', value)
//     return other()
// }).then(value =>{
//     console.log(2)
//     console.log('resolve', value)
// })

// const promise = new MyPromise((resolve, reject) => {
//   // resolve('success')
//   throw new Error("error executor");
// });

// 这个时候将promise定义一个p1，然后返回的时候返回p1这个promise
// const p1 = promise.then(value => {
//    console.log(1)
//    console.log('resolve', value)
//    return p1
// })

// promise
//   .then(
//     (value) => {
//       console.log(2);
//       console.log("resolve", value);
//       throw new Error("error");
//     },
//     (reason) => {
//       console.log(3);
//       console.log(reason.message);
//     }
//   )
//   .then((value) => {
//     console.log(4);
//     console.log(reason.message);
//   });

// const promise = new MyPromise((resolve, reject) => {
//     resolve('success')
//     // throw new Error('执行器错误')
//  })

// // 第一个then方法中的错误要在第二个then方法中捕获到
// promise.then(value => {
//   console.log(1)
//   console.log('resolve', value)
//   throw new Error('then error')
// }, reason => {
//   console.log(2)
//   console.log(reason.message)
// }).then(value => {
//   console.log(3)
//   console.log(value);
// }, reason => {
//   console.log(4)
//   console.log(reason.message)
// })

const promise = new MyPromise((resolve, reject) => {
  resolve("succ");
});

promise
  .then()
  .then()
  .then((value) => console.log(value));
