const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');
const rules = new Map([['A X', 3], ['B X', 1], ['C X', 2], ['A Y', 4], ['B Y', 5], ['C Y', 6], ['A Z', 8], ['B Z', 9], ['C Z', 7]]);

let score = 0;

data.forEach(value => score += rules.get(value));

console.log(score);