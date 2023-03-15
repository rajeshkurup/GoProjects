/**
 * @brief Queue to store 2D coordinates
 *
 * @author rajeshkurup@live.com
 */
package matrix

import "errors"

/**
 * @brief Queue to store 2D coordinates
 */
type PosQueue struct {
	queue []Pos
}

/**
 * @brief Constructor for {@link PosQueue}
 *
 * @return Instance of {@link PosQueue}
 */
func MakePosQueue() *PosQueue {
	queue := make([]Pos, 0)
	return &PosQueue{queue}
}

/**
 * @brief Push new set of 2D coordinates into Queue
 *
 * @param p 2D coordinates to be added into Queue
 *
 * @return true if succeeded, error message if failed
 */
func (posQueue *PosQueue) Push(p *Pos) (bool, error) {

	if p == nil {
		return false, errors.New("Invalid argument")
	}

	posQueue.queue = append(posQueue.queue, *p)
	return true, nil
}

/**
 * @brief Remove a set of 2D coordinates from Queue
 *
 * @return Coordinates that was removed from Queue, error message if failed
 */
func (posQueue *PosQueue) Pop() (*Pos, error) {

	if len(posQueue.queue) == 0 {
		return nil, errors.New("Queue is empty")
	}

	pos := posQueue.queue[0]
	posQueue.queue = posQueue.queue[1:]
	return &pos, nil
}

/**
 * @brief Remove all coordinates from Queue
 */
func (posQueue *PosQueue) Clear() {
	posQueue.queue = make([]Pos, 0)
}

/**
 * @brief Remove a set of 2D coordinates from Queue
 *
 * @return Coordinates that are present in Queue, empty list if Queue is empty
 */
func (posQueue *PosQueue) Get() []Pos {
	return posQueue.queue
}

/**
 * @brief Returns the size of the Queue
 *
 * @return Number of elements in Queue
 */
func (posQueue *PosQueue) Size() int {
	return len(posQueue.queue)
}

/**
 * @brief Check whether given 2D coordinates are present in Queue
 *
 * @param pos 2D coordinates to be checked
 *
 * @return true if given 2D coordinates are present in Queue, false otherwise
 */
func (posQueue *PosQueue) Find(pos *Pos) bool {

	if pos != nil {
		for _, elem := range posQueue.queue {
			if pos.X == elem.X && pos.Y == elem.Y {
				return true
			}
		}
	}

	return false
}
