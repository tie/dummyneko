package dummyneko

import (
	"math"
)

type NekoState struct {
	X, Y   float64
	Action Action
}

type MouseState struct {
	X, Y float64
}

type Options struct {
	// speed per single action/tick
	Step float64
	// max distance between mouse and neko
	Dmax float64
	// StillTransition is the next state after Still state
	//
	// Values
	//   0: Yawn
	//   1: Itch
	//   2: Scratch
	StillTransition uint

	// ticks per state

	StillTicks               uint
	YawnTicks, PostYawnTicks uint
	SleepTicks               uint
	AlertTicks               uint
	RunTicks                 uint

	ItchTicks, ItchCount, PostItchTicks          uint
	ScratchTicks, ScratchCount, PostScratchTicks uint
	// Disable transition from Scratch to Alert state
	ScratchDisableAlert bool
}

type Action string

const (
	ActionAlert     = "alert"
	ActionStill     = "still"
	ActionYawn      = "yawn"
	ActionItch1     = "itch1"
	ActionItch2     = "itch2"
	ActionSleep1    = "sleep1"
	ActionSleep2    = "sleep2"
	ActionNRun1     = "nrun1"
	ActionNRun2     = "nrun2"
	ActionNERun1    = "nerun1"
	ActionNERun2    = "nerun2"
	ActionERun1     = "erun1"
	ActionERun2     = "erun2"
	ActionSERun1    = "serun1"
	ActionSERun2    = "serun2"
	ActionSRun1     = "srun1"
	ActionSRun2     = "srun2"
	ActionSWRun1    = "swrun1"
	ActionSWRun2    = "swrun2"
	ActionWRun1     = "wrun1"
	ActionWRun2     = "wrun2"
	ActionNWRun1    = "nwrun1"
	ActionNWRun2    = "nwrun2"
	ActionNScratch1 = "nscratch1"
	ActionNScratch2 = "nscratch2"
	ActionEScratch1 = "escratch1"
	ActionEScratch2 = "escratch2"
	ActionSScratch1 = "sscratch1"
	ActionSScratch2 = "sscratch2"
	ActionWScratch1 = "wscratch1"
	ActionWScratch2 = "wscratch2"
)

// SupportedActions is a list of implemented action.
var SupportedActions = []Action{
	ActionAlert,
	ActionStill,
	ActionYawn,
	// itch
	ActionItch1,
	ActionItch2,
	// sleep
	ActionSleep1,
	ActionSleep2,
	// run
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
	// scratch
	ActionNScratch1,
	ActionNScratch2,
	ActionEScratch1,
	ActionEScratch2,
	ActionSScratch1,
	ActionSScratch2,
	ActionWScratch1,
	ActionWScratch2,
}

type dir string

const (
	dirW  = "W"
	dirNW = "NW"
	dirN  = "N"
	dirNE = "NE"
	dirE  = "E"
	dirSW = "SW"
	dirS  = "S"
	dirSE = "SE"
)

var DefaultOptions = Options{
	Step:            15,
	Dmax:            20,
	StillTransition: 1,
	StillTicks:      4,
	SleepTicks:      2,
	AlertTicks:      2,

	YawnTicks:     4,
	PostYawnTicks: 4,

	ItchTicks:     1,
	ItchCount:     6,
	PostItchTicks: 4,

	ScratchTicks:        2,
	ScratchCount:        4,
	PostScratchTicks:    4,
	ScratchDisableAlert: true,
}

const (
	π = math.Pi
)

// runAction returns action name for Run state given the direction (see direction function) and whether the action is even.
func runAction(d dir, even bool) Action {
	// Return values are hardcoded because go tool gives better coverage reports this way.
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

// scratchAction returns action name for Scratch state given the major direction (see majorDirection function) and whether the action is even.
func scratchAction(d dir, even bool) Action {
	switch d {
	case dirN:
		if even {
			return ActionNScratch2
		} else {
			return ActionNScratch1
		}
	case dirE:
		if even {
			return ActionEScratch2
		} else {
			return ActionEScratch1
		}
	case dirS:
		if even {
			return ActionSScratch2
		} else {
			return ActionSScratch1
		}
	case dirW:
		if even {
			return ActionWScratch2
		} else {
			return ActionWScratch1
		}
	}
	return ""
}

// direction function computes the compass direction at the (x,y) to the destination (mx,my) coordinates.
//
// Return values are: E, SE, S, SW, W, NW, N, NE.
//
// Note that it assumes a coordinate system with inverted Y axis, i.e. with origin at the top-left corner.  Invoke it with inverted ordinate sign to use the classic cartesian coordinate system.
func direction(x, y, mx, my float64) dir {
	// Return values are hardcoded because go tool gives better coverage reports this way.
	dx := x - mx
	dy := y - my
	α := math.Atan2(dy, dx)
	switch {
	case -π*8/8 <= α && α <= -π*7/8:
		return dirE
	case -π*7/8 <= α && α <= -π*5/8:
		return dirSE
	case -π*5/8 <= α && α <= -π*3/8:
		return dirS
	case -π*3/8 <= α && α <= -π*1/8:
		return dirSW
	case -π*1/8 <= α && α <= +π*1/8:
		return dirW
	case +π*1/8 <= α && α <= +π*3/8:
		return dirNW
	case +π*3/8 <= α && α <= +π*5/8:
		return dirN
	case +π*5/8 <= α && α <= +π*7/8:
		return dirNE
	case +π*7/8 <= α && α <= +π*8/8:
		return dirE
	}
	return ""
}

// majorDirection function computes the major compass direction at the (x,y) to the destination (mx,my) coordinates.
//
// The major compass directions are: E, S, W, N.
//
// Note that it assumes a coordinate system with inverted Y axis, i.e. with origin at the top-left corner.  Invoke it with inverted ordinate sign to use the classic cartesian coordinate system.
func majorDirection(x, y, mx, my float64) dir {
	// Return values are hardcoded because go tool gives better coverage reports this way.
	dx := x - mx
	dy := y - my
	α := math.Atan2(dy, dx)
	switch {
	case -π*4/4 <= α && α <= -π*3/4:
		return dirE
	case -π*3/4 <= α && α <= -π*1/4:
		return dirS
	case -π*1/4 <= α && α <= +π*1/4:
		return dirW
	case +π*1/4 <= α && α <= +π*3/4:
		return dirN
	case +π*3/4 <= α && α <= +π*4/4:
		return dirE
	}
	return ""
}

func pointerNearby(n NekoState, m MouseState, b Options) bool {
	dx := n.X - m.X
	dy := n.Y - m.Y
	d := math.Hypot(dx, dy)
	return d <= b.Dmax
}

func makeStep(n *NekoState, m MouseState, b Options) {
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
	Next(NekoState, MouseState, Options) ActionState
	Render(NekoState, MouseState, Options) NekoState
}

func NewInitialState() ActionState {
	return initialState{}
}

type initialState struct{}

func (initialState) Next(n NekoState, m MouseState, b Options) ActionState {
	return stateStill{}
}

func (initialState) Render(n NekoState, m MouseState, b Options) NekoState {
	return n
}

type stateStill struct {
	tick uint
}

func (s stateStill) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.StillTicks {
		return s.selectNext(n, m, b)
	}
	return s
}

