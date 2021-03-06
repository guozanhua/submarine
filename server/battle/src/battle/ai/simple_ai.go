package ai

import (
	"math"

	"github.com/ungerik/go3d/float64/vec2"

	"github.com/shiwano/submarine/server/battle/src/battle/scene"
)

// SimpleAI represents a simple battle AI.
type SimpleAI struct {
	*ai
	isNextDestStartPosition bool
}

// NewSimpleAI creates a SimpleAI.
func NewSimpleAI(scn scene.Scene) *SimpleAI {
	return &SimpleAI{ai: newAI(scn)}
}

// Update the SimpleAI.
func (a *SimpleAI) Update(submarine scene.Actor) {
	if !a.navigator.isStarted() {
		nextDest := a.nextDest(submarine.Player().StartPosition)
		path := a.scn.Stage().FindPath(submarine.Position(), nextDest)
		a.navigator.start(path, submarine.Position())
	}

	if ok, dir := a.navigator.navigate(submarine.Position()); ok {
		if math.Abs(dir-submarine.Direction()) > 0.01 {
			a.accelerateActor(submarine, dir)
		}
	} else {
		if submarine.IsAccelerating() {
			a.brakeActor(submarine, submarine.Direction())
		}
	}
}

func (a *SimpleAI) nextDest(startPosition *vec2.T) *vec2.T {
	var dest *vec2.T
	if a.isNextDestStartPosition {
		dest = &vec2.T{startPosition[0], startPosition[1]}
	} else {
		dest = &vec2.T{-startPosition[0], -startPosition[1]}
	}
	a.isNextDestStartPosition = !a.isNextDestStartPosition
	return dest
}
