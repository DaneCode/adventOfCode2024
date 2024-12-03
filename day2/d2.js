const fs = require('fs');

// Function to check if a report is safe
function isSafe(report) {
    let increasing = true;
    let decreasing = true;

    for (let i = 0; i < report.length - 1; i++) {
        let diff = report[i + 1] - report[i];
        if (diff < 1 || diff > 3) {
            increasing = false;
        }
        if (diff > -1 || diff < -3) {
            decreasing = false;
        }
    }

    return increasing || decreasing;
}

// Step 1: Read the data from d2input.txt
const data = fs.readFileSync('d2input.txt', 'utf8');
const lines = data.trim().split('\n');

// Step 2: Split each line into a list of numbers
let reports = lines.map(line => line.trim().split(/\s+/).map(Number));

// Step 3-5: Check each report and count the number of safe reports
let safeCount = reports.reduce((count, report) => count + (isSafe(report) ? 1 : 0), 0);

console.log("Number of safe reports:", safeCount);