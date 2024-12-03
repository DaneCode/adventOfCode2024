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

// Step 3: Sort both arrays
array1.sort((a, b) => a - b);
array2.sort((a, b) => a - b);

// Step 4: Calculate the absolute difference between corresponding elements of the two arrays
let array3 = array1.map((num, index) => Math.abs(num - array2[index]));

// Step 5: Sum all the values in the third array
let totalDifference = array3.reduce((acc, curr) => acc + curr, 0);

console.log("Total difference:", totalDifference);