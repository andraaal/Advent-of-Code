const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

const field = [
    'GBDCPR',
    'GVH',
    'MPJDQSN',
    'MNCDGLSP',
    'SLFPCNBJ',
    'STGVZDBQ',
    'QTFHMZB',
    'FBDMC',
    'GQCF'
];

let state = field.map(value => value.split('').reverse());

let result = '';
let exit = false;

data.forEach(value => {
    if (exit) return;
    let commandContents = value.split(' ');
    const origin = parseInt(commandContents[3]);
    const times = parseInt(commandContents[1]);
    const target = parseInt(commandContents[5]);
    for (let i = 0; i < times; i++) {
        const buffer = state[origin - 1][state[origin - 1].length - 1];
        if(!buffer) {
            console.log(value);
            console.log(state);
            exit = true;
            return;
        }
        state[origin - 1].pop();
        state[target - 1].push(buffer);
    }
});
state.forEach(value => result += value[value.length - 1]);

console.log(result);