const fs = require('fs');

// Step 1: Read the data from d1Input.txt
const data = fs.readFileSync('d1Input.txt', 'utf8');
const lines = data.trim().split('\n');

// Step 2: Split each line into two numbers and store them in two separate arrays
let array1 = [];
let array2 = [];
lines.forEach(line => {
    const nums = line.trim().split(/\s+/).map(Number); // Use regex to handle multiple spaces
    if (nums.length !== 2) {
        console.log("Invalid line format:", line);
        return;
    }
    array1.push(nums[0]);
    array2.push(nums[1]);
});

// Step 3: Count the occurrences of each number from the first array in the second array
let counter = {};
array2.forEach(num => {
    counter[num] = (counter[num] || 0) + 1;
});

// Step 4: Multiply each number from the first array by its count in the second array and store the result in a third array
let array3 = array1.map(num => num * (counter[num] || 0));

// Step 5: Sum all the values in the third array to get the total
let total = array3.reduce((acc, curr) => acc + curr, 0);

console.log("Total:", total);