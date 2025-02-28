const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let highest = 0;
let current = 0;

data.forEach((value, _index, _array) => {
    if(!value && current > highest) {
        highest = current;
        current = 0;
    } else if(!value) {
        current = 0;
    } else {
        current += parseInt(value);
    }
});

console.log(highest);
console.log(data);