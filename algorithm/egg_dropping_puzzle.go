package algorithm
//import "fmt"
/**
 *	use DP approach
 *  floor and egg should be greater than zero
 */
func EggDroppingPuzzle(egg, floor uint64) uint64 {
	// row == egg.
	// column == floor.
	d2Array := make([][]uint64, egg)
	for row := range d2Array {
		d2Array[row] = make([]uint64, floor)
		// Initial first column.
		// If floor == 1, then egg will be 1.
		d2Array[row][0] = 1
	}

	// Initial first row.
	for column := range d2Array[0] {
		d2Array[0][column] = uint64(column + 1)
	}

	// Bottom Up.
	for row := uint64(1) ; row < uint64(len(d2Array)) ; row = row + 1 {
		for column := uint64(1)  ; column < uint64(len(d2Array[row])) ; column = column + 1 {
			// Consider first column to avoid out of array index
			// Worst case: first floor did not break then need to try rest of floor
			var min uint64 = d2Array[row][column-1] 

			for index := uint64(1)  ; index < column; index = index + 1 {
				var max uint64				

				if max = d2Array[row-1][index-1]; max < d2Array[row][column-index-1] {
					max = d2Array[row][column-index-1]
				}

				if max < min {
					min = max
				}
			}

			// Move last column out of for loop to avoid out of array index
			// Worst case: last floor break then need to try rest of floor by using minus one egg
			if min > d2Array[row - 1][column - 1] {
				min = d2Array[row - 1][column - 1]
			}
			
			d2Array[row][column] = min + 1
			
		}
	}
	return d2Array[egg-1][floor-1];
}