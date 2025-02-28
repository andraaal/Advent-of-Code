const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');
const rules = new Map([['A X', 4], ['B X', 1], ['C X', 7], ['A Y', 8], ['B Y', 5], ['C Y', 2], ['A Z', 3], ['B Z', 9], ['C Z', 6]]);

let score = 0;

data.forEach(value => score += rules.get(value));

console.log(score);