/*
_Binary Search_

This JavaScript module implements a binary search for two lists.

input: sorted list 

output: 
    - index of value if in list and the number of steps
    - -1 if the value is not in the list (this return value could vary depending on requirements) and the number of steps

*/

// create unordered lists
let myListIntegers = [3, 4, 510, 45, 60, 23, 303];
let myListNames = ["Andrew", "Jake", "Russell", "Aaron", "Darius"];

// print unordered lists
console.log(myListIntegers, myListNames);

// create ordered lists; note that in JavaScript the sort method
// only works on strings; to sort with Integers, you need to provide a comparison function
myListIntegers.sort(function (a, b) {
  return a - b;
});
myListNames.sort();

// print ordered lists
console.log(myListIntegers, myListNames);

// binary search function
function binary_search(inputList, searchItem) {
  low = 0;
  high = inputList.length - 1;
  steps = 0;
  error = -1;

  while (low <= high) {
    steps++;
    mid = Math.floor((low + high) / 2);

    guess = inputList[mid];

    if (guess === searchItem) {
      return { mid, steps };
    }

    if (guess > searchItem) {
      high = mid - 1;
    } else {
      low = mid + 1;
    }
  }
  return { error, steps };
}

console.log(binary_search(myListIntegers, 5));
console.log(binary_search(myListIntegers, 303));
console.log(binary_search(myListNames, "Terry"));
console.log(binary_search(myListNames, "Aaron"));
