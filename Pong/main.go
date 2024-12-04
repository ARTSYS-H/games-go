package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// === Initialization ===
	const screenWidth = 800
	const screenHeight = 450

	// Init Window
	rl.SetConfigFlags(rl.FlagWindowHighdpi | rl.FlagMsaa4xHint)
	rl.InitWindow(screenWidth, screenHeight, "[Raylib] Pong")
	defer rl.CloseWindow()

	// Init Sounds
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	hitPong := rl.LoadSound("assets/sounds/001.wav")
	defer rl.UnloadSound(hitPong)
	hitPaddle := rl.LoadSound("assets/sounds/002.wav")
	defer rl.UnloadSound(hitPaddle)
	hitLoose := rl.LoadSound("assets/sounds/003.wav")
	defer rl.UnloadSound(hitLoose)

	// Init Paddles
	var paddleWidth, paddleHeight float32 = 15.0, 70.0
	leftPaddle := rl.NewRectangle(50, screenHeight/2-paddleHeight/2, paddleWidth, paddleHeight)
	rightPaddle := rl.NewRectangle(screenWidth-50-paddleWidth, screenHeight/2-paddleHeight/2, paddleWidth, paddleHeight)

	// Init Ball
	ball := rl.NewRectangle(screenWidth/2, screenHeight/2, 10, 10)
	ballSpeed := rl.NewVector2(6, 6)

	// Init Scores
	var leftScore, rightScore int

	// Paddle's speed
	var paddleSpeed float32 = 6.0

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		// === Logiq ===

		// Player 1 controls (left paddle)
		if rl.IsKeyDown(rl.KeyW) && leftPaddle.Y > 0 {
			leftPaddle.Y -= paddleSpeed
		}
		if rl.IsKeyDown(rl.KeyS) && leftPaddle.Y < screenHeight-paddleHeight {
			leftPaddle.Y += paddleSpeed
		}

		// Player 2 controls (right paddle)
		if rl.IsKeyDown(rl.KeyUp) && rightPaddle.Y > 0 {
			rightPaddle.Y -= paddleSpeed
		}
		if rl.IsKeyDown(rl.KeyDown) && rightPaddle.Y < screenHeight-paddleHeight {
			rightPaddle.Y += paddleSpeed
		}

		// Moving the ball!
		ball.X += ballSpeed.X
		ball.Y += ballSpeed.Y

		// Walls Hit
		if ball.Y <= 0 || ball.Y+ball.Height >= screenHeight {
			ballSpeed.Y *= -1
			rl.PlaySound(hitPong)
		}

		// Paddles Hit
		if rl.CheckCollisionRecs(ball, leftPaddle) {
			ballSpeed.X *= -1 // Reverse horizontal direction
			// Calcul angle from relative position
			relativeImpact := (ball.Y + ball.Height/2) - (leftPaddle.Y + leftPaddle.Height/2)
			ballSpeed.Y = relativeImpact * 0.2 // Adjust vertical speed from impact position
			rl.PlaySound(hitPaddle)
		}

		if rl.CheckCollisionRecs(ball, rightPaddle) {
			ballSpeed.X *= -1
			relativeImpact := (ball.Y + ball.Height/2) - (rightPaddle.Y + rightPaddle.Height/2)
			ballSpeed.Y = relativeImpact * 0.2
			rl.PlaySound(hitPaddle)
		}

		// Scores detection
		if ball.X <= 0 {
			rightScore++
			ball = rl.NewRectangle(screenWidth/2, screenHeight/2, 10, 10)
			ballSpeed = rl.NewVector2(6, 6)
			rl.SetSoundPan(hitLoose, 0.8)
			rl.PlaySound(hitLoose)
		}
		if ball.X+ball.Width >= screenWidth {
			leftScore++
			ball = rl.NewRectangle(screenWidth/2, screenHeight/2, 10, 10)
			ballSpeed = rl.NewVector2(-6, -6)
			rl.SetSoundPan(hitLoose, 0.2)
			rl.PlaySound(hitLoose)
		}

		// === RENDER ===
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		// Drawing paddles and ball
		rl.DrawRectangleRec(leftPaddle, rl.White)
		rl.DrawRectangleRec(rightPaddle, rl.White)
		rl.DrawRectangleRec(ball, rl.White)

		// Draw Scores
		rl.DrawText(fmt.Sprintf("%d", leftScore), screenWidth/4, 20, 40, rl.White)
		rl.DrawText(fmt.Sprintf("%d", rightScore), 3*screenWidth/4, 20, 40, rl.White)

		// Draw middle line
		pass := true
		for i := 0; i < screenHeight; i += 10 {
			if pass {
				pass = !pass
				continue
			}
			rl.DrawRectangle(screenWidth/2, int32(i), 3, 10, rl.White)
			pass = !pass
		}

		rl.EndDrawing()
	}
}
