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

// Function to check if a report can be made safe by removing one level
function canBeSafeByRemovingOne(report) {
    for (let i = 0; i < report.length; i++) {
        let newReport = report.slice(0, i).concat(report.slice(i + 1));
        if (isSafe(newReport)) {
            return true;
        }
    }
    return false;
}

// Step 1: Read the data from d2input.txt
const data = fs.readFileSync('d2input.txt', 'utf8');
const lines = data.trim().split('\n');

// Step 2: Split each line into a list of numbers
let reports = lines.map(line => line.trim().split(/\s+/).map(Number));

// Step 3-6: Check each report and count the number of safe reports
let safeCount = 0;
for (let report of reports) {
    if (isSafe(report) || canBeSafeByRemovingOne(report)) {
        safeCount++;
    }
}

console.log("Number of safe reports:", safeCount);