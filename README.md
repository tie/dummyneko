>*Neko* (猫, ねこ) is the Japanese word for *cat*.

# Web/猫

[Neko](https://en.wikipedia.org/wiki/Neko_%28software%29) for the web, somewhat compatible with [WebNeko](https://webneko.net).

## Preview

You can find a cat waiting for mouse on all pages at [b1nary.tk](https://b1nary.tk).
Don't let the neko catch your mouse!

## Credits

- [Daniil Zakirullin](https://github.com/Vftdan) for hacking on ECMA5-compatibility.

## Features. What's working?
- state_still
- state_yawn
- state_sleep
- state_alert
- state_run

## Roadmap. What's not implemented?

- state_scratch
- state_itch
- home position
- refactoring

## Known bugs and workarounds.

- Default `display_state` updates image source URL, and some browsers (e.g.  Chrome) cancel unfinished downloads — low-bandwidth network users never receive the neko (unless they manually preload it).

  Temporary workaround: display preloaded images in `display_state` function.

- Web/猫 was not managed using Git from the start.  Unfortunately, there is no fix or workaround for this problem.

## License

### Public domain

Unlike the [webneko.net](https://webneko.net/n200504.js) JavaScript code, Web/猫 is published and distributed under the Unlicense and WTFPL.  Attribution is optional, but desirable.

Rationales for placing software in public domain are listed in [nothings/stb docs](https://github.com/nothings/stb/blob/master/docs/why_public_domain.md).

### Traditional license

Want a traditional copyright-ish license?

>You are granted a perpetual, irrevocable license to copy, modify, publish, and distribute this software as you see fit.
