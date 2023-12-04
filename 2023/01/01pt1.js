const fs = require('fs');
const path = require('path');
var sum = 0;

// current directory
const currentDirectory = __dirname;

// txt path
const filePath = path.join(currentDirectory, '01Input.txt');

// read file
fs.readFile(filePath, 'utf8', (err, data) => {
  if (err) {
    console.error('Error reading the file:', err);
    return;
  }

  // split into lines
  const lines = data.split('\n');

  // process each line
  lines.forEach((line) => {
    // extract first and last digits
    const { firstDigit, lastDigit } = extractNumbers(line);
    sum += combineIntegers(firstDigit, lastDigit);
  });
  console.log(`Sum = ${sum}`);
});

function combineIntegers(firstDigit, lastDigit) {
  // convert integers to strings and combine
  const combinedString = `${firstDigit}${lastDigit}`;
  // convert string to integer
  return parseInt(combinedString, 10);
}

function extractNumbers(line) {
    // match all single digits in the string
    const digitsArray = line.match(/\d/g);

    // check if digits in string
    if (digitsArray && digitsArray.length > 0) {
      // extract first and last digits
      const firstDigit = digitsArray[0];
      const lastDigit = digitsArray[digitsArray.length - 1];
      return { firstDigit, lastDigit };
    } 
}