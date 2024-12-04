const fs = require('fs');

function readGrid(filename) {
    const data = fs.readFileSync(filename, 'utf8');
    return data.split('\n').map(line => line.trim());
}

function isXmasPattern(grid, row, col) {
    // Check if the X-MAS pattern fits within the grid boundaries
    if (row + 2 >= grid.length || col + 2 >= grid[0].length) {
        return false;
    }

    // Check the X-MAS pattern
    if (grid[row][col] === 'M' && grid[row][col + 2] === 'S' &&
        grid[row + 1][col + 1] === 'A' &&
        grid[row + 2][col] === 'M' && grid[row + 2][col + 2] === 'S') {
        return true;
    }

    // Check the reversed X-MAS pattern
    if (grid[row][col] === 'S' && grid[row][col + 2] === 'M' &&
        grid[row + 1][col + 1] === 'A' &&
        grid[row + 2][col] === 'S' && grid[row + 2][col + 2] === 'M') {
        return true;
    }
    // Check Downward X-MAS pattern
    if (grid[row][col] === 'M' && grid[row][col + 2] === 'M' &&
        grid[row + 1][col + 1] === 'A' &&
        grid[row + 2][col] === 'S' && grid[row + 2][col + 2] === 'S') {
        return true;
    }
    // Check Upward X-MAS pattern
    if (grid[row][col] === 'S' && grid[row][col + 2] === 'S' &&
        grid[row + 1][col + 1] === 'A' &&
        grid[row + 2][col] === 'M' && grid[row + 2][col + 2] === 'M') {
        return true;
    }

    return false;
}

function countXmasPatterns(grid) {
    let count = 0;
    for (let row = 0; row < grid.length; row++) {
        for (let col = 0; col < grid[0].length; col++) {
            if (isXmasPattern(grid, row, col)) {
                count += 1;
            }
        }
    }
    return count;
}

const grid = readGrid('day4Input.txt');
const xmasCount = countXmasPatterns(grid);
console.log(`Number of X-MAS patterns: ${xmasCount}`);