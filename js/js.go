package main

import (
	"strconv"
	"time"

	"github.com/gopherjs/gopherwasm/js"

	neko "github.com/tie/dummyneko"
)

func main() {
	n, m, b := neko.State{}, neko.Pos{}, neko.DefaultOptions

	mouseUpdate := js.NewEventCallback(0, func(ev js.Value) {
		m.X, m.Y = ev.Get("clientX").Float(), ev.Get("clientY").Float()
	})

	global := js.Global()

	go func() {
		image := global.Get("Image")
		for _, a := range neko.SupportedActions {
			img := image.New()
			img.Set("src", imgUrl(a))
		}
	}()

	doc := global.Get("document")

	doc.Call("addEventListener", "mousemove", mouseUpdate, false)
	doc.Call("addEventListener", "mouseenter", mouseUpdate, false)

	global.Get("window").Call("addEventListener", "load", js.NewEventCallback(0, func(js.Value) {
		e := doc.Call("createElement", "img")
		setupElement(e)
		doc.Get("body").Call("appendChild", e)
		s := neko.NewInitialState()
		ticker := time.NewTicker(300 * time.Millisecond)
		for {
			s = s.Next(n, m, b)
			n = s.Render(n, m, b)
			displayState(e, n)

			select {
			case <-ticker.C:
				switch b.StillTransition {
				case 0:
					b.StillTransition = 2
				case 1:
					b.StillTransition = 0
				case 2:
					b.StillTransition = 1
				}
				continue
			}
		}
	}))
}

func setupElement(e js.Value) {
	e.Set("draggable", false)
	styles := e.Get("style")
	styles.Set("position", "fixed")
	styles.Set("width", "32px")
	styles.Set("top", "0px")
	styles.Set("left", "0px")
	styles.Set("imageRendering", "pixelated")
}

func displayState(e js.Value, n neko.State) {
	style := e.Get("style")
	style.Set("left", f2px(n.X))
	style.Set("top", f2px(n.Y))
	e.Set("src", imgUrl(n.Action))
}

func imgUrl(a neko.Action) neko.Action {
	return "https://b1nary.tk/ass/webneko.net/socks/" + a + ".gif"
}

func f2px(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64) + "px"
}
