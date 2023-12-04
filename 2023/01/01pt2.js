const fs = require('fs');
const path = require('path');

const digitsDict = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
};

function scanLine(line, accumulate) {
    let acc = '';

    for (const x of line) {
        if (x.match(/\d/)) {
            return x;
        } else {
            acc = accumulate(acc, x);
            for (const [k, v] of Object.entries(digitsDict)) {
                if (acc.includes(k)) {
                    return v;
                }
            }
        }
    }

    console.error('Not found in', acc);
    return '0';
}

function main() {
    const inputFilePath = path.join(__dirname, '01input.txt');
    if (!fs.existsSync(inputFilePath)) {
        console.error('No such file:', inputFilePath);
        return;
    }

    let total = 0;

    const data = fs.readFileSync(inputFilePath, 'utf8').split('\n');
    for (const line of data) {
        const num = scanLine(line, (x, y) => `${x}${y}`) + scanLine([...line].reverse(), (x, y) => `${y}${x}`);
        total += parseInt(num, 10);
    }

    console.log(total);
}

if (require.main === module) {
    main();
}
