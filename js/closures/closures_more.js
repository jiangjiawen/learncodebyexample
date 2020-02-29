//https://stackoverflow.com/questions/111102/how-do-javascript-closures-work
function say667() {
  // Local variable that ends up within closure
  var num = 42;
  var say = function() {
    console.log(num);
  };
  num++;
  return say;
}
var sayNumber = say667();
sayNumber(); // logs 43

//(In JavaScript, whenever you declare a function inside another function,
//the inside function(s) is/are recreated again each time the outside function is called.
var gLogNumber, gIncreaseNumber, gSetNumber;
function setupSomeGlobals() {
  // Local variable that ends up within closure
  var num = 42;
  // Store some references to functions as global variables
  gLogNumber = function() {
    console.log(num);
  };
  gIncreaseNumber = function() {
    num++;
  };
  gSetNumber = function(x) {
    num = x;
  };
}

setupSomeGlobals();
gIncreaseNumber();
gLogNumber(); // 43
gSetNumber(5);
gLogNumber(); // 5

var oldLog = gLogNumber;

setupSomeGlobals();
gLogNumber(); // 42

oldLog(); // 5

//variable hosting
function sayAlice() {
  var say = function() {
    console.log(alice);
  };
  // Local variable that ends up within closure
  var alice = "Hello Alice";
  return say;
}
sayAlice()(); // logs "Hello Alice"

function buildList(list) {
  var result = [];
  for (var i = 0; i < list.length; i++) {
    let item = "item" + i;
    // var item = "item" + i;
    result.push(function() {
      console.log(item + " " + list[i]);
    });
  }
  return result;
}

function testList() {
  var fnlist = buildList([1, 2, 3]);
  // Using j only to help prevent confusion -- could use i.
  for (var j = 0; j < fnlist.length; j++) {
    fnlist[j]();
  }
}

testList(); //logs "item2 undefined" 3 times


function newClosure(someNum, someRef) {
    // Local variables that end up within closure
    var num = someNum;
    var anArray = [1,2,3];
    var ref = someRef;
    return function(x) {
        num += x;
        anArray.push(num);
        console.log('num: ' + num +
            '; anArray: ' + anArray.toString() +
            '; ref.someVar: ' + ref.someVar + ';');
      }
}
obj = {someVar: 4};
fn1 = newClosure(4, obj);
fn2 = newClosure(5, obj); // attention here: new closure assigned to a new variable!
fn1(1); // num: 5; anArray: 1,2,3,5; ref.someVar: 4;
fn2(1); // num: 6; anArray: 1,2,3,6; ref.someVar: 4;
obj.someVar++;
fn1(2); // num: 7; anArray: 1,2,3,5,7; ref.someVar: 5;
fn2(2); // num: 8; anArray: 1,2,3,6,8; ref.someVar: 5;
