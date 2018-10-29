package dummyneko

import (
	"testing"
)

func TestDirection(t *testing.T) {
	cases := []struct{
		x, y float64
		mx, my float64
		dir string
	}{
		{
			0, 0,
			0, 1,
			"s",
		},
		{
			0, 0,
			1, 1,
			"se",
		},
		{ // exactly E
			0, 0,
			1, 0,
			"e",
		},
		{ // E -π
			0, 0,
			1, +.1,
			"e",
		},
		{ // E +π
			0, 0,
			1, -.1,
			"e",
		},
		{
			0, 0,
			1, -1,
			"ne",
		},
		{
			0, 0,
			0, -1,
			"n",
		},
		{
			0, 0,
			-1, -1,
			"nw",
		},
		{
			0, 0,
			-1, 0,
			"w",
		},
		{
			0, 0,
			-1, 1,
			"sw",
		},
	}

	for _, c := range cases {
		dir := direction(c.x, c.y, c.mx, c.my)
		if dir != c.dir {
			t.Errorf(
				"Direction(%f, %f, %f, %f) expected %q, got %q",
				c.x, c.y, c.mx, c.my,
				c.dir, dir,
			)
		}
	}
}

func TestStatesChain(t *testing.T) {
	states := []struct{
		e NekoState
		m MouseState
		b NekoBehavior
	}{
		{
			e: NekoState{Action:ActionStill},
			b: NekoBehavior{
				StillTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionStill},
			b: NekoBehavior{
				StillTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionYawn},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionYawn},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionStill},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionStill},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionSleep1},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionSleep1},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionSleep2},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionSleep2},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 2,
			},
		},
		{
			e: NekoState{Action:ActionSleep1},
			b: NekoBehavior{
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
			},
		},
		{
			e: NekoState{Action:ActionSleep2},
			m: MouseState{X:1,Y:1},
			b: NekoBehavior{
				Dmax: 2,
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
			},
		},
		{
			e: NekoState{Action:ActionAlert},
			m: MouseState{X:1,Y:1},
			b: NekoBehavior{
				Dmax: 1,
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:3,Y:4,Action:ActionSERun1},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionSERun2},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 2,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionStill},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 2,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionYawn},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionStill},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionStill},
			m: MouseState{X:6,Y:8},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:6,Y:8,Action:ActionAlert},
			m: MouseState{X:1,Y:-4},
			b: NekoBehavior{
				Step: 13,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionNWRun1},
			m: MouseState{X:-4,Y:-16},
			b: NekoBehavior{
				Step: 13,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionStill},
			m: MouseState{X:1,Y:-4},
			b: NekoBehavior{
				Step: 13,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionAlert},
			m: MouseState{X:1920,Y:1080},
			b: NekoBehavior{
				Step: 20,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionStill},
			m: MouseState{X:1,Y:-4},
			b: NekoBehavior{
				Step: 20,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionYawn},
			m: MouseState{X:1,Y:-4},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 2,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionAlert},
			m: MouseState{X:4,Y:0},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 2,
			},
		},
		{
			e: NekoState{X:1,Y:-4,Action:ActionAlert},
			m: MouseState{X:4,Y:0},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 2,
			},
		},
		{
			e: NekoState{X:4,Y:0,Action:ActionSERun1},
			m: MouseState{X:4,Y:0},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
		{
			e: NekoState{X:4,Y:0,Action:ActionStill},
			m: MouseState{X:4,Y:0},
			b: NekoBehavior{
				Step: 5,
				Dmax: 1,
				StillTicks: 1,
				YawnTicks: 1,
				YawnStillTicks: 2,
				SleepTicks: 1,
				AlertTicks: 1,
			},
		},
	}

	var n NekoState
	var s ActionState
	for _, c := range states {
		if s != nil {
			s = s.Next(n, c.m, c.b)
			if s == nil {
				t.Fatal("Next returned nil state")
			}
		} else {
			// initial state: check Render return value
			s = stateStill{}
		}

		n = s.Render(n, c.m, c.b)
		if n != c.e {
			t.Errorf("expected %#v, got %#v", c.e, n)
		}
	}
}
