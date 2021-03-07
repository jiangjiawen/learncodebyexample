//https://www.youtube.com/watch?v=oBt53YbR9Kk&t=50s&ab_channel=freeCodeCamp.org
const fib = (n) => {
    if(n<=2) return 1;
    return fib(n-1) + fib(n-2);
};

console.log(fib(6));
console.log(fib(7));
console.log(fib(8));
// console.log(fib(50));

//memorization
// js object, keys will be arg to fn, value will the be the return value

const fib2 = (n, memo = {}) =>{
    if (n in memo) return memo[n];
    if (n<=2) return 1;
    memo[n] = fib2(n-1, memo) + fib2(n-2, memo);
    return memo[n];
};
console.log(fib2(6));
console.log(fib2(7));
console.log(fib2(8));
console.log(fib2(50));