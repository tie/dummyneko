package main

import (
	"fmt"
	"time"
	"github.com/gopherjs/gopherwasm/js"

	neko "github.com/tie/dummyneko"
)

func main() {
	// neko.DefaultBehavior
	n, m, b := neko.NekoState{}, neko.MouseState{}, neko.DefaultBehavior

	global := js.Global()

	doc := global.Get("document")
	mouseUpdate := js.NewEventCallback(0, func(ev js.Value) {
		m.X, m.Y = ev.Get("pageX").Float(), ev.Get("pageY").Float()
	})

	doc.Call("addEventListener", "mousemove", mouseUpdate, false)
	doc.Call("addEventListener", "mouseenter", mouseUpdate, false)

	global.Get("window").Call("addEventListener", "load", js.NewEventCallback(0, func(js.Value) {
		e := doc.Call("createElement", "img")
		styles := e.Get("style")
		styles.Set("position", "absolute")
		styles.Set("width", "32px")
		styles.Set("top", "0px")
		styles.Set("left", "0px")
		styles.Set("imageRendering", "pixelated")
		styles.Set("draggable", false)

		doc.Get("body").Call("appendChild", e)

		s := neko.InitialState
		n = s.Render(n, m, b)
		displayState(e, n)

		ticker := time.NewTicker(300 * time.Millisecond)
		for {
			select {
			case <-ticker.C:
				s = s.Next(n, m, b)
				n = s.Render(n, m, b)
				displayState(e, n)
			}
		}
	}))
}

func displayState(e js.Value, n neko.NekoState) {
	style := e.Get("style")
	style.Set("left", fmt.Sprintf("%fpx", n.X))
	style.Set("top", fmt.Sprintf("%fpx", n.Y))
	e.Set("src", fmt.Sprintf(
		"https://b1nary.tk/ass/webneko.net/socks/%s.gif",
		n.Action,
	))
}
