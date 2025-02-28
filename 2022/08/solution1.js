const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n').map(value => value.split(''));
const total = data.length * data[0].length;
let sum = 0;

data.forEach((value, Yindex) => {
    value.forEach((value, Xindex) => {
        if(checkTree(Xindex, Yindex)) sum++;
    });
});

function checkTree(x, y) {
    let numberOfBlocked = 0;

    for (let i = y - 1; i >= 0; i--) {
        if (data[i][x] >= data[y][x]) {
            numberOfBlocked++;
            break;
        }
    }
    for (let i = y + 1; i < data.length; i++) {
        if (data[i][x] >= data[y][x]) {
            numberOfBlocked++;
            break;
        }
    }
    for (let i = x - 1; i >= 0; i--) {
        if (data[y][i] >= data[y][x]) {
            numberOfBlocked++;
            break;
        }
    }
    for (let i = x + 1; i < data[0].length; i++) {
        if (data[y][i] >= data[y][x]) {
            numberOfBlocked++;
            break;
        }
    }

    return numberOfBlocked === 4;
}

console.log(total - sum);