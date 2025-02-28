const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let sum = 0;

data.forEach(value => {
    let same;
    const rucksack = value.split('');
    const comp1 = new Set(rucksack.slice(0, rucksack.length / 2));
    const comp2 = new Set(rucksack.slice(-rucksack.length / 2));
    for (snack of comp1) {
        if (comp2.has(snack)) { same = snack; break; }
    }
    if (same === same.toLowerCase()) {
        sum += same.charCodeAt(0) - 96;
    } else {
        sum += same.charCodeAt(0) - 38;
    }
});

console.log(sum);