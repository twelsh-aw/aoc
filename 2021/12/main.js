const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n")
let edges = parseInput(input);

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let initialPaths = [] // array of string
    for (let edge of edges) {
        let i = edge.indexOf('start')
        if (i > -1) {
            let path = [edge[i], edge[invertBit(i)]]
            initialPaths.push(path)
        }
    }

    let completedPaths = []
    let prevPaths = [...initialPaths]
    while (true) {
        let nextPaths = traversePath(prevPaths, completedPaths)
        if (nextPaths.length === 0) {
            break
        } else {
            prevPaths = [...nextPaths]
        }
    }

    return completedPaths.length
}

function puzzle2() {
    let initialPaths = [] // array of string
    for (let edge of edges) {
        let i = edge.indexOf('start')
        if (i > -1) {
            let path = [edge[i], edge[invertBit(i)]]
            initialPaths.push(path)
        }
    }

    let completedPaths = []
    let prevPaths = [...initialPaths]
    while (true) {
        let nextPaths = traversePath2(prevPaths, completedPaths)
        if (nextPaths.length === 0) {
            break
        } else {
            prevPaths = [...nextPaths]
        }
    }

    return completedPaths.length
}

function parseInput() {
    let edges = []
    for (let row of input) {
        let split = row.trim().split('-')
        let sorted = [split[0], split[1]]
        edges.push(sorted)
    }

    return edges;
}

function invertBit(i) {
    if (i === 0) {
        return 1
    } else if (i === 1) {
        return 0
    } else {
        throw 'bad index'
    }
}

function traversePath(pathsSoFar, completedPaths) {
    let nextPaths = []
    for (let path of pathsSoFar) {
        let cur = path[path.length - 1]
        for (let edge of edges) {
            let i = edge.indexOf(cur)
            if (i > -1) {
                let other = edge[invertBit(i)]
                if (other === other.toLowerCase() && path.some(v => v === other)) {
                    continue
                }

                let p = [...path, other]
                if (other === 'end') {
                    completedPaths.push(p)
                } else {
                    nextPaths.push(p)
                }
            }
        }
    }

    return nextPaths
}

function traversePath2(pathsSoFar, completedPaths) {
    let nextPaths = []
    for (let path of pathsSoFar) {
        let cur = path[path.length - 1]
        for (let edge of edges) {
            let i = edge.indexOf(cur)
            if (i > -1) {
                let other = edge[invertBit(i)]
                if (other === other.toLowerCase() && !checkPathForNewCave(path, other)) {
                    continue
                }

                let p = [...path, other]
                if (other === 'end') {
                    completedPaths.push(p)
                } else {
                    nextPaths.push(p)
                }
            }
        }
    }

    return nextPaths
}

function checkPathForNewCave(path, newCave) {
    if (newCave === 'start') {
        return false
    }

    let lc = path.filter(p => p === p.toLowerCase())
    let counts = {}
    for (let c of [...lc, newCave]) {
        if (counts[c] !== undefined) {
            counts[c]++
        } else {
            counts[c] = 1
        }
    }

    return Object.values(counts).filter(v => v === 2).length <= 1 && Object.values(counts).filter(v => v > 2).length === 0
}
