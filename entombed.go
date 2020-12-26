package entombed

import "math/rand"

const (
	IS_WALL     = 1
	NO_WALL     = 0
	RANDOM_WALL = -1
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

/*
	int MAGIC(unsigned int lasttwo, unsigned int threeabove);
public:
	std::string render_line(unsigned int row);
	void rowGen(std::vector<unsigned int> &lastrows);
*/
