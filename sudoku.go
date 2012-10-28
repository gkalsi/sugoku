// Package sudoku provdes an interface for solving 9x9 sudoku puzzles
package sudoku

// field is a slice of 9 element arrays.
// Each element represent a set of indices 
var field [][9]int

func init() {
	field = make([][9]int, 27)

	// Squares
	field[0] = [...]int{1, 2, 3, 10, 11, 12, 19, 20, 21}
	field[1] = [...]int{4, 5, 6, 13, 14, 15, 22, 23, 24}
	field[2] = [...]int{7, 8, 9, 16, 17, 18, 25, 26, 27}

	field[3] = [...]int{28, 29, 30, 37, 38, 39, 46, 47, 48}
	field[4] = [...]int{31, 32, 33, 40, 41, 42, 49, 50, 51}
	field[5] = [...]int{34, 35, 36, 43, 44, 45, 52, 53, 54}

	field[6] = [...]int{55, 56, 57, 64, 65, 66, 73, 74, 75}
	field[7] = [...]int{58, 59, 60, 67, 68, 69, 76, 77, 78}
	field[8] = [...]int{61, 62, 63, 70, 71, 72, 79, 80, 81}

	// Rows
	field[9] = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	field[10] = [...]int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	field[11] = [...]int{19, 20, 21, 22, 23, 24, 25, 26, 27}

	field[12] = [...]int{28, 29, 30, 31, 32, 33, 34, 35, 36}
	field[13] = [...]int{37, 38, 39, 40, 41, 42, 43, 44, 45}
	field[14] = [...]int{46, 47, 48, 49, 50, 51, 52, 53, 54}

	field[15] = [...]int{55, 56, 57, 58, 59, 60, 61, 62, 63}
	field[16] = [...]int{64, 65, 66, 67, 68, 69, 70, 71, 72}
	field[17] = [...]int{73, 74, 75, 76, 77, 78, 79, 80, 81}

	// Columns
	field[18] = [...]int{1, 10, 19, 28, 37, 46, 55, 64, 73}
	field[19] = [...]int{2, 11, 20, 29, 38, 47, 56, 65, 74}
	field[20] = [...]int{3, 12, 21, 30, 39, 48, 57, 66, 75}

	field[21] = [...]int{4, 13, 22, 31, 40, 49, 58, 67, 76}
	field[22] = [...]int{5, 14, 23, 32, 41, 50, 59, 68, 77}
	field[23] = [...]int{6, 15, 24, 33, 42, 51, 60, 69, 78}

	field[24] = [...]int{7, 16, 25, 34, 43, 52, 61, 70, 79}
	field[25] = [...]int{8, 17, 26, 35, 44, 53, 62, 71, 80}
	field[26] = [...]int{9, 18, 27, 36, 45, 54, 63, 72, 81}
}
// Solve takes a slice of ints representing a sudoku board.
// It returns true if the puzzle was successfully solved, 
// false otherwise.
func Solve(arr []int) bool {
	result := solveRecursive(arr, 0)
	return result
}


// solveRecursive is the implementation of the Solve method.
// It uses recursive backtracking to solve the puzzle.
func solveRecursive(arr []int, index int) bool {
	if index == 81 {
		return check(arr)	// If the board is full and the current configuration is valid, the puzzle is complete
	} else if arr[index] == 0 {
		for i := 0; i < 9; i++ {
			arr[index] = i + 1
			if check(arr) {
				if solveRecursive(arr, i) {
					return true
				}
			}
		}
		arr[index] = 0
		return false
	} else {
		return solveRecursive(arr, index+1)
	}
	return false
}

// check Checks to see if a given sudoku puzzle is correct.
// It will return true if the puzzle's current state is legal
// even if the puzzle is not fully complete.
func check(arr []int) bool {
	var used [10]bool
	for _, f := range field {
		for i := 0; i <= 9; i++ {
			used[i] = false
		}
		for i := 0; i < 9; i++ {
			if arr[f[i]-1] == 0 {
				continue
			}
			if used[arr[f[i]-1]] == true {
				return false
			}
			used[arr[f[i]-1]] = true
		}
	}
	return true
}
