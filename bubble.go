package main

import (
	"math/rand"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	arraySize     = 750
	maxInt        = 100
	rectangleSize = 1
)

var (
	comparisons   = 0
	arrayAccesses = 0
	done          = false
	iterations    = 0
)

func main() {
	arr := []int{}
	for i := 0; i < arraySize; i++ {
		r := rand.Intn(maxInt)
		if r == 0 {
			arr = append(arr, 1)
		} else {
			arr = append(arr, rand.Intn(maxInt))
		}
	}

	rl.InitWindow(800, 450, "sorting")
	defer rl.CloseWindow()
	rl.SetTargetFPS(3000)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		for i := 0; i < len(arr); i++ {
			rl.DrawRectangle(int32(i), maxInt+10, rectangleSize, int32(arr[i]), rl.LightGray)
		}
		if !done {
			if !checkIfSorted(arr) {
				arr = selectionSort(arr)
				iterations++
			} else {
				done = true
			}
		}
		rl.DrawText("Bubble sort", 190, 70, 20, rl.LightGray)
		rl.DrawText("Comparisons: "+strconv.Itoa(comparisons), 190, 200, 20, rl.LightGray)
		rl.DrawText("Array Accesses: "+strconv.Itoa(arrayAccesses), 190, 230, 20, rl.LightGray)
		rl.DrawText("Iterations: "+strconv.Itoa(iterations), 190, 260, 20, rl.LightGray)
		rl.DrawFPS(20, 20)
		rl.EndDrawing()
	}
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr = swap(arr, i+1, i)
		}
		comparisons += 2
		arrayAccesses += 2
	}
	return arr
}

func selectionSort(arr []int) []int {
	arr = 
	return arr
}

func max(arr []int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if i > m {
			m = i
		}
	}
	return m
}

func checkIfSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
		comparisons += 2
		arrayAccesses += 2
	}
	return true
}

func swap(arr []int, i0, i1 int) []int {
	arr[i0], arr[i1] = arr[i1], arr[i0]
	arrayAccesses += 4
	return arr
}
