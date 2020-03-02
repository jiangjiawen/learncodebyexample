//https://www.jianshu.com/p/9adb185a8188

// for (var i = 0; i < 5; i++) {
//   setTimeout(function() {
//     console.log(i);
//   }, 1000 * i);
// }

// for (var i = 0; i < 5; i++) {
//   (function(j) {
//     // j = i
//     setTimeout(function() {
//       console.log(new Date(), j);
//     }, 1000);
//   })(i);
// }

// for (var i = 1; i <= 5; i++) {
//   (function() {
//     var j = i;
//     setTimeout(function timer() {
//       console.log(j);
//     }, j * 1000);
//   })();
// }

// //error, IIFE is undefined
// // for (var i = 0; i < 5; i++) {
// //   setTimeout(
// //     (function(i) {
// //       console.log(i);
// //     })(i),
// //     i * 1000
// //   );
// // }
// for (let i = 0; i < 5; i++) {
//   setTimeout(function() {
//     console.log(i);
//   }, 1000 * i);
// }

const tasks = [];
for (let i = 0; i < 5; i++) {
  // 这里 i 的声明不能改成 let，如果要改该怎么做？
  (j => {
    tasks.push(
      new Promise(resolve => {
        setTimeout(() => {
          console.log(new Date(), j);
          resolve(); // 这里一定要 resolve，否则代码不会按预期 work
        }, 1000 * j); // 定时器的超时时间逐步增加
      })
    );
  })(i);
}

// const sleep = timeountMS =>
//   new Promise(resolve => {
//     setTimeout(resolve, timeountMS);
//   });
// (async () => {
//   // 声明即执行的 async 函数表达式
//   for (var i = 0; i < 5; i++) {
//     await sleep(1000);
//     console.log(new Date(), i);
//   }
// })();
