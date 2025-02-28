const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let sizes = [];
let sum = 0;

for(command of data) {
    const parsed = command.split(' ');
    if(parseInt(parsed[0]) > 0) sizes[sizes.length - 1] += parseInt(parsed[0]);

    if(parsed[2] === '..') {
        console.log(sizes[sizes.length - 1], sum);
        sizes[sizes.length - 2] += sizes[sizes.length - 1];
        if(sizes[sizes.length - 1] <= 100000) sum += sizes[sizes.length - 1];
        sizes.pop();
    } else if(parsed[1] === 'cd' && parsed[0] === '$') {
        sizes.push(0);
    }
}

console.log(sum);