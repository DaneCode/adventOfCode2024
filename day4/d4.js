const fs = require('fs');

function readWordSearch(filePath) {
    const data = fs.readFileSync(filePath, 'utf8');
    return data.split('\n').map(line => line.trim());
}

function countWordOccurrences(grid, word) {
    const rows = grid.length;
    const cols = grid[0].length;
    const wordLen = word.length;
    let count = 0;

    function checkDirection(x, y, dx, dy) {
        for (let i = 0; i < wordLen; i++) {
            if (!(0 <= x + i * dx && x + i * dx < rows && 0 <= y + i * dy && y + i * dy < cols)) {
                return false;
            }
            if (grid[x + i * dx][y + i * dy] !== word[i]) {
                return false;
            }
        }
        return true;
    }

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < cols; j++) {
            if (grid[i][j] === word[0]) {
                const directions = [
                    [1, 0], [0, 1], [1, 1], [1, -1],
                    [-1, 0], [0, -1], [-1, -1], [-1, 1]
                ];
                for (const [dx, dy] of directions) {
                    if (checkDirection(i, j, dx, dy)) {
                        count += 1;
                    }
                }
            }
        }
    }

    return count;
}

if (require.main === module) {
    const grid = readWordSearch('day4Input2.txt');
    const word = 'MAS'; // Example word to search for
    const occurrences = countWordOccurrences(grid, word);
    console.log(`Number of occurrences of the word "${word}": ${occurrences}`);
}