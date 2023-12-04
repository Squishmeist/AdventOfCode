const fs = require('fs');
const path = require('path');

function minimumSet(gameInfo) {
    const currentCounts = { red: 0, green: 0, blue: 0 };

    for (const subset of gameInfo) {
        for (const cubeInfo of subset) {
            const [count, colour] = cubeInfo.split(' ');
            const parsedCount = parseInt(count);

            // Update currentCounts if the parsed count is greater than the current count
            if (parsedCount > currentCounts[colour]) {
                currentCounts[colour] = parsedCount;
            }
        }
    }

    // Extract the counts for each color
    const redCount = currentCounts.red;
    const greenCount = currentCounts.green;
    const blueCount = currentCounts.blue;

    // Multiply the counts together
    const result = redCount * greenCount * blueCount;

    // Return the result
    return result;
}


function main() {
    const targetCounts = { red: 12, green: 13, blue: 14 };
    let sum = 0;

    const inputFilePath = path.join(__dirname, '02input.txt');

    try {
        const data = fs.readFileSync(inputFilePath, 'utf8');
        const gameRecords = data.split('\n');

        for (const gameRecord of gameRecords) {
            try {
                const match = gameRecord.match(/Game (\d+):/);
                if (!match) {
                    throw new Error("Invalid game record format");
                }

                const gameInfo = gameRecord.split(":")[1].split(';').map(subset => subset.trim().split(", "));

                const result = minimumSet(gameInfo);
                sum += result;

            } catch (parseError) {
                console.error("Error parsing game record:", gameRecord);
                console.error(parseError.stack);
            }
        }

        console.log("Sum of set powers:", sum);
    } catch (error) {
        console.error("Error reading the file:", error.message);
        console.error(error.stack);
    }
}

main();
