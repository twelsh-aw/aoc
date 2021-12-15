const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let input = fs.readFileSync(inputPath).toString().split("\n").map(row => row.split('').map(Number))
    return puzzle(input)
}

function puzzle2() {
    let inputRaw = fs.readFileSync(inputPath).toString().split("\n").map(row => row.split('').map(Number))
    let expanded = expandInput(inputRaw)
    return puzzle(expanded)
}

function puzzle(input) {
    let costs = makeArray(input[0].length, input.length, Infinity)
    let unvisited = makeArray(input[0].length, input.length, true)
    costs[0][0] = 0
    let current = [0, 0]
    let prev

    while (current !== undefined) {
        step(costs, input, unvisited, current)
        prev = [...current]
        current = getNext(costs, unvisited)
    }

    return costs[input.length - 1][input[0].length - 1]
}

function getNext(costs, unvisited) {
    let minCost = Infinity
    let next = undefined
    for (let i = 0; i < costs.length; i++) {
        for (let j = 0; j < costs[0].length; j++) {
            if (unvisited[i][j] && costs[i][j] < minCost) {
                next = [i, j]
                minCost = costs[i][j]
            }
        }
    }

    return next
}

function step(costs, input, unvisited, current) {
    let toVisit = getPossibleMoves(current, input, unvisited)
    for (let p of toVisit) {
        if (input[p[0]][p[1]] <= 0 || input[p[0]][p[1]] > 9) {
            throw p
        }

        let cost = input[p[0]][p[1]] + costs[current[0]][current[1]]
        if (costs[p[0]][p[1]] > cost) {
            costs[p[0]][p[1]] = cost
        }
    }

    unvisited[current[0]][current[1]] = false
}

function getPossibleMoves(pos, input, unvisited) {
    let ret = []
    if (input[pos[0]][pos[1]+1]) {
        ret.push([pos[0],pos[1]+1])
    }

    if (input[pos[0] + 1]) {
        ret.push([pos[0] + 1,pos[1]])
    }

    if (input[pos[0]][pos[1]-1]) {
        ret.push([pos[0],pos[1]-1])
    }

    if (input[pos[0]-1]) {
        ret.push([pos[0]-1,pos[1]])
    }

    return ret.filter(r => unvisited[r[0]][r[1]])
}

function makeArray(d1, d2, fill) {
    var arr = [];
    for(let i = 0; i < d2; i++) {
        arr.push(new Array(d1).fill(fill));
    }
    return arr;
}

function expandInput(inputRaw) {
    let n = inputRaw.length
    let expanded = makeArray(n * 5, n * 5)
    for (let k = 0; k < 5; k++) {
        for (let l = 0; l < 5; l++) {
            for (let i = 0; i < inputRaw.length; i++) {
                for (let j = 0; j < inputRaw[0].length; j++) {
                    let initial = inputRaw[i][j]
                    let nextVal = 1 + ((initial + k + l - 1) % 9)
                    expanded[i+(k*n)][j+(l*n)] = nextVal
                }
            }
        }
    }

    return expanded
}
