package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gmae199boy/card_dungeon/src/setting"
)

func main() {
	rl.InitWindow(setting.WINDOW_WIDTH, setting.WINDOW_HEIGHT, "raylib [core] example - basic window")

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
