const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n")
let initial = input[0]
let transforms = input.slice(2).map(row => row.split(' -> '))

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1(n) {
    if (!n) {
        n = 10
    }

    let pairs = getPairs(initial)
    let pairCounts = getPairCounts(pairs)
    for (let i = 1; i <= n; i++) {
        let newPairCounts = {}
        for (let pair of Object.keys(pairCounts)) {
            let newPairs = getNewPairs(pair)
            for (let p2 of newPairs) {
                if (!newPairCounts[p2]) {
                    newPairCounts[p2] = 0
                }

                newPairCounts[p2] += pairCounts[pair]
            }
        }

        pairCounts = {...newPairCounts}
    }

    // double counted counts
    let charCounts = {}
    for (let pair of Object.keys(pairCounts)) {
        for (let char of pair.split('')) {
            if (!charCounts[char]) {
                charCounts[char] = 0
            }

            charCounts[char] += pairCounts[pair]
        }
    }

    let maxCount = Math.max(...Object.values(charCounts))
    let minCount = Math.min(...Object.values(charCounts))
    // remove double counting
    return ((maxCount + 1) - minCount) / 2
}

function puzzle2() {
    return puzzle1(40)
}

function getPairs(str) {
    let pairs = []
    let strArr = str.split('')
    for (let i = 0; i < strArr.length - 1; i++) {
        let pair = strArr.slice(i, i + 2).join('')
        pairs.push(pair)
    }

    return pairs
}

function getPairCounts(pairs) {
    let counts = {}
    for (let pair of pairs) {
        if (!counts[pair]) {
            counts[pair] = 0
        }

        counts[pair]++
    }

    return counts
}

function getNewPairs(pair) {
    let transform = transforms.find(t => t[0] === pair)
    if (transform) {
        let split = transform[0].split('')
        return [split[0] + transform[1], transform[1] + split[1]]
    }

    return [pair]
}
