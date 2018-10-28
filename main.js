'use strict'

// time between state transition, ms
const dt = 300
// max distance between cursor and cat
const dmax = 15
// speed per `dt` ms
const step = 15

/* Current mouse position */
var mx = 0
var my = 0

document.addEventListener('mousemove', mouseUpdate, false)
document.addEventListener('mouseenter', mouseUpdate, false)
function mouseUpdate(e) {
  mx = e.pageX
  my = e.pageY
}

// - utils
// you can change how element changes its position (`commit` function),
//                            displays image for some state (`display_state`),
// change `make_step` function or define new state display names.
// `states` table exists only as a reference for designers but you may use it
// for your needs (e.g., return default value for undefined state names in
// `display_name` function).

// List of all possible state display names
const states = [ "alert"
               , "still"
               , "yawn"
               , "itch1"    , "itch2"
               , "sleep1"   , "sleep2"
               // run
               , "nrun1"    , "nrun2"
               , "nerun1"   , "nerun2"
               , "erun1"    , "erun2"
               , "serun1"   , "serun2"
               , "srun1"    , "srun2"
               , "swrun1"   , "swrun2"
               , "wrun1"    , "wrun2"
               , "nwrun1"   , "nwrun2"
               // scratch
               , "nscratch1", "nscratch2"
               , "escratch1", "escratch2"
               , "sscratch1", "sscratch2"
               , "wscratch1", "wscratch2"
               ]

// name - state name
// dir  - direction (optional)
// n    - state iteration (optional)
let display_name = function(name, n, dir) {
  return (dir ? dir : '') + name + (n ? n : '')
}

let display_state = function(e, name) {
  e.src = '/ass/webneko.net/socks/' + name + '.gif'
}

// commit changes in element's position
let commit = function(e, x, y) {
  e.style.top  = y + 'px'
  e.style.left = x + 'px'
}

// get one step closer to the pointer
let make_step = function(p) {
  let d = distance(p.x, p.y)
  if(d > 0) {
    let dx = p.x - mx
    let dy = p.y - my
    let dstep = step / d

    p.x -= dstep * dx
    p.y -= dstep * dy
  }
}

// state transition with delay
let transit = function(state, nrepeat) {
  if(nrepeat == undefined) nrepeat = 1
  setTimeout(state, dt * nrepeat)
}

// - general math
// most likely you do not have to change anything in this section.

const π = Math.PI
// If you want to want to have the major directions only:
//   angleMap.filter((_, i) => !(i&1))
const angleMap = [ [     0, 'w' ]
                 , [ π*1/4, 'nw']
                 , [ π / 2, 'n' ]
                 , [ π*3/4, 'ne']
                 , [     π, 'e' ]
                 , [-π*1/4, 'sw']
                 , [-π / 2, 's' ]
                 , [-π*3/4, 'se']
                 , [    -π, 'e' ]
                 ]


/* * * * * * */
/* NW  ↑  NE */
/*     N     */
/* ← W ☼ E → */
/*     S     */
/* SW  ↓  SE */
/* * * * * * */
function direction(x, y, dirs) {
  if(dirs == undefined) dirs = angleMap

  let dx = x - mx
  let dy = y - my
  let α = Math.atan2(dy, dx)

  for (let element of dirs) {
   let key = element[0], val = element[1]
   let β = α - key
   if (-π/8 <= β && β <= π/8) {
     return val
   }
  }

  return ''
}

function distance(x, y) {
  let dx = x - mx
  let dy = y - my
  return Math.sqrt(dx*dx + dy*dy)
}

function pointerNearby(p) {
  return distance(p.x, p.y) <= dmax
}

// - states

function state_still(p, e, n) {
  if(n == undefined) n = 1
  let next = (n + 1) & 0xFF

  if(!pointerNearby(p)) {
    return state_alert(p, e)
  }

  if(15 <= n && n <= 18) {
    return state_yawn(p, e, next)
  }
  if(n > 20) {
    return state_sleep(p, e)
  }

  let state = display_name('still')
  display_state(e, state)

  return transit(
    function() { return state_still(p, e, next) }
  )
}

// here `n` is iteration of still state, not yawn state
function state_yawn(p, e, n) {
  if(!pointerNearby(p)) {
    return state_still(p, e)
  }

  let state = display_name('yawn')
  display_state(e, state)

  return transit(
    function() { return state_still(p, e, n) }
  )
}

function state_sleep(p, e, n) {
  if(n != 1 && n != 2) n = 1

  // wake up!
  if(!pointerNearby(p)) {
    return state_still(p, e)
  }

  let state = display_name('sleep', n)
  display_state(e, state)

  return transit(
    function() { return state_sleep(p, e, (n&1)+1) },
    2
  )
}

function state_alert(p, e) {
  let state = display_name('alert')
  display_state(e, state)

  return transit(
    function() { return state_run(p, e) },
    2 // * -> alert -> alert -> run
  )
}

function state_run(p, e, n) {
  if(n != 1 && n != 2) n = 1

  if(pointerNearby(p)) {
    return state_still(p, e)
  }

  make_step(p)
  commit(e, p.x, p.y)
  let state = display_name('run', n, direction(p.x, p.y))
  display_state(e, state)

  return transit(
    function() { return state_run(p, e, (n&1)+1) }
  )
}
