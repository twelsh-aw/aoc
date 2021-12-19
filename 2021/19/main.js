const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')
let parsed = parseInput(input)
let transforms = getTransforms()

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let coords = []
    let total = 0
    for (let i = 0; i < parsed.length; i++) {
        for (let coord of parsed[i]) {
            let adjusted = getCoordRelativeToScannerZero(coord, i, transforms)
            if (!coords.some(c => c[0] === adjusted[0] && c[1] === adjusted[1] && c[2] === adjusted[2])) {
                coords.push(adjusted)
            }
            total++
        }
    }

    return coords.length
}

function puzzle2() {
    let scannerCoords = []
    for (let i = 0; i < parsed.length; i++) {
        let adjusted = getCoordRelativeToScannerZero([0, 0, 0], i, transforms)
        scannerCoords.push(adjusted)
    }

    let maxDistance = 0
    let distances = getBeaconDistances(scannerCoords)
    for (let i = 0; i < distances.length; i++) {
        for (let j = 0; j < distances[i].length; j++) {
            let d = Math.abs(distances[i][j][0]) + Math.abs(distances[i][j][1]) + Math.abs(distances[i][j][2])
            if (d > maxDistance) {
                maxDistance = d
            }
        }
    }

    return maxDistance
}

function getTransforms() {
    let transforms = []
    let trackedIndexes = [0]
    while (trackedIndexes.length !== parsed.length) {
        for (let i of trackedIndexes) {
            let baseBeacon = parsed[i]
            for (let j = 0; j < parsed.length; j++) {
                if (trackedIndexes.some(k => j === k)) {
                    continue
                }

                let match;
                let otherBeacon = parsed[j]
                let otherBeaconPerms = getPermutations(otherBeacon)
                for (let permBeacon of otherBeaconPerms) {
                    match = compareBeaconCoords(baseBeacon, permBeacon.permBeacons)
                    if (match !== undefined) {
                        let c = match.common[1]
                        let pair = []
                        pair.push({index: i, coord: baseBeacon[c.i]})
                        pair.push({index: j, coord: permBeacon.permBeacons[c.j]})
                        let offset = [
                            (pair[0].coord[0] - pair[1].coord[0]),
                            (pair[0].coord[1] - pair[1].coord[1]),
                            (pair[0].coord[2] - pair[1].coord[2])
                        ]
                        transforms.push({from: j, to: i, offset, perm: permBeacon.perm})
                        trackedIndexes.push(j)
                        break
                    }
                }
            }
        }
    }

    return transforms
}

function getCoordRelativeToScannerZero(coord, index, transforms) {
    if (index === 0) {
        return [...coord]
    }

    let transformFrom = transforms.find(t => t.from === index)
    let permuted =  permuteCoordinate(coord, transformFrom.perm)
    let offset = transformFrom.offset
    let adjusted = [permuted[0] + offset[0], permuted[1] + offset[1], permuted[2] + offset[2]]
    // console.log(coord, adjusted, transformFrom.from, transformFrom.to)
    return getCoordRelativeToScannerZero(adjusted, transformFrom.to, transforms)
}

// how many possible in common... >= 12 => match
function compareBeaconCoords(beacon1, beacon2) {
    let d1 = getBeaconDistances(beacon1)
    let d2 = getBeaconDistances(beacon2)
    // console.log(d1, d2)
    for (let i = 0; i < d1.length; i++) {
        for (let j = 0; j < d2.length; j++) {
            let common = getNumInCommon(d1[i], d2[j])
            if (common.length >= 12) {
                return {i, j, common}
            }
        }
    }
}

function getNumInCommon(distances1, distances2) {
    let common = []
    for (let i = 0; i < distances1.length; i++) {
        let d1 = distances1[i]
        for (let j = 0; j < distances2.length; j++) {
            let d2 = distances2[j]
            if (d1[0] === d2[0] && d1[1] === d2[1] && d1[2] === d2[2]) {
                common.push({i, j})
            }
        }
    }

    return common
}

//
function getPermutations(beacons) {
    let permuted = []
    for (let forward of [0, 1, 2]) {
        for (let fRefl of [1, -1]) {
            for (let up of [2, 1, 0]) {
                if (forward === up) {
                    continue
                }

                for (let upRefl of [1, -1]) {
                    for (let backRefl of [1, -1]) {
                        let perm = {forward, fRefl, up, upRefl, backRefl}
                        let permBeacons = beacons.map(coord => permuteCoordinate(coord, perm))
                        permuted.push({perm, permBeacons})
                    }
                }
            }
        }
    }

    return permuted
}

function permuteCoordinate(coord, perm) {
    let newCoord = []
    newCoord[0] = perm.fRefl * coord[perm.forward]
    newCoord[2] = perm.upRefl * coord[perm.up]
    let back = [0, 1, 2].filter(d => d !== perm.forward && d !== perm.up)
    newCoord[1] = perm.backRefl * coord[back]
    return newCoord
}

// distances from x to z
function getBeaconDistances(beacons) {
    let distances = []
    for (let i = 0; i < beacons.length; i++) {
        let distancesForBeacon = []
        for (let j = 0; j < beacons.length; j++) {
            let distance = [
                beacons[i][0] - beacons[j][0],
                beacons[i][1] - beacons[j][1],
                beacons[i][2] - beacons[j][2],
            ]

            distancesForBeacon.push(distance)
        }

        distances.push(distancesForBeacon)
    }

    return distances
}

function parseInput(input) {
    let scanners = []
    let curBeacons = []
    for (let row of input) {
        let scanSplit = row.split('--- scanner ')
        if (scanSplit.length > 1) {
            // let scanNum = parseInt(scanSplit[1].substr(0, scanSplit[1].length - 4))
            curBeacons = []
        } else if (row.trim().length > 0) {
            curBeacons.push(row.split(',').map(Number))
        } else {
            scanners.push(curBeacons)
        }
    }

    return scanners
}
