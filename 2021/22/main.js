const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')
let instructions = parseInput(input)
let debug = false

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let flipped = {}
    for (let ins of instructions) {
        for (let x = Math.max(ins.x[0], -50); x <= Math.min(50, ins.x[1]); x++) {
            for (let y = Math.max(ins.y[0], -50); y <= Math.min(50, ins.y[1]); y++) {
                for (let z = Math.max(ins.z[0], -50); z <= Math.min(50, ins.z[1]); z++) {
                    let key = `${x}_${y}_${z}`
                    flipped[key] = ins.flip
                }
            }
        }
    }

    return Object.values(flipped).filter(v => v === 1).length
}

function puzzle2() {
    // disjoint regions
    let regionsOn = []
    for (let instruction of instructions) {
        regionsOn = flipRegion(regionsOn, instruction)
    }

    let n = countFast(regionsOn)
    return n
}

function flipRegion(regionsOn, instruction) {
    let instructionRegion = getRegion(instruction.x, instruction.y, instruction.z)
    if (!regionsOn.length) {
        if (instruction.flip) {
            return [instructionRegion]
        } else {
            return []
        }
    }

    if (instruction.flip) {
        let newRegionsOn = [{...instructionRegion}]

        for (let region of regionsOn) {
            let removeDoubleCountedRegion = deleteFromCube(instructionRegion, region)
            if (debug && countFast([region]) < countFast(removeDoubleCountedRegion)) {
                throw 'removed double counted region, but gained points'
            }

            for (let cube of removeDoubleCountedRegion) {
                newRegionsOn.push({...cube})
            }
        }

        if (debug && countFast(regionsOn) > countFast(newRegionsOn)) {
            throw 'added region and removed double counted, but lost points'
        }

        return [...newRegionsOn]
    } else {
        let newRegionsOn = []
        for (let region of regionsOn) {
            let removeDoubleCounted = deleteFromCube(instructionRegion, region)
            for (let cube of removeDoubleCounted) {
                newRegionsOn.push({...cube})
            }
        }

        return [...newRegionsOn]
    }
}

function countFast(newRegions) {
    let n = 0
    for (let region of newRegions) {
        let regionPoints = (region.x[1] - region.x[0] + 1) * (region.y[1] - region.y[0] + 1) * (region.z[1] - region.z[0] + 1)
        n += regionPoints
    }

    return n
}

function getRegion(x, y, z) {
    return {x, y, z}
}

function deleteFromCube(deleteCube, baseCube) {
    if (deleteCube.x[1] < baseCube.x[0] || deleteCube.x[0] > baseCube.x[1]) {
        return [{...baseCube}]
    } else if (deleteCube.x[1] < baseCube.x[1] || deleteCube.x[0] > baseCube.x[0]) {
        // baseCube not entirely contained in deleteCube
        let regionsAfterDelete = []
        if (baseCube.x[0] < deleteCube.x[0]) {
            let beforeX = [baseCube.x[0], deleteCube.x[0] - 1]
            regionsAfterDelete.push(getRegion(beforeX, baseCube.y, baseCube.z))
        }

        if (baseCube.x[1] > deleteCube.x[1]) {
            let afterX = [deleteCube.x[1] + 1, baseCube.x[1]]
            regionsAfterDelete.push(getRegion(afterX, baseCube.y, baseCube.z))
        }

        let middleX = [Math.max(deleteCube.x[0], baseCube.x[0]), Math.min(deleteCube.x[1], baseCube.x[1])]
        let middleRegions = deleteFromCube(deleteCube, getRegion(middleX, baseCube.y, baseCube.z))
        regionsAfterDelete = regionsAfterDelete.concat([...middleRegions])
        return [...regionsAfterDelete]
    }

    if (deleteCube.y[1] < baseCube.y[0] || deleteCube.y[0] > baseCube.y[1]) {
        return [{...baseCube}]
    } else if (deleteCube.y[1] < baseCube.y[1] || deleteCube.y[0] > baseCube.y[0]) {
        // baseCube not entirely contained in deleteCube
        let regionsAfterDelete = []
        if (baseCube.y[0] < deleteCube.y[0]) {
            let beforeY = [baseCube.y[0], deleteCube.y[0] - 1]
            regionsAfterDelete.push(getRegion(baseCube.x, beforeY, baseCube.z))
        }

        if (baseCube.y[1] > deleteCube.y[1]) {
            let afterY = [deleteCube.y[1] + 1, baseCube.y[1]]
            regionsAfterDelete.push(getRegion(baseCube.x, afterY, baseCube.z))
        }

        let middleY = [Math.max(deleteCube.y[0], baseCube.y[0]), Math.min(deleteCube.y[1], baseCube.y[1])]
        let middleRegions = deleteFromCube(deleteCube, getRegion(baseCube.x, middleY, baseCube.z))
        regionsAfterDelete = regionsAfterDelete.concat([...middleRegions])
        return [...regionsAfterDelete]
    }

    if (deleteCube.z[1] < baseCube.z[0] || deleteCube.z[0] > baseCube.z[1]) {
        return [{...baseCube}]
    } else if (deleteCube.z[1] < baseCube.z[1] || deleteCube.z[0] > baseCube.z[0]) {
        let regionsAfterDelete = []
        if (baseCube.z[0] < deleteCube.z[0]) {
            let beforeZ = [baseCube.z[0], deleteCube.z[0] - 1]
            regionsAfterDelete.push(getRegion(baseCube.x, baseCube.y, beforeZ))
        }

        if (baseCube.z[1] > deleteCube.z[1]) {
            let afterZ = [deleteCube.z[1] + 1, baseCube.z[1]]
            regionsAfterDelete.push(getRegion(baseCube.x, baseCube.y, afterZ))
        }

        let middleZ = [Math.max(deleteCube.z[0], baseCube.z[0]), Math.min(deleteCube.z[1], baseCube.z[1])]
        let middleRegions = deleteFromCube(deleteCube, getRegion(baseCube.x, baseCube.y, middleZ))
        regionsAfterDelete = regionsAfterDelete.concat([...middleRegions])
        return [...regionsAfterDelete]
    }

    return []
}

function parseInput(input) {
    let instructions = []
    for (let row of input) {
        if (row.length === 0) {
            break
        }

        let flip = row.startsWith("on") ? 1 : 0
        let x = row.split("x=")[1].split(",")[0].split("..").map(Number)
        let y = row.split("y=")[1].split(",")[0].split("..").map(Number)
        let z = row.split("z=")[1].split(",")[0].split("..").map(Number)
        let instruction = {flip, x, y, z}
        instructions.push(instruction)
    }

    return instructions
}
