package dummyneko

import (
	"math"
)

type NekoState struct {
	X, Y float64
	Action Action
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

type Action string

const (
	ActionAlert = "alert"
	ActionStill = "still"
	ActionYawn = "yawn"
	// itch
	ActionItch1 = "itch1"
	ActionItch2 = "itch2"
	// sleep
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
	// scratch
	ActionNScratch1 = "nscratch1"
	ActionNScratch2 = "nscratch2"
	ActionEScratch1 = "escratch1"
	ActionEScratch2 = "escratch2"
	ActionSScratch1 = "sscratch1"
	ActionSScratch2 = "sscratch2"
	ActionWScratch1 = "wscratch1"
	ActionWScratch2 = "wscratch2"
)

var (
	// SupportedActions is a list of implemented action.
	SupportedActions = []Action{
		ActionAlert,
		ActionStill,
		ActionYawn,
		ActionSleep1,
		ActionSleep2,
		ActionNRun1,
		ActionNRun2,
		ActionNERun1,
		ActionNERun2,
		ActionERun1,
		ActionERun2,
		ActionSERun1,
		ActionSERun2,
		ActionSRun1,
		ActionSRun2,
		ActionSWRun1,
		ActionSWRun2,
		ActionWRun1,
		ActionWRun2,
		ActionNWRun1,
		ActionNWRun2,
	}
)

type dir string

const (
	dirW = "W"
	dirNW = "NW"
	dirN = "N"
	dirNE = "NE"
	dirE = "E"
	dirSW = "SW"
	dirS = "S"
	dirSE = "SE"
)

var (
	DefaultBehavior = NekoBehavior{
		Step: 15,
		Dmax: 20,
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

// runAction returns action name given the direction (see direction function) and whether the action is even.
//
// Returns values are hardcoded because go tool gives better coverage reports this way.
func runAction(d dir, even bool) Action {
	switch d {
	case dirN:
		if even {
			return ActionNRun2
		} else {
			return ActionNRun1
		}
	case dirNE:
		if even {
			return ActionNERun2
		} else {
			return ActionNERun1
		}
	case dirE:
		if even {
			return ActionERun2
		} else {
			return ActionERun1
		}
	case dirSE:
		if even {
			return ActionSERun2
		} else {
			return ActionSERun1
		}
	case dirS:
		if even {
			return ActionSRun2
		} else {
			return ActionSRun1
		}
	case dirSW:
		if even {
			return ActionSWRun2
		} else {
			return ActionSWRun1
		}
	case dirW:
		if even {
			return ActionWRun2
		} else {
			return ActionWRun1
		}
	case dirNW:
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
func direction(x, y, mx, my float64) dir {
	dx := x - mx
	dy := y - my
	α := math.Atan2(dy, dx)
	switch {
	case -π * 8/8 <= α && α <= -π * 7/8:
		return dirE
	case -π * 7/8 <= α && α <= -π * 5/8:
		return dirSE
	case -π * 5/8 <= α && α <= -π * 3/8:
		return dirS
	case -π * 3/8 <= α && α <= -π * 1/8:
		return dirSW
	case -π * 1/8 <= α && α <= +π * 1/8:
		return dirW
	case +π * 1/8 <= α && α <= +π * 3/8:
		return dirNW
	case +π * 3/8 <= α && α <= +π * 5/8:
		return dirN
	case +π * 5/8 <= α && α <= +π * 7/8:
		return dirNE
	case +π * 7/8 <= α && α <= +π * 8/8:
		return dirE
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

func NewInitialState() ActionState {
	return initialState{}
}

type initialState struct{}

func (initialState) Next(n NekoState, m MouseState, b NekoBehavior) ActionState {
	return stateStill{}
}

func (initialState) Render(n NekoState, m MouseState, b NekoBehavior) NekoState {
	return n
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
	d := direction(n.X, n.Y, m.X, m.Y)
	n.Action = runAction(d, s.even)
	makeStep(&n, m, b)
	return n
}
