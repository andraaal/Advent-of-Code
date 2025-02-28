const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

result = 0;

data.forEach(line => {
    console.log(line);
    console.log(line);
    let digits = line.match(/[1-9]/g);
    console.log(digits);
    let first = digits.at(0);
    let last = digits.at(-1);
    let outputNumber = parseInt(first + last);
    console.log(outputNumber);
    result += outputNumber;
});

console.log(result);