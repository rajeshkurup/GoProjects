/**
 * @brief Coordinates of 2D matrix
 *
 * @author rajeshkurup@live.com
 */
package matrix

/**
 * @brief Cointainer for coordinates in a 2D matrix
 *
 */
type Pos struct {
	X int
	Y int
}

/**
 * @brief Compare two coordinates in 2D matrix
 *
 * @param src Source coordinates to be compared
 *
 * @return true if source coordinates is same as current {@link Pos} instance, false otherwise
 */
func (pos *Pos) IsEqual(src *Pos) bool {
	if src != nil {
		return pos.X == src.X && pos.Y == src.Y
	}

	return false
}
