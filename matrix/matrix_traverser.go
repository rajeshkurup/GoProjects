/**
 * @brief Traverse 2D matrix to find shortest path.
 *
 * @author rajeshkurup@live.com
 */
package matrix

import (
	"errors"
)

/**
 * @brief Traverse 2D matrix to find shortest path.
 */
type MatrixTraverser struct {
	matrix [][]int
}

/**
 * @brief Constructor for {@link MatrixTraverser}
 *
 * @return Instance of {@link MatrixTraverser} if succeeded, error message if failed
 */
func MakeMatrixTraverser(matrix [][]int) (*MatrixTraverser, error) {

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil, errors.New("Invalid matrix")
	}

	matrix2d := make([][]int, 0)
	for _, mat := range matrix {
		row := make([]int, len(mat))
		copy(row, mat)
		matrix2d = append(matrix2d, row)
	}

	return &MatrixTraverser{matrix2d}, nil
}

/**
 * @brief Find shorest path to reach from start to end in a given 2D matrix. True cells represent blocked route.
 *
 * @param matrix 2D Matrics to be traversed
 * @param start Start position (coordinates)
 * @param end End position (coordinates)
 *
 * @return Collection of {@link Pos} instances contains shortest path from start to end
 */
func (matrixTraverser *MatrixTraverser) FindShortestPath(start *Pos, end *Pos) ([]Pos, error) {

	if matrixTraverser.matrix == nil || start == nil || end == nil || len(matrixTraverser.matrix) == 0 {
		return nil, errors.New("Invalid arguments")
	}

	posFound := matrixTraverser.findPos(start)
	if !posFound {
		return nil, errors.New("Start position not found in Matrix")
	}

	posFound = matrixTraverser.findPos(end)
	if !posFound {
		return nil, errors.New("End position not found in Matrix")
	}

	current := start
	paths := make([][]Pos, 0)
	paths = append(paths, append(make([]Pos, 0), *current))
	currentPathIdx := 0

	for !matrixTraverser.isVisitDone() {
		newPaths := matrixTraverser.getAdjCells(current, &paths[currentPathIdx])
		for _, newPath := range newPaths {
			paths = append(paths, newPath)
		}

		if !current.IsEqual(end) {
			matrixTraverser.matrix[current.X][current.Y] = 1
		}

		noMoreCells := true
		for idx, currentPath := range paths {

			lastPos := currentPath[len(currentPath)-1]
			if !lastPos.IsEqual(end) && matrixTraverser.matrix[lastPos.X][lastPos.Y] == 0 {
				current = &lastPos
				currentPathIdx = idx
				noMoreCells = false
				break
			}
		}

		if noMoreCells {
			matrixTraverser.matrix[end.X][end.Y] = 1
		}
	}

	return matrixTraverser.getShortestPath(paths, end)
}

func (matrixTraverser *MatrixTraverser) getShortestPath(paths [][]Pos, end *Pos) ([]Pos, error) {

	minLen := 0
	minLenIdx := 0
	for idx, path := range paths {

		if len(path) > 0 && path[len(path)-1].IsEqual(end) {
			if minLen == 0 {
				minLen = len(path)
				minLenIdx = idx
			} else if len(path) < minLen {
				minLen = len(path)
				minLenIdx = idx
			}
		}
	}

	if minLen == 0 {
		return nil, errors.New("No path found")
	}

	return paths[minLenIdx], nil
}

func (matrixTraverser *MatrixTraverser) isVisitDone() bool {

	for i := 0; i < len(matrixTraverser.matrix); i++ {
		for j := 0; j < len(matrixTraverser.matrix[i]); j++ {
			if matrixTraverser.matrix[i][j] == 0 {
				return false
			}
		}
	}

	return true
}

func (matrixTraverser *MatrixTraverser) getAdjCells(pos *Pos, path *[]Pos) [][]Pos {

	adjRight := pos.X + 1
	adjLeft := pos.X - 1
	adjTop := pos.Y - 1
	adjBottom := pos.Y + 1

	adjCells := make([]Pos, 0)
	if adjRight < len(matrixTraverser.matrix) && matrixTraverser.matrix[adjRight][pos.Y] == 0 && matrixTraverser.matrix[adjRight][pos.Y] == 0 {
		adjCells = append(adjCells, Pos{X: adjRight, Y: pos.Y})
	}

	if adjBottom < len(matrixTraverser.matrix[pos.X]) && matrixTraverser.matrix[pos.X][adjBottom] == 0 && matrixTraverser.matrix[pos.X][adjBottom] == 0 {
		adjCells = append(adjCells, Pos{X: pos.X, Y: adjBottom})
	}

	if adjLeft >= 0 && matrixTraverser.matrix[adjLeft][pos.Y] == 0 && matrixTraverser.matrix[adjLeft][pos.Y] == 0 {
		adjCells = append(adjCells, Pos{X: adjLeft, Y: pos.Y})
	}

	if adjTop >= 0 && matrixTraverser.matrix[pos.X][adjTop] == 0 && matrixTraverser.matrix[pos.X][adjTop] == 0 {
		adjCells = append(adjCells, Pos{X: pos.X, Y: adjTop})
	}

	paths := make([][]Pos, 0)
	for i := 0; i < len(adjCells); i++ {
		newPath := make([]Pos, len(*path))
		copy(newPath, *path)
		paths = append(paths, append(newPath, adjCells[i]))
	}

	return paths
}

/**
 * @brief Check whether given coordinates are present in Matrix
 *
 * @param matrix 2D Matrics to be traversed
 * @param pos Coordinates to be checked
 *
 * @return true if coordinates found, false otherwise
 */
func (matrixTraverser *MatrixTraverser) findPos(pos *Pos) bool {

	for x := 0; x < len(matrixTraverser.matrix); x++ {
		if x == pos.X {
			for y := 0; y < len(matrixTraverser.matrix[x]); y++ {
				if y == pos.Y {
					return true
				}
			}
		}
	}

	return false
}
