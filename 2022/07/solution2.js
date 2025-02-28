const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let sizes = [];
const resultSizes = [];
const totalSize = 70000000 - 42558312;

for (command of data) {
    const parsed = command.split(' ');
    if (parseInt(parsed[0]) > 0) sizes[sizes.length - 1] += parseInt(parsed[0]);

    if (parsed[2] === '..') {
        sizes[sizes.length - 2] += sizes[sizes.length - 1];
        resultSizes.push(sizes[sizes.length - 1])
        sizes.pop();
    } else if (parsed[1] === 'cd' && parsed[0] === '$') {
        sizes.push(0);
    }
}

console.log(resultSizes.sort((a, b) => a - b).find(value => value + totalSize >= 30000000));