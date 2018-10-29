package dummyneko

import (
	"math"
)

type NekoState struct {
	X, Y float64
	Action ActionName
}

type MouseState struct {
	X, Y float64
}

type NekoBehavior struct {
	// speed per single action/tick
	Step float64
	// max distance between mouse and neko
	Dmax float64

	// ticks per state

	StillTicks uint
	YawnTicks uint
	YawnStillTicks uint
	SleepTicks uint
	AlertTicks uint
}

type ActionName string

const (
	ActionAlert = "alert"
	ActionStill = "still"
	ActionYawn = "yawn"
	ActionSleep1 = "sleep1"
	ActionSleep2 = "sleep2"
	// run
	ActionNRun1 = "nrun1"
	ActionNRun2 = "nrun2"
	ActionNERun1 = "nerun1"
	ActionNERun2 = "nerun2"
	ActionERun1 = "erun1"
	ActionERun2 = "erun2"
	ActionSERun1 = "serun1"
	ActionSERun2 = "serun2"
	ActionSRun1 = "srun1"
	ActionSRun2 = "srun2"
	ActionSWRun1 = "swrun1"
	ActionSWRun2 = "swrun2"
	ActionWRun1 = "wrun1"
	ActionWRun2 = "wrun2"
	ActionNWRun1 = "nwrun1"
	ActionNWRun2 = "nwrun2"
)

var (
	InitialState ActionState = stateStill{}
	DefaultBehavior = NekoBehavior{
		Step: 15,
		Dmax: 15,
		StillTicks: 15,
		YawnTicks: 4,
		YawnStillTicks: 1,
		SleepTicks: 2,
		AlertTicks: 2,
	}
)

const (
	π = math.Pi
)

// dirActionName returns action name given the direction (see direction function) and whether the action is even.
//
// Returns values are hardcoded because go tool gives better coverage reports this way.
func dirActionName(dir string, even bool) ActionName {
	switch dir {
	case "n":
		if even {
			return ActionNRun2
		} else {
			return ActionNRun1
		}
	case "ne":
		if even {
			return ActionNERun2
		} else {
			return ActionNERun1
		}
	case "E":
		if even {
			return ActionERun2
		} else {
			return ActionERun1
		}
	case "se":
		if even {
			return ActionSERun2
		} else {
			return ActionSERun1
		}
	case "s":
		if even {
			return ActionSRun2
		} else {
			return ActionSRun1
		}
	case "sw":
		if even {
			return ActionSWRun2
		} else {
			return ActionSWRun1
		}
	case "w":
		if even {
			return ActionWRun2
		} else {
			return ActionWRun1
		}
	case "nw":
		if even {
			return ActionNWRun2
		} else {
			return ActionNWRun1
		}
	}
	return ""
}

// direction function computes the compass direction to the destination given the current (x,y) and the destination (mx,my) coordinates.
//
// Note that it assumes a coordinate system with inverted Y axis, i.e. with origin at the top-left corner.  Invoke it with inverted ordinate sign to use the classic cartesian coordinate system.
//
// Returns values are hardcoded because go tool gives better coverage reports this way.
func direction(x, y, mx, my float64) string {
	dx := x - mx
	dy := y - my
	α := math.Atan2(dy, dx)
	switch {
	case -π * 8/8 <= α && α <= -π * 7/8:
		return "e"
	case -π * 7/8 <= α && α <= -π * 5/8:
		return "se"
	case -π * 5/8 <= α && α <= -π * 3/8:
		return "s"
	case -π * 3/8 <= α && α <= -π * 1/8:
		return "sw"
	case -π * 1/8 <= α && α <= +π * 1/8:
		return "w"
	case +π * 1/8 <= α && α <= +π * 3/8:
		return "nw"
	case +π * 3/8 <= α && α <= +π * 5/8:
		return "n"
	case +π * 5/8 <= α && α <= +π * 7/8:
		return "ne"
	case +π * 7/8 <= α && α <= +π * 8/8:
		return "e"
	}
	return ""
}

func pointerNearby(n NekoState, m MouseState, b NekoBehavior) bool {
	dx := n.X - m.X
	dy := n.Y - m.Y
	d := math.Hypot(dx, dy)
	return d <= b.Dmax
}

func makeStep(n *NekoState, m MouseState, b NekoBehavior) {
	dx := n.X - m.X
	dy := n.Y - m.Y
	d := math.Hypot(dx, dy)
	if d > 0 {
		dstep := b.Step / d
		n.X -= dstep * dx
		n.Y -= dstep * dy
		return
	}
	return
}

type ActionState interface {
	Next(NekoState, MouseState, NekoBehavior) ActionState
	Render(NekoState, MouseState, NekoBehavior) NekoState
}

type stateStill struct {
	tick uint
}

func (s stateStill) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.StillTicks {
		return stateYawn{}
	}
	return s
}

func (s stateStill) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	n.Action = ActionStill
	return n
}

type stateYawn struct {
	tick uint
}

func (s stateYawn) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.YawnTicks {
		return stateYawnStill{}
	}
	return s
}

func (s stateYawn) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	n.Action = ActionYawn
	return n
}

type stateYawnStill struct {
	tick uint
}

func (s stateYawnStill) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.YawnStillTicks {
		return stateSleep{}
	}
	return s
}

func (s stateYawnStill) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	n.Action = ActionStill
	return n
}

type stateSleep struct {
	tick uint
	even bool
}

func (s stateSleep) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.SleepTicks {
		s.tick = 0
		s.even = !s.even
	}
	return s
}

func (s stateSleep) Render(n NekoState, m MouseState, b NekoBehavior)  NekoState {
	if s.even {
		n.Action = ActionSleep2
	} else {
		n.Action = ActionSleep1
	}
	return n
}

type stateAlert struct {
	tick uint
}

func (s stateAlert) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if pointerNearby(n, m, b) {
		return stateStill{}
	}
	s.tick += 1
	if s.tick >= b.AlertTicks {
		return stateRun{}
	}
	return s
}

func (s stateAlert) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	n.Action = ActionAlert
	return n
}

type stateRun struct {
	even bool
}

func (s stateRun) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	if pointerNearby(n, m, b) {
		return stateStill{}
	}
	s.even = !s.even
	return s
}

func (s stateRun) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	if !pointerNearby(n, m, b) {
		dir := direction(n.X, n.Y, m.X, m.Y)
		n.Action = dirActionName(dir, s.even)
	}
	makeStep(&n, m, b)
	return n
}
