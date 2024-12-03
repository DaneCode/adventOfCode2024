const fs = require('fs');

// Function to extract and multiply numbers
function extractAndMultiply(filename) {
    const results = [];
    // Updated pattern to strictly match valid mul(X,Y) instructions
    const pattern = /mul\((\d{1,3}),(\d{1,3})\)/g;

    const data = fs.readFileSync(filename, 'utf8');
    let match;
    while ((match = pattern.exec(data)) !== null) {
        const n1 = parseInt(match[1], 10);
        const n2 = parseInt(match[2], 10);
        results.push(n1 * n2);
    }

    return results;
}

// File path
const filePath = 'd3.txt';

// Get the results
const multiplicationResults = extractAndMultiply(filePath);

// Calculate the sum of the results
const totalSum = multiplicationResults.reduce((sum, value) => sum + value, 0);

// Print the sum of the results
console.log('Sum of multiplication results:');
console.log(totalSum);