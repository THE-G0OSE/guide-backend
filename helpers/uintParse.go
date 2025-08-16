package helpers

import "strconv"

func UintParse(strid string) (uint, error) {
	id64, err := strconv.ParseUint(strid, 10, 64)
	id := uint(id64)
	return id, err
}
