const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n").map(row => row.split('').map(Number))

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let costs = makeArray(input[0].length, input.length)
    let end = [input.length - 1, input[0].length - 1]
    costs[end[0]][end[1]] = input[end[0]][end[[1]]]
    for (let i = input.length - 1; i >= 0; i--) {
        for (let j = input[0].length - 1; j >= 0; j--) {
            if (costs[i][j] > -1) {
                continue
            }

            let possible = getPossibleMoves([i, j])
            let costsForPos = []
            for (let pos of possible) {
                if (costs[pos[0]][pos[1]] <= 0) {
                    console.error(i, j, pos, costsForPos)
                    throw 'unknown'
                }

                costsForPos.push(input[i][j] + costs[pos[0]][pos[1]])
            }

            costs[i][j] = Math.min(...costsForPos)
        }
    }

    console.log(costs[0][0] - input[0][0])
}

function puzzle2() {

}

function getPossibleMoves(pos, considerBackwards) {
    let ret = []
    if (input[pos[0]][pos[1]+1]) {
        ret.push([pos[0],pos[1]+1])
    }

    if (input[pos[0] + 1]) {
        ret.push([pos[0] + 1,pos[1]])
    }

    return ret
}

function makeArray(d1, d2) {
    var arr = [];
    for(let i = 0; i < d2; i++) {
        arr.push(new Array(d1).fill(-1));
    }
    return arr;
}
