package dummyneko

import (
	"testing"
)

func TestDirection(t *testing.T) {
	cases := []struct {
		x, y   float64
		mx, my float64
		d      dir
	}{
		{
			0, 0,
			0, 1,
			dirS,
		},
		{
			0, 0,
			1, 1,
			dirSE,
		},
		{ // exactly E
			0, 0,
			1, 0,
			dirE,
		},
		{ // E -π
			0, 0,
			1, +.1,
			dirE,
		},
		{ // E +π
			0, 0,
			1, -.1,
			dirE,
		},
		{
			0, 0,
			1, -1,
			dirNE,
		},
		{
			0, 0,
			0, -1,
			dirN,
		},
		{
			0, 0,
			-1, -1,
			dirNW,
		},
		{
			0, 0,
			-1, 0,
			dirW,
		},
		{
			0, 0,
			-1, 1,
			dirSW,
		},
	}

	for _, c := range cases {
		d := direction(c.x, c.y, c.mx, c.my)
		if d != c.d {
			t.Errorf(
				"Direction(%f, %f, %f, %f) expected %q, got %q",
				c.x, c.y, c.mx, c.my,
				c.d, d,
			)
		}
	}
}

func TestMajorDirection(t *testing.T) {
	cases := []struct {
		x, y   float64
		mx, my float64
		d      dir
	}{
		{
			0, 0,
			0, 1,
			dirS,
		},
		{
			0, 0,
			1.1, 1,
			dirE,
		},
		{
			0, 0,
			1, 1.1,
			dirS,
		},
		{ // exactly E
			0, 0,
			1, 0,
			dirE,
		},
		{ // E -π
			0, 0,
			1, +.1,
			dirE,
		},
		{ // E +π
			0, 0,
			1, -.1,
			dirE,
		},
		{
			0, 0,
			1.1, -1,
			dirE,
		},
		{
			0, 0,
			1, -1.1,
			dirN,
		},
		{
			0, 0,
			0, -1,
			dirN,
		},
		{
			0, 0,
			-1.1, -1,
			dirW,
		},
		{
			0, 0,
			-1, -1.1,
			dirN,
		},
		{
			0, 0,
			-1, 0,
			dirW,
		},
		{
			0, 0,
			-1.1, 1,
			dirW,
		},
		{
			0, 0,
			-1, 1.1,
			dirS,
		},
	}

	for _, c := range cases {
		d := majorDirection(c.x, c.y, c.mx, c.my)
		if d != c.d {
			t.Errorf(
				"Direction(%f, %f, %f, %f) expected %q, got %q",
				c.x, c.y, c.mx, c.my,
				c.d, d,
			)
		}
	}
}

func TestStatesChain(t *testing.T) {
	states := []struct {
		e State
		m Pos
		b Options
	}{
		{
			e: State{Action: ActionStill},
			b: Options{
				StillTicks: 2,
			},
		},
		{
			e: State{Action: ActionStill},
			b: Options{
				StillTicks: 2,
			},
		},
		{
			e: State{Action: ActionYawn},
			b: Options{
				StillTicks: 2,
				YawnTicks:  2,
			},
		},
		{
			e: State{Action: ActionYawn},
			b: Options{
				StillTicks: 2,
				YawnTicks:  2,
			},
		},
		{
			e: State{Action: ActionStill},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
			},
		},
		{
			e: State{Action: ActionStill},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
			},
		},
		{
			e: State{Action: ActionSleep1},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    2,
			},
		},
		{
			e: State{Action: ActionSleep1},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    2,
			},
		},
		{
			e: State{Action: ActionSleep2},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    2,
			},
		},
		{
			e: State{Action: ActionSleep2},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    2,
			},
		},
		{
			e: State{Action: ActionSleep1},
			b: Options{
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
			},
		},
		{
			e: State{Action: ActionSleep2},
			m: Pos{X: 1, Y: 1},
			b: Options{
				Dmax:          2,
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
			},
		},
		{
			e: State{Action: ActionAlert},
			m: Pos{X: 1, Y: 1},
			b: Options{
				Dmax:          1,
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 3, Y: 4, Action: ActionSERun1},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionSERun2},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    2,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionStill},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     2,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionYawn},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionStill},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionStill},
			m: Pos{X: 6, Y: 8},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 6, Y: 8, Action: ActionAlert},
			m: Pos{X: 1, Y: -4},
			b: Options{
				Step:          13,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionNWRun1},
			m: Pos{X: -4, Y: -16},
			b: Options{
				Step:          13,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionStill},
			m: Pos{X: 1, Y: -4},
			b: Options{
				Step:          13,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionAlert},
			m: Pos{X: 1920, Y: 1080},
			b: Options{
				Step:          20,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionStill},
			m: Pos{X: 1, Y: -4},
			b: Options{
				Step:          20,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionYawn},
			m: Pos{X: 1, Y: -4},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    2,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionAlert},
			m: Pos{X: 4, Y: 0},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    2,
			},
		},
		{
			e: State{X: 1, Y: -4, Action: ActionAlert},
			m: Pos{X: 4, Y: 0},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    2,
			},
		},
		{
			e: State{X: 4, Y: 0, Action: ActionSERun1},
			m: Pos{X: 4, Y: 0},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 4, Y: 0, Action: ActionStill},
			m: Pos{X: 4, Y: 0},
			b: Options{
				Step:          5,
				Dmax:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 4, Y: 0, Action: ActionAlert},
			m: Pos{X: 0, Y: 0},
			b: Options{
				Step:          1,
				StillTicks:    1,
				YawnTicks:     1,
				PostYawnTicks: 2,
				SleepTicks:    1,
				AlertTicks:    1,
			},
		},
		{
			e: State{X: 3, Y: 0, Action: ActionWRun1},
			m: Pos{X: 0, Y: 0},
			b: Options{Step: 1},
		},
		{
			e: State{X: 2, Y: 0, Action: ActionWRun2},
			m: Pos{X: 0, Y: 0},
			b: Options{Step: 1},
		},
		{
			e: State{X: 1, Y: 0, Action: ActionWRun1},
			m: Pos{X: 0, Y: 0},
			b: Options{Step: 1},
		},
		{
			e: State{X: 0, Y: 0, Action: ActionWRun2},
			m: Pos{X: 0, Y: 0},
			b: Options{Step: 1},
		},
		{
			e: State{X: -3, Y: -4, Action: ActionNWRun1},
			m: Pos{X: -6, Y: -8},
			b: Options{Step: 5},
		},
		{
			e: State{X: -6, Y: -8, Action: ActionNWRun2},
			m: Pos{X: -6, Y: -8},
			b: Options{Step: 5},
		},
	}

	var n State
	s := NewInitialState()
	for _, c := range states {
		s = s.Next(n, c.m, c.b)
		if s == nil {
			t.Fatal("Next returned nil state")
		}

		n = s.Render(n, c.m, c.b)
		if n != c.e {
			t.Errorf("expected %#v, got %#v", c.e, n)
		}
	}
}
