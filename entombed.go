package entombed

import "math/rand"

const (
	isWall     = 1
	noWall     = 0
	randomWall = -1
)

func getRandomBit() uint {
	return uint(rand.Intn(2))
}

func isZeroPresent(lastRows []uint) bool {
	for _, row := range lastRows {
		if row == 0 {
			return true
		}
	}
	return false
}

func doTheMagicBitNow(lasttwo, threeabove uint) int {
	var response int
	lasttwo = lasttwo << 3
	switch lasttwo | threeabove {
	case 0b00000:
		response = isWall
		break
	case 0b00001:
		response = isWall
		break
	case 0b00010:
		response = isWall
		break
	case 0b00011:
		response = randomWall
		break
	case 0b00100:
		response = noWall
		break
	case 0b00101:
		response = noWall
		break
	case 0b00110:
		response = randomWall
		break
	case 0b00111:
		response = randomWall
		break
	case 0b01000:
		response = isWall
		break
	case 0b01001:
		response = isWall
		break
	case 0b01010:
		response = isWall
		break
	case 0b01011:
		response = isWall
		break
	case 0b01100:
		response = randomWall
		break
	case 0b01101:
		response = noWall
		break
	case 0b01110:
		response = noWall
		break
	case 0b01111:
		response = noWall
		break
	case 0b10000:
		response = isWall
		break
	case 0b10001:
		response = isWall
		break
	case 0b10010:
		response = isWall
		break
	case 0b10011:
		response = randomWall
		break
	case 0b10100:
		response = noWall
		break
	case 0b10101:
		response = noWall
		break
	case 0b10110:
		response = noWall
		break
	case 0b10111:
		response = noWall
		break
	case 0b11000:
		response = randomWall
		break
	case 0b11001:
		response = noWall
		break
	case 0b11010:
		response = isWall
		break
	case 0b11011:
		response = randomWall
		break
	case 0b11100:
		response = randomWall
		break
	case 0b11101:
		response = noWall
		break
	case 0b11110:
		response = noWall
		break
	case 0b11111:
		response = noWall
		break
	default:
		break
	}

	return response
}

// RenderLine prints an integer added to its mirrored bits as a string
func RenderLine(row uint) string {
	PF12 := ""

	for i := 0; i < 8; i++ {
		if row&1 == 1 {
			PF12 = "X" + PF12
		} else {
			PF12 = " " + PF12
		}
		row >>= 1
	}
	PF012 := "X" + PF12
	mirrorPF012 := ""
	for n := len(PF012) - 1; n >= 0; n-- {
		mirrorPF012 += string(PF012[n])
	}
	return PF012 + mirrorPF012
}

// GenerateRow takes a slice of 12 rows, removes the first row and adds one to the end
func GenerateRow(lastrows *[]uint) {}
