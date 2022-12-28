package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/broyojo/sorting/sort"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	min       = 1
	max       = 500
	arraySize = 1000
	width     = 1280
	height    = 720
)

var comparisons, accesses, iterations int

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr []int
	for i := 0; i < arraySize; i++ {
		arr = append(arr, rand.Intn(max-min+1)+min)
	}

	rl.InitWindow(width, height, "sorting visualization")
	defer rl.CloseWindow()
	rl.SetTargetFPS(10000)

	barWidth := int32(width / int32(len(arr)))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawFPS(10, 10)
		rl.DrawText("Number of Elements: "+strconv.Itoa(arraySize), 10, 40, 20, rl.White)
		rl.DrawText("Comparisons: "+strconv.Itoa(comparisons), 10, 70, 20, rl.White)
		rl.DrawText("Array Accesses: "+strconv.Itoa(accesses), 10, 100, 20, rl.White)
		rl.DrawText("Iterations: "+strconv.Itoa(iterations), 10, 130, 20, rl.White)

		for k, v := range arr {
			rl.DrawRectangle(barWidth*int32(k), height-int32(v), barWidth, int32(v), rl.White)
		}

		if !sort.CheckSorted(arr) {
			arr, comparisons, accesses = sort.Bubble(arr, iterations%len(arr))
			iterations++
		}

		rl.EndDrawing()
	}
}
