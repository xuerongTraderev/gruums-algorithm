package algorithm

import (
	"github.com/gruums-common-go/container/stack"
)

type D2ArrayIndex struct {
	row int
	column int
}
/**
 *	DFS approach.
 */
func FloodFill(d2metrix [][]int, startPoint *D2ArrayIndex, color int) {
	stack := stack.New()

	exploredNodes := make(map[D2ArrayIndex]bool)
	originalColor := d2metrix[startPoint.row][startPoint.column]

	// Initial stack and exploreNodes
	stack.Push(startPoint)
	exploredNodes[*startPoint] = true

	for stack.Empty() != true {
		var currentNode *D2ArrayIndex = stack.Pop().(*D2ArrayIndex)
		d2metrix[currentNode.row][currentNode.column] = color
		var newNode *D2ArrayIndex
		// Find valid neighbor nodes: UP, DOWN, LEFT, RIGHT.
		// LEFT.
		if currentNode.column - 1 >= 0 && 
			d2metrix[currentNode.row][currentNode.column - 1] == originalColor &&
			! exploredNodes[D2ArrayIndex{currentNode.row,currentNode.column - 1}] {
			newNode = &D2ArrayIndex{currentNode.row, currentNode.column - 1}
			stack.Push(newNode)
			exploredNodes[*newNode] = true
		}
		// Right.
		if currentNode.column + 1 < len(d2metrix[currentNode.row]) && 
			d2metrix[currentNode.row][currentNode.column + 1] == originalColor &&
			! exploredNodes[D2ArrayIndex{currentNode.row,currentNode.column + 1}]{
			newNode = &D2ArrayIndex{currentNode.row, currentNode.column + 1}
			stack.Push(newNode)
			exploredNodes[*newNode] = true
		}
		// Top.
		if currentNode.row - 1 >= 0 && 
			d2metrix[currentNode.row - 1][currentNode.column] == originalColor &&
			! exploredNodes[D2ArrayIndex{currentNode.row - 1,currentNode.column}]{
			
			newNode = &D2ArrayIndex{currentNode.row - 1, currentNode.column}
			stack.Push(newNode)
			exploredNodes[*newNode] = true
		}
		// Down.
		if currentNode.row + 1 < len(d2metrix) && 
			d2metrix[currentNode.row + 1][currentNode.column] == originalColor &&
			! exploredNodes[D2ArrayIndex{currentNode.row + 1,currentNode.column}]{
			
			newNode = &D2ArrayIndex{currentNode.row + 1, currentNode.column}
			stack.Push(newNode)
			exploredNodes[*newNode] = true
		}

	}
}
