const fs = require('fs');

function evaluateExpression(numbers, operators) {
    let result = numbers[0];
    for (let i = 0; i < operators.length; i++) {
        if (operators[i] === '+') {
            result += numbers[i + 1];
        } else if (operators[i] === '*') {
            result *= numbers[i + 1];
        }
    }
    return result;
}

function canBeTrue(testValue, numbers) {
    if (numbers.length === 1) {
        return numbers[0] === testValue;
    }

    const operators = ['+', '*'];
    const combinations = getCombinations(operators, numbers.length - 1);

    for (const ops of combinations) {
        if (evaluateExpression(numbers, ops) === testValue) {
            return true;
        }
    }
    return false;
}

function getCombinations(operators, length) {
    if (length === 1) {
        return operators.map(op => [op]);
    }

    const combinations = [];
    const smallerCombinations = getCombinations(operators, length - 1);

    for (const op of operators) {
        for (const smallerCombination of smallerCombinations) {
            combinations.push([op, ...smallerCombination]);
        }
    }

    return combinations;
}

function main() {
    let totalCalibrationResult = 0;
    const data = fs.readFileSync('d7Input.txt', 'utf8');
    const lines = data.split('\n');

    for (const line of lines) {
        if (line.trim() === '') continue;

        const [testValueStr, numbersStr] = line.split(':');
        const testValue = parseInt(testValueStr.trim());
        const numbers = numbersStr.trim().split(' ').map(Number);

        if (canBeTrue(testValue, numbers)) {
            totalCalibrationResult += testValue;
        }
    }

    console.log(totalCalibrationResult);
}

main();