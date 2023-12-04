const fs = require('fs');
const path = require('path');

function isPossible(gameInfo, targetCounts) {

    for (const subset of gameInfo) {
        for (const cubeInfo of subset) {
            const [count, colour] = cubeInfo.split(' ');
            if(targetCounts[colour] < count){
                return false;
            }
        }
    }

    return true;
}

function main() {
    const targetCounts = { red: 12, green: 13, blue: 14 };
    const possibleGames = [];

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

                const gameId = match[1];
                const gameInfo = gameRecord.split(":")[1].split(';').map(subset => subset.trim().split(", "));

                if (isPossible(gameInfo, targetCounts)) {
                    possibleGames.push(parseInt(gameId));
                }

            } catch (parseError) {
                console.error("Error parsing game record:", gameRecord);
                console.error(parseError.stack);
            }
        }

        const totalSum = possibleGames.reduce((sum, gameId) => sum + gameId, 0);
        console.log("Possible Games:", possibleGames);
        console.log("Sum of IDs:", totalSum);
    } catch (error) {
        console.error("Error reading the file:", error.message);
        console.error(error.stack);
    }
}

main();
