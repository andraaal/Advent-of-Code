const fs = require('fs');
let input = fs.readFileSync('AdventOfCode.txt', 'utf-8');

let numbers = input.split("\r\n");
numbers = numbers.map(i => parseInt(i));
for (let i = 0; i < numbers.length; i++) {
    for (let j = 0; j < numbers.length; j++) {
        for (let k = 0; k < numbers.length; k++) {
            if(numbers[i] + numbers[j] + numbers[k] === 2020) {
                answer = numbers[i] * numbers[j] * numbers[k];
            }
        }
    }
}
console.log(answer);
