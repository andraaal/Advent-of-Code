const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let highest = new Set();
let current = 0;
let result = 0;

data.forEach((value, _index, _array) => {
    if(!value) {
        add(current);
        current = 0;
    } else {
        current += parseInt(value);
    }
});

highest.forEach(value => result += value);

console.log(result);


function add(value) {
    highest.add(parseInt(value));
    if(highest.size > 3) {
        highest.delete(Math.min(... highest));
    }
}