func (s stateStill) selectNext(n NekoState, m MouseState, b Options) ActionState {
	switch b.StillTransition {
	case 1:
		return stateItch{}
	case 2:
		return stateScratch{}
	}
	return stateYawn{}
}

func (s stateStill) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionStill
	return n
}

type stateItch struct {
	tick  uint
	even  bool
	count uint
}

func (s stateItch) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.ItchTicks {
		s.tick = 0
		s.even = !s.even
		s.count += 1
	}
	if s.count >= b.ItchCount {
		return statePostItch{}
	}
	return s
}

func (s stateItch) Render(n NekoState, m MouseState, b Options) NekoState {
	if s.even {
		n.Action = ActionItch2
	} else {
		n.Action = ActionItch1
	}
	return n
}

type statePostItch struct {
	tick uint
}

func (s statePostItch) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.PostItchTicks {
		return stateYawn{}
	}
	return s
}

func (s statePostItch) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionStill
	return n
}

type stateScratch struct {
	tick  uint
	even  bool
	count uint
}

func (s stateScratch) Next(n NekoState, m MouseState, b Options) ActionState {
	if !b.ScratchDisableAlert && !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.ScratchTicks {
		s.tick = 0
		s.even = !s.even
		s.count += 1
	}
	if s.count >= b.ScratchCount {
		return statePostScratch{}
	}
	return s
}

func (s stateScratch) Render(n NekoState, m MouseState, b Options) NekoState {
	d := majorDirection(n.X, n.Y, m.X, m.Y)
	n.Action = scratchAction(d, s.even)
	return n
}

type statePostScratch struct {
	tick uint
}

func (s statePostScratch) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.PostScratchTicks {
		return stateYawn{}
	}
	return s
}

func (s statePostScratch) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionStill
	return n
}

type stateYawn struct {
	tick uint
}

func (s stateYawn) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.YawnTicks {
		return statePostYawn{}
	}
	return s
}

func (s stateYawn) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionYawn
	return n
}

type statePostYawn struct {
	tick uint
}

func (s statePostYawn) Next(n NekoState, m MouseState, b Options) ActionState {
	if !pointerNearby(n, m, b) {
		return stateAlert{}
	}
	s.tick += 1
	if s.tick >= b.PostYawnTicks {
		return stateSleep{}
	}
	return s
}

func (s statePostYawn) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionStill
	return n
}

type stateSleep struct {
	tick uint
	even bool
}

func (s stateSleep) Next(n NekoState, m MouseState, b Options) ActionState {
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

func (s stateSleep) Render(n NekoState, m MouseState, b Options) NekoState {
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

func (s stateAlert) Next(n NekoState, m MouseState, b Options) ActionState {
	if pointerNearby(n, m, b) {
		return stateStill{}
	}
	s.tick += 1
	if s.tick >= b.AlertTicks {
		return stateRun{}
	}
	return s
}

func (s stateAlert) Render(n NekoState, m MouseState, b Options) NekoState {
	n.Action = ActionAlert
	return n
}

type stateRun struct {
	tick uint
	even bool
}

func (s stateRun) Next(n NekoState, m MouseState, b Options) ActionState {
	if pointerNearby(n, m, b) {
		return stateStill{}
	}
	s.tick += 1
	if s.tick >= b.RunTicks {
		s.tick = 0
		s.even = !s.even
	}
	return s
}

func (s stateRun) Render(n NekoState, m MouseState, b Options) NekoState {
	d := direction(n.X, n.Y, m.X, m.Y)
	n.Action = runAction(d, s.even)
	makeStep(&n, m, b)
	return n
}
