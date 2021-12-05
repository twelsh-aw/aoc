const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");
let lines = parseInput(input)

let p1 = puzzle1(lines)
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1(lines) {
    let coords = makeArray(1000, 1000)
    for (let line of lines.filter(line => line.x1 === line.x2 || line.y1 === line.y2)) {
        if (line.x1 === line.x2) {
            let yStart = Math.min(line.y1, line.y2)
            let yEnd = Math.max(line.y1, line.y2)
            for (let y = yStart; y <= yEnd; y++) {
                coords[y][line.x1]++
            }
        }

        if (line.y1 === line.y2) {
            let xStart = Math.min(line.x1, line.x2)
            let xEnd = Math.max(line.x1, line.x2)
            for (let x = xStart; x <= xEnd; x++) {
                coords[line.y1][x]++
            }
        }
    }

    let num = 0
    for (let row of coords) {
        num += row.filter(r => r > 1).length
    }

    return num
}

function puzzle2() {
    let coords = makeArray(1000, 1000)
    for (let line of lines) {
        if (line.x1 === line.x2) {
            let yStart = Math.min(line.y1, line.y2)
            let yEnd = Math.max(line.y1, line.y2)
            for (let y = yStart; y <= yEnd; y++) {
                coords[y][line.x1]++
            }
        } else if (line.y1 === line.y2) {
            let xStart = Math.min(line.x1, line.x2)
            let xEnd = Math.max(line.x1, line.x2)
            for (let x = xStart; x <= xEnd; x++) {
                coords[line.y1][x]++
            }
        } else if (line.y2 > line.y1 && line.x2 > line.x1) { // up diag
            let diff = line.x2 - line.x1
            for (let i = 0; i <= diff; i++) {
                coords[line.y1+i][line.x1+i]++
            }
        } else if (line.y1 > line.y2 && line.x1 > line.x2) { // up diag case 2
            let diff = line.x1 - line.x2
            for (let i = 0; i <= diff; i++) {
                coords[line.y2+i][line.x2+i]++
            }
        } else if (line.y2 < line.y1 && line.x2 > line.x1) { // down diag
            let diff = line.x2 - line.x1
            for (let i = 0; i <= diff; i++) {
                coords[line.y1-i][line.x1+i]++
            }
        } else if (line.y2 > line.y1 && line.x2 < line.x1) { // down diag case 2
            let diff = line.x1 - line.x2
            for (let i = 0; i <= diff; i++) {
                coords[line.y2-i][line.x2+i]++
            }
        } else {
            throw line
        }
    }

    let num = 0
    for (let row of coords) {
        num += row.filter(r => r > 1).length
    }

    return num
}

function parseInput(input) {
    let lines = []
    for (let line of input) {
        let s = line.split(" -> ")
        let x1 = parseInt(s[0].split(',')[0])
        let y1 = parseInt(s[0].split(',')[1])
        let x2 = parseInt(s[1].split(',')[0])
        let y2 = parseInt(s[1].split(',')[1])
        lines.push({x1, y1, x2, y2})
    }

    return lines
}

function makeArray(d1, d2) {
    var arr = [];
    for(let i = 0; i < d2; i++) {
        arr.push(new Array(d1).fill(0));
    }
    return arr;
}
