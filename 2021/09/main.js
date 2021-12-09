const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");
let parsed = parseInput(input);
let allBasinPoints = []

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let lp = getLowPoints()
    let sum = 0
    for (let p of lp) {
        sum += (parsed[p[0]][p[1]] + 1)
    }

    return sum
}

function puzzle2() {
    let lp = getLowPoints()
    let topBasins = [0, 0, 0]
    for (let p of lp) {
        topBasins = topBasins.sort((a, b) => a - b)
        let curBasin = removeDups(getBasinPoints(p))
        // console.log(p, curBasin)
        // allBasinPoints = allBasinPoints.concat(curBasin)
        // checkForDups(allBasinPoints)
        if (curBasin.length > topBasins[0]) {
            // console.log(p, curBasin)
            topBasins = [curBasin.length, topBasins[1], topBasins[2]]
        }
    }

    // checkPoints()
    return topBasins[0] * topBasins[1] * topBasins[2]
}

function checkPoints() {
    let rows = parsed.length;
    let cols = parsed[0].length
    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            let val = parsed[row][col]
            if (val !== 9) {
                if (!allBasinPoints.some(ap => ap[0] == row && ap[1] == col)) {
                    console.log('point not in basin ' + [row, col] + ' with value ' + val)
                }
            }
        }
    }
}

function checkForDups(allBasinPoints) {
    for (let p of allBasinPoints) {
        if (allBasinPoints.filter(ap => ap[0] == p[0] && ap[1] == p[1]).length > 1) {
            throw 'dup basin point ' + p.toString()
        }
    }
}

function removeDups(basinPoints) {
    let all = []
    for (let p of basinPoints) {
        if (!all.some(ap => ap[0] == p[0] && ap[1] == p[1]) && parsed[p[0]][p[1]] !== 9) {
            all.push(p)
        }
    }

    return all
}

function getBasinPoints(point) {
    let all = [point]
    let adj = getBasinAdjacentPoints(point)
    for (let a of adj) {
        all = all.concat(getBasinPoints(a))
    }

    return all
}

function getBasinAdjacentPoints(point) {
    let b = []
    let row = point[0]
    let col = point[1]
    let value = parsed[row][col]
    // left
    if (parsed[row][col - 1] >= value + 1) {
        b.push([row, col - 1])
    }

    // right
    if (parsed[row][col + 1] >= value + 1) {
        b.push([row, col + 1])
    }

    // up
    if (parsed[row+1] && parsed[row+1][col] >= value + 1) {
        b.push([row + 1, col])
    }

    // down
    if (parsed[row-1] && parsed[row-1][col] >= value + 1) {
        b.push([row - 1, col])
    }

    return b
}

function parseInput() {
    let parsed = []
    for (let i of input) {
        let row = i.trim().split('').map(Number)
        parsed.push(row)
    }

    return parsed
}

function getLowPoints() {
    let rows = parsed.length;
    let cols = parsed[0].length
    let lowPoints = []
    for (let row = 0; row < rows; row++) {
        for (let col = 0; col < cols; col++) {
            let value = parsed[row][col]
            let left, right, up, down
            try {
                left = parsed[row][col-1]
                left.toString()
            } catch{
                left = 99
                // console.log('left', row, col)
            }
            try {
                right = parsed[row][col+1]
                right.toString()
            } catch{
                right = 99
                // console.log('right', row, col)
            }
            try {
                up = parsed[row - 1][col]
            } catch{
                up = 99
                // console.log('up', row, col)
            }
            try {
                down = parsed[row + 1][col]
            } catch{
                down = 99
                // console.log('down', row, col)
            }

            if (value < Math.min(...[up, down, left, right])) {
                lowPoints.push([row, col])
            }
        }
    }

    return lowPoints
}
