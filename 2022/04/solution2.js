const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let sum = 0;

data.forEach(value => {
    const pairs = value.split(',');
    const firstElf = pairs[0].split('-');
    const secondElf = pairs[1].split('-');
    const doesContainFirst = parseInt(firstElf[0]) <= parseInt(secondElf[1]) && parseInt(firstElf[1]) >= parseInt(secondElf[0]);
    const doesContainSecond = parseInt(firstElf[1]) >= parseInt(secondElf[0]) && parseInt(firstElf[0]) <= parseInt(secondElf[1]);

    if (doesContainFirst || doesContainSecond) sum++;
});

console.log(sum);