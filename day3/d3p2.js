const fs = require('fs');

// Function to extract and multiply numbers
function extractAndMultiply(filename) {
    const results = [];
    // Updated pattern to strictly match valid mul(X,Y) instructions
    const mulPattern = /mul\((\d{1,3}),(\d{1,3})\)/g;

    const data = fs.readFileSync(filename, 'utf8');
    let match;
    let mulEnabled = true; // Initially, mul instructions are enabled

    for (let i = 0; i < data.length; i++) {
        // Check for do() and don't() instructions
        if (data.slice(i).startsWith('do()')) {
            mulEnabled = true;
            i += 3; // Move index past the do() instruction
        } else if (data.slice(i).startsWith("don't()")) {
            mulEnabled = false;
            i += 6; // Move index past the don't() instruction
        }

        // Check for mul(X,Y) instructions
        if (mulEnabled) {
            mulPattern.lastIndex = i;
            match = mulPattern.exec(data);
            if (match && match.index === i) {
                const n1 = parseInt(match[1], 10);
                const n2 = parseInt(match[2], 10);
                results.push(n1 * n2);
                i += match[0].length - 1; // Move index past the mul(X,Y) instruction
            }
        }
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