const fs = require('fs');
const path = require('path');

function extractNumbers(line) {
    const numbers = [];
    let currentNumber = '';

    for (const char of line) {
        if (char === '.' && currentNumber !== '') {
            numbers.push(parseInt(currentNumber, 10));
            currentNumber = '';
        } else if (!isNaN(parseInt(char, 10))) {
            currentNumber += char;
        }
    }

    return numbers;
}

function main() {
    const inputFilePath = path.join(__dirname, '03test.txt');

    try {
        const data = fs.readFileSync(inputFilePath, 'utf8');
        const engineSchematics = data.split('\n');

        for (const engineSchematic of engineSchematics) {
            try {
                const numbers = extractNumbers(engineSchematic);
                console.log(numbers);

            } catch (parseError) {
                console.error("Error parsing record:", engineSchematic);
                console.error(parseError.stack);
            }
        }

        
    } catch (error) {
        console.error("Error reading the file:", error.message);
        console.error(error.stack);
    }
}

main();
