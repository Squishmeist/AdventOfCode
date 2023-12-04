const fs = require('fs');
const path = require('path');

function hasWon(winningNumbers, givenCards) {
    let currentPoints = 0;
    let isFirstMatch = true;

    for (const number of givenCards) {
        if (winningNumbers.includes(number)) {
            if (isFirstMatch) {
                currentPoints = 1;
                isFirstMatch = false;
            } else {
                currentPoints *= 2;
            }
        }
    }

    return currentPoints;
}

function main() {
    const inputFilePath = path.join(__dirname, '04input.txt');
    let sum = 0;

    try {
        const data = fs.readFileSync(inputFilePath, 'utf8');
        const cardRecords = data.split('\n');

        for (const cardRecord of cardRecords) {
            try {
                const match = cardRecord.match(/Card\s+(\d+):/);
                if (!match) {
                    throw new Error("Invalid card record format");
                }

                const cardNumber = parseInt(match[1], 10);

                // Extracting the winning numbers and given cards sections
                const sections = cardRecord.split(':');
                if (sections.length !== 2) {
                    throw new Error("Invalid card record format");
                }

                const winningNumbersStr = sections[1].split('|')[0];
                const givenCardsStr = sections[1].split('|')[1];

                const winningNumbers = winningNumbersStr.trim().split(/\s+/).filter(num => num !== '').map(num => parseInt(num, 10));
                const givenCards = givenCardsStr.trim().split(/\s+/).filter(num => num !== '').map(num => parseInt(num, 10));

                const points = hasWon(winningNumbers, givenCards);
                console.log(`Card ${cardNumber} Points: ${points}`);

                sum += points;

            } catch (parseError) {
                console.error("Error parsing card record:", cardRecord);
                console.error(parseError.stack);
            }
        }

        console.log("Scratchcard points:", sum);
    } catch (error) {
        console.error("Error reading the file:", error.message);
        console.error(error.stack);
    }
}

main();