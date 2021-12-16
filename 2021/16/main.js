const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let hex = fs.readFileSync(inputPath).toString();

const hexMap = {
    '0': '0000',
    '1': '0001',
    '2': '0010',
    '3': '0011',
    '4': '0100',
    '5': '0101',
    '6': '0110',
    '7': '0111',
    '8': '1000',
    '9': '1001',
    'A': '1010',
    'B': '1011',
    'C': '1100',
    'D': '1101',
    'E': '1110',
    'F': '1111'
}

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let binary = hex.split('').map(h => hexMap[h]).join('')
    let packets = []
    parsePackets(binary, packets)
    let sum = sumVersions(packets)
    // console.log(JSON.stringify(packets, null, 2))
    return sum
}

function puzzle2() {
    let binary = hex.split('').map(h => hexMap[h]).join('')
    let packets = []
    parsePackets(binary, packets)
    console.log(JSON.stringify(packets, null, 2))
    calculateValues(packets[0])
    return packets[0].value
}

function calculateValues(packet) {
    if (packet.value !== undefined) {
        return
    }

    switch (packet.id) {
        case 4:
            packet.value = packet.literalValue
            break
        case 0:
            packet.subpackets.forEach(p => calculateValues(p))
            packet.value = packet.subpackets.map(p => p.value).reduce((a, b) => a + b)
            break
        case 1:
            packet.subpackets.forEach(p => calculateValues(p))
            packet.value = packet.subpackets.map(p => p.value).reduce((a, b) => a * b, 1)
            break
        case 2:
            packet.subpackets.forEach(p => calculateValues(p))
            packet.value = Math.min(...packet.subpackets.map(p => p.value))
            break
        case 3:
            packet.subpackets.forEach(p => calculateValues(p))
            packet.value = Math.max(...packet.subpackets.map(p => p.value))
            break
        case 5:
            packet.subpackets.forEach(p => calculateValues(p))
            if (packet.subpackets[0].value > packet.subpackets[1].value) {
                packet.value = 1
            } else {
                packet.value = 0
            }
            break
        case 6:
            packet.subpackets.forEach(p => calculateValues(p))
            if (packet.subpackets[0].value < packet.subpackets[1].value) {
                packet.value = 1
            } else {
                packet.value = 0
            }
            break
        case 7:
            packet.subpackets.forEach(p => calculateValues(p))
            if (packet.subpackets[0].value === packet.subpackets[1].value) {
                packet.value = 1
            } else {
                packet.value = 0
            }
            break
        default:
            throw packet.id
    }
}

function sumVersions(packets) {
    let sum = 0
    for (let packet of packets) {
        sum += packet.version
        sum += sumVersions(packet.subpackets)
    }

    return sum
}

function parsePackets(binary, packetsSoFar) {
    if (binary.split('').every(b => b === '0')) {
        return
    }

    let version = parseInt(binary.substr(0, 3), 2)
    let id = parseInt(binary.substr(3, 3), 2)
    let packet = {version, id, subpackets: []}
    packetsSoFar.push(packet)

    if (id === 4) {
        let literal = parseLiteral(binary)
        packet.literalValue = literal.value
        packet.literalLength = literal.length
        let remaining = binary.substring(packet.literalLength)
        parsePackets(remaining, packetsSoFar)
    } else {
        let lengthTypeID = binary.substr(6, 1)
        packet.lengthTypeID = lengthTypeID
        if (lengthTypeID === '0') {
            let length = parseInt(binary.substr(7, 15), 2)
            packet.lengthTypeValue = length
            let subpackets = binary.substr(22, length)
            let remaining = binary.substring(22 + length)
            parsePackets(subpackets, packet.subpackets)
            parsePackets(remaining, packetsSoFar)
        } else {
            let numPackets = parseInt(binary.substr(7, 11), 2)
            packet.lengthTypeValue = numPackets
            let subpackets = binary.substr(18)
            parsePackets(subpackets, packet.subpackets)
            let extra = packet.subpackets.splice(numPackets)
            for (let x of extra) {
                packetsSoFar.push(x)
            }
        }
    }
}

function parseLiteral(binary) {
    let literal = binary.substring(6)
    let literalPieces = []
    while (true) {
        let piece = literal.substr(literalPieces.length * 5, 5)
        if (piece.length < 5) {
            break
        }

        if (piece.substr(0, 1) === '1') {
            literalPieces.push(piece)
        } else {
            literalPieces.push(piece)
            break
        }
    }

    let length = (literalPieces.length * 5) + 6
    let literalValue = parseInt(literalPieces.map(p => p.substring(1)).join(''), 2)
    return { value: literalValue, length: length }
}
