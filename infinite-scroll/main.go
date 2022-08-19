package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/oakmound/oak/v4/key"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/scene"
)

const (
	// The only collision label we need for this demo is 'ground',
	// indicating something we shouldn't be able to fall or walk through
	Ground   collision.Label = 1
	Obstacle collision.Label = 2
	WaitTime time.Duration   = 50 * time.Millisecond
)

func main() {
	oak.AddScene("firstScene", scene.Scene{
		Start: func(sc *scene.Context) {
			// ... draw entities, bind callbacks ...

			sc.Window.SetTitle("Infinite Dash")

			char := entities.New(sc,
				entities.WithRect(floatgeom.NewRect2WH(100, 100, 16, 32)),
				entities.WithColor(color.RGBA{255, 0, 0, 255}),
				entities.WithSpeed(floatgeom.Point2{3, 7}),
			)

			const fallSpeed = .2

			event.Bind(sc, event.Enter, char, func(c *entities.Entity, ev event.EnterPayload) event.Response {

				oldX, oldY := char.X(), char.Y()
				char.ShiftDelta()
				aboveGround := false

				hit := collision.HitLabel(char.Space, Ground)
				gameOver := collision.HitLabel(char.Space, Obstacle)

				if gameOver != nil {
					fmt.Println("Game Over")
					os.Exit(0)
				}

				// If we've moved in y value this frame and in the last frame,
				// we were below what we're trying to hit, we are still falling
				if hit != nil && !(oldY != char.Y() && oldY+char.H() > hit.Y()) {
					// Correct our y if we started falling into the ground
					char.SetY(hit.Y() - char.H())
					// Stop falling
					char.Delta[1] = 0
					// Jump with Space when on the ground
					if oak.IsDown(key.Spacebar) {
						char.Delta[1] -= char.Speed.Y()
					}
					aboveGround = true
				} else {

					// Fall if there's no ground
					char.Delta[1] += fallSpeed
				}

				if hit != nil {
					// If we walked into a piece of ground, move back
					xover, yover := char.Space.Overlap(hit)
					// We, perhaps unintuitively, need to check the Y overlap, not
					// the x overlap
					// if the y overlap exceeds a superficial value, that suggests
					// we're in a state like
					//
					// G = Ground, C = Character
					//
					// GG C
					// GG C
					//
					// moving to the left
					if math.Abs(yover) > 1 {
						// We add a buffer so this doesn't retrigger immediately
						xbump := 1.0
						if xover > 0 {
							xbump = -1
						}
						char.SetX(oldX + xbump)
						if char.Delta.Y() < 0 {
							char.Delta[1] = 0
						}
					}

					// If we're below what we hit and we have significant xoverlap, by contrast,
					// then we're about to jump from below into the ground, and we
					// should stop the character.
					if !aboveGround && math.Abs(xover) > 1 {
						// We add a buffer so this doesn't retrigger immediately
						char.SetY(oldY + 1)
						char.Delta[1] = fallSpeed
					}

				}

				return 0
			})

			platform := floatgeom.NewRect2WH(0, 420, 640, 20)

			entities.New(sc,
				entities.WithRect(platform),
				entities.WithColor(color.RGBA{0, 168, 0, 255}),
				entities.WithLabel(Ground),
			)

			go func() {

				for {

					wait := rand.Intn(20)
					timeOut := time.Duration(wait) * time.Second

					time.Sleep(timeOut)

					fmt.Println("Spawn")
					obs := entities.New(sc,
						entities.WithRect(floatgeom.NewRect2WH(640, 390, 20, 32)),
						entities.WithColor(color.RGBA{100, 100, 0, 255}),
						entities.WithSpeed(floatgeom.Point2{3, 7}),
						entities.WithLabel(Obstacle),
					)

					go HandleEnt(obs)

				}

			}()

		},
	})
	oak.Init("firstScene")
}

func HandleEnt(e *entities.Entity) {

	for {

		time.Sleep(WaitTime)

		// set speed to move left, hence `-` sign
		e.Delta[0] = -e.Speed.X()

		// perform shift
		e.ShiftDelta()

		// if x is less then 0, we're out of
		// the screen. I'll check if -10 so the entity
		// is destroyed out of screen.
		if e.X() < -10 {
			break
		}

	}

	fmt.Println("Destroying")

	e.Destroy()

}
