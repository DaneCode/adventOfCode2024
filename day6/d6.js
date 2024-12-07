const fs = require('fs');

// Read the input file
const input = fs.readFileSync('d6Input.txt', 'utf-8').split('\n').map(line => line.split(''));

// Directions: up, right, down, left
const directions = [
  { dx: -1, dy: 0 }, // up
  { dx: 0, dy: 1 },  // right
  { dx: 1, dy: 0 },  // down
  { dx: 0, dy: -1 }  // left
];

// Find the initial position and direction of the guard
let guardPos = null;
let guardDir = 0; // Initially facing up

for (let i = 0; i < input.length; i++) {
  for (let j = 0; j < input[i].length; j++) {
    if (input[i][j] === '^') {
      guardPos = { x: i, y: j };
      input[i][j] = '.'; // Replace guard's initial position with empty space
      break;
    }
  }
  if (guardPos) break;
}

const visited = new Set();
visited.add(`${guardPos.x},${guardPos.y}`);

while (true) {
  const nextPos = {
    x: guardPos.x + directions[guardDir].dx,
    y: guardPos.y + directions[guardDir].dy
  };

  // Check if the guard has left the mapped area
  if (
    nextPos.x < 0 || nextPos.x >= input.length ||
    nextPos.y < 0 || nextPos.y >= input[0].length
  ) {
    break;
  }

  if (input[nextPos.x][nextPos.y] === '#') {
    // Turn right 90 degrees
    guardDir = (guardDir + 1) % 4;
  } else {
    // Move forward
    guardPos = nextPos;
    visited.add(`${guardPos.x},${guardPos.y}`);
  }
}

console.log(`Distinct positions visited: ${visited.size}`);