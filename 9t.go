package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
)

var ( // MARK: var ███████████████████████████████

	// portal
	hintson                                        bool
	portalh, portalv, portalblock, portaldirection int
	// weapons
	destroyon                       bool
	destroyx, destroyy              int
	destroytimer, destroytimercount int
	bullets                         = make([]bullet, worlda)
	// intro
	startnewgameon bool
	scrolltextx    = -200
	scrolltext2x   = monw + 200
	introtextfade  = 0.0
	backfade       = 1.0
	backfade2      = 0.0
	dinomultiplier = make([]int, 500)
	dinotype       = make([]int, 500)
	//game
	introon, levelon, levelend                            bool
	treetype, gametime, gametimecount, score, levelnumber int
	//fx
	emotecount                      int
	emotex                          float32
	cloudtimer, cloudtimercount     int
	cloudspeed                      int
	cloudlr, cloudson               bool
	cloudsv2                        = make([]rl.Vector2, 100)
	cloudsl                         = make([]int, 100)
	weathertimer, weathertimercount int
	lineson, weather, rain, snow    bool
	snowimg                         = make([]rl.Rectangle, 200)
	snowv2                          = make([]rl.Vector2, 200)
	snowl                           = make([]int, 200)
	rainv2                          = make([]rl.Vector2, 200)
	rainl                           = make([]float32, 200)
	//enemies & monsters
	monsternumber = 500
	enemiesnumber = 1000
	enemies       = make([]enemyblok, worlda)
	monsters      = make([]monsterblok, worlda)
	//options
	optionselect int
	//powerups
	flyingrotation, axerotation, bombrotation                         float32
	pickuptexttimer, pickuptexttimercount, pickuptextychange          int
	pickuptexton                                                      bool
	currentpickuptext                                                 string
	flyingobjectblok, flyingh, flyingv, flyingy, flyingx, flyingspeed int
	flyingtimer, flyingtimercount                                     int
	flyingon, flyingdrop                                              bool
	//platforms
	plattile1, plattile2, plattile3    rl.Rectangle
	plat1color, plat2color, plat3color rl.Color
	//backgrounds
	backgzoomon            bool
	backcolor1, backcolor2 rl.Color
	backtiles              = make([]rl.Rectangle, drawa)
	backtile1, backtile2   rl.Rectangle
	backgon                bool
	background             = make([]backg, 10)
	//player
	currentemote                rl.Rectangle
	emoteon                     bool
	emotetimer, emotetimercount int
	coinstotal                  int
	jumph                       = 5
	jumpon, fallon              bool
	player                      = playerblok{}
	// map
	tilesize, textsize1, textsize2 int
	freemoveon                     = true
	screenw                        = 80 // 16px blocks - 1 screen 1280
	screenh                        = 45 // 16px blocks - 1 screen 720
	worldw                         = screenw * 12
	worldh                         = screenh * 4
	worlda                         = worldw * worldh
	world                          = make([]blok, worlda)
	objects                        = make([]blok, worlda)
	//img
	portal         = rl.NewRectangle(126, 358, 48, 48)
	monster1       = rl.NewRectangle(383, 917, 16, 16)
	monster2       = rl.NewRectangle(463, 917, 16, 16)
	monster3       = rl.NewRectangle(523, 917, 16, 16)
	monster4       = rl.NewRectangle(382, 937, 16, 16)
	monster5       = rl.NewRectangle(463, 937, 16, 16)
	monster6       = rl.NewRectangle(382, 957, 16, 16)
	monster7       = rl.NewRectangle(462, 957, 16, 16)
	monster8       = rl.NewRectangle(382, 977, 16, 16)
	monster9       = rl.NewRectangle(460, 977, 16, 16)
	monster10      = rl.NewRectangle(382, 997, 16, 16)
	monster11      = rl.NewRectangle(468, 997, 16, 16)
	monster12      = rl.NewRectangle(382, 1017, 16, 16)
	monster13      = rl.NewRectangle(502, 1017, 16, 16)
	monster14      = rl.NewRectangle(384, 1037, 16, 16)
	monster15      = rl.NewRectangle(502, 1037, 16, 16)
	monster16      = rl.NewRectangle(383, 1057, 16, 16)
	monster17      = rl.NewRectangle(483, 1057, 16, 16)
	monster18      = rl.NewRectangle(383, 1077, 16, 16)
	monster19      = rl.NewRectangle(501, 1077, 16, 16)
	monster20      = rl.NewRectangle(382, 1097, 16, 16)
	monster21      = rl.NewRectangle(501, 1097, 16, 16)
	monster22      = rl.NewRectangle(382, 1117, 16, 16)
	monster23      = rl.NewRectangle(382, 1137, 16, 16)
	monster24      = rl.NewRectangle(481, 1137, 16, 16)
	monster25      = rl.NewRectangle(380, 1157, 16, 16)
	propellor2     = rl.NewRectangle(1328, 603, 31, 31)
	propellor      = rl.NewRectangle(1301, 575, 228, 8)
	helicopterr    = rl.NewRectangle(1295, 328, 237, 112)
	helicopterl    = rl.NewRectangle(1296, 452, 237, 112)
	timeimg        = rl.NewRectangle(623, 64, 15, 15)
	snowflake1     = rl.NewRectangle(574, 41, 16, 16)
	snowflake2     = rl.NewRectangle(590, 41, 16, 16)
	snowflake3     = rl.NewRectangle(606, 41, 16, 16)
	snowflake4     = rl.NewRectangle(622, 41, 16, 16)
	raddishimg     = rl.NewRectangle(9, 1032, 30, 38)
	raddishlimg    = rl.NewRectangle(160, 984, 30, 38)
	spikesimg      = rl.NewRectangle(13, 940, 44, 26)
	spikeslimg     = rl.NewRectangle(321, 896, 44, 26)
	mushroomimg    = rl.NewRectangle(5, 840, 32, 32)
	mushroomlimg   = rl.NewRectangle(488, 794, 32, 32)
	ghostimg       = rl.NewRectangle(6, 748, 44, 30)
	ghostlimg      = rl.NewRectangle(403, 709, 44, 30)
	bunnyimg       = rl.NewRectangle(2, 552, 34, 44)
	bunnylimg      = rl.NewRectangle(378, 492, 34, 44)
	chickenimg     = rl.NewRectangle(3, 660, 32, 34)
	chickenlimg    = rl.NewRectangle(422, 616, 32, 34)
	springimg      = rl.NewRectangle(590, 19, 16, 16)
	scorpionimg    = rl.NewRectangle(520, 3, 16, 16)
	bombimg        = rl.NewRectangle(542, 2, 16, 16)
	bomblimg       = rl.NewRectangle(494, 3, 16, 16)
	shotgunimg     = rl.NewRectangle(564, 4, 16, 16)
	shotgunlimg    = rl.NewRectangle(472, 5, 16, 16)
	uzziimg        = rl.NewRectangle(586, 3, 16, 16)
	uzzilimg       = rl.NewRectangle(450, 4, 16, 16)
	bazookaimg     = rl.NewRectangle(604, 2, 16, 16)
	bazookalimg    = rl.NewRectangle(432, 3, 16, 16)
	axeimg         = rl.NewRectangle(624, 0, 16, 16)
	axelimg        = rl.NewRectangle(412, 1, 16, 16)
	hpimg          = rl.NewRectangle(0, 367, 42, 36)
	coinimg        = rl.NewRectangle(0, 343, 16, 16)
	powerupblokimg = rl.NewRectangle(0, 327, 16, 16)
	emotes         = make([]rl.Rectangle, 110)
	trees          = make([]rl.Rectangle, 15)
	trees2         = make([]rl.Rectangle, 15)
	trees3         = make([]rl.Rectangle, 20)
	tiles          = make([]rl.Rectangle, 40)
	dinor          = rl.NewRectangle(0, 456, 24, 24)
	dinol          = rl.NewRectangle(314, 431, 24, 24)
	dino2r         = rl.NewRectangle(846, 797, 24, 24)
	dino2l         = rl.NewRectangle(1180, 765, 24, 24)
	dino3r         = rl.NewRectangle(1216, 728, 24, 24)
	dino3l         = rl.NewRectangle(1552, 692, 24, 24)
	dino4r         = rl.NewRectangle(1227, 800, 24, 24)
	dino4l         = rl.NewRectangle(1560, 765, 24, 24)
	// core
	options, paused, scanlines, pixelnoise, ghosting           bool
	centerblok, drawblok, drawbloknext, draww, drawh, drawa    int
	drawwextend, drawhextend, drawaextend, drawblocknextextend int
	mouseblok                                                  int
	mousepos                                                   rl.Vector2
	gridon, debugon, fadeblinkon, fadeblink2on                 bool
	monw, monh                                                 int
	fps                                                        = 30
	framecount                                                 int
	imgs                                                       rl.Texture2D
	camera, camerabackg                                        rl.Camera2D
	fadeblink                                                  = float32(0.2)
	fadeblink2                                                 = float32(0.1)
	onoff2, onoff3, onoff6, onoff10, onoff15, onoff30, onoff60 bool
)

// MARK: struct
type bullet struct {
	name                    string
	activ, direction, pause bool
	hp, nextblock           int
}
type monsterblok struct {
	name                    string
	activ, switch1, switch2 bool
	hp                      int
	img                     rl.Rectangle
}
type enemyblok struct {
	name                                                                        string
	pause, activ, lr, switch1, switch2                                          bool
	x, y, pausecount, hp, count1, count2, nextblock, direction, spikesdirection int
	img                                                                         rl.Rectangle
}
type backg struct {
	v2             rl.Vector2
	w, h           int
	color1, color2 rl.Color
	opac           float32
}
type playerblok struct {
	holding                               string
	hp, bloknumber, h, v, x, y            int
	isholding, direction, moving, jumping bool
	objimg, objlimg                       rl.Rectangle
}
type blok struct {
	portal, pickup, coin, switch1, switch2, powerup, activ, imgon, solid, ground bool
	rotation                                                                     float32
	x, y, hp, nextblock                                                          int
	img                                                                          rl.Rectangle
	color1, color2                                                               rl.Color
}

func raylib() { // MARK: raylib
	rl.InitWindow(monw, monh, "GAME TITLE")
	//rl.ToggleFullscreen()
	rl.SetExitKey(rl.KeyEnd) // key to end the game and close window
	// MARK: load images
	imgs = rl.LoadTexture("imgs.png") // load images
	createimgs()
	createmap()
	rl.SetTargetFPS(fps)
	//rl.HideCursor()
	//	rl.ToggleFullscreen()
	for !rl.WindowShouldClose() {
		framecount++
		mousepos = rl.GetMousePosition()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		drawnocameraback()
		rl.BeginMode2D(camera)
		if !paused {
			drawlayers()
			if destroyon {
				drawdestroy()
			}
			if flyingon {
				drawflying()
			}
		}
		if gridon {
			drawgrid()
		}

		rl.EndMode2D()
		drawnocamera()

		if debugon {
			drawdebug()
		}
		update()
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

func update() { // MARK: update

	input()
	if !paused {
		timers()
		updateplayer()
		updatebojects()
		updateenemies()
		updatecollisions()
		updatehv()
	}

	if startnewgameon {
		if backfade2 < 1.0 {
			backfade2 += 0.02
		} else {
			introon = false
		}
	}

	if introon {

		if backfade > 0 {
			backfade -= 0.02
		} else {
			if introtextfade < 1.0 {
				introtextfade += 0.02
			}
		}
		if introtextfade >= 1.0 {
			scrolltext2x -= 5
			scrolltextx += 5

			if scrolltext2x < -200 {
				scrolltext2x = monw + 200
			}
			if scrolltextx > monw+200 {
				scrolltextx = -200
			}
		}

	}
	if levelend {
		paused = true
		drawlevelend()

	}

}
func updatecollisions() { // MARK: updatecollisions

	count := 0
	checkblock := drawblocknextextend
	for a := 0; a < drawaextend; a++ {

		if bullets[checkblock].activ && enemies[checkblock].activ {
			enemies[checkblock].hp--
		}

		checkblock++
		count++
		if count == drawwextend {
			count = 0
			checkblock -= drawwextend
			checkblock += worldw
		}
	}

}
func updateenemies() { // MARK: updateenemies

	count := 0
	checkblock := drawblocknextextend
	for a := 0; a < drawaextend; a++ {
		if enemies[checkblock].activ {

			if enemies[checkblock].pause {
				enemies[checkblock].pausecount++
			}
			if enemies[checkblock].pausecount >= 6 {
				enemies[checkblock].pausecount = 0
				enemies[checkblock].pause = false
			}

			enemies[checkblock].count1++
			switch enemies[checkblock].name {
			case "mushroom":
				if framecount%8 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			case "spikes":
				if framecount%6 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			case "raddish":
				if framecount%5 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			case "bunny":
				if framecount%3 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			case "chicken":
				if framecount%12 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			case "ghost":
				if framecount%4 == 0 {
					moveenemy(enemies[checkblock].direction, checkblock)
				}
			}
		}

		if monsters[checkblock].activ {
			if !world[checkblock-worldw].solid {
				monsters[checkblock-worldw] = monsters[checkblock]
				monsters[checkblock] = monsterblok{}
			}

		}

		checkblock++
		count++
		if count == drawwextend {
			count = 0
			checkblock -= drawwextend
			checkblock += worldw
		}
	}

	count = 0
	checkblock = drawblocknextextend
	for a := 0; a < drawaextend; a++ {
		if enemies[checkblock].activ {
			if enemies[checkblock].nextblock != checkblock {
				enemies[enemies[checkblock].nextblock] = enemies[checkblock]
				enemies[checkblock] = enemyblok{}
			}

			enemies[checkblock].count1++

			switch enemies[checkblock].name {
			case "mushroom":
				if !world[checkblock-worldw].solid && !enemies[checkblock].switch1 {
					enemies[checkblock].direction = 2
				} else {
					if rolldice() == 6 {
						if flipcoin() {
							enemies[checkblock].direction = 4
						} else {
							enemies[checkblock].direction = 6
						}
					}
					if rolldice()+rolldice() == 12 {
						enemies[checkblock].direction = 8
						enemies[checkblock].switch1 = true
						enemies[checkblock].count2 = 0
					}
					if enemies[checkblock].switch1 {
						enemies[checkblock].count2++
					}
					if enemies[checkblock].count2 > 15 {
						enemies[checkblock].switch1 = false
					}
					if enemies[checkblock].direction == 4 {
						enemies[checkblock].lr = false
					} else if enemies[checkblock].direction == 6 {
						enemies[checkblock].lr = true
					}

				}
			case "spikes":
				if !world[checkblock-worldw].solid && !enemies[checkblock].switch1 && enemies[checkblock].spikesdirection != 1 {
					enemies[checkblock].direction = 2
					enemies[checkblock].spikesdirection = 3
				} else if world[checkblock+1].solid && !world[checkblock-worldw].solid {
					enemies[checkblock].direction = 2
					enemies[checkblock].spikesdirection = 3
				} else if world[checkblock-1].solid && !world[checkblock-worldw].solid {
					enemies[checkblock].direction = 2
					enemies[checkblock].spikesdirection = 3
				} else if world[checkblock+1].solid && !world[checkblock+worldw].solid {
					enemies[checkblock].direction = 8
					enemies[checkblock].spikesdirection = 1
				} else if world[checkblock-1].solid && !world[checkblock+worldw].solid {
					enemies[checkblock].direction = 8
					enemies[checkblock].spikesdirection = 1
				}
				if rolldice()+rolldice() == 12 {
					switch rInt(1, 5) {
					case 1:
						enemies[checkblock].direction = 8
					case 2:
						enemies[checkblock].direction = 6
					case 3:
						enemies[checkblock].direction = 2
					case 4:
						enemies[checkblock].direction = 4
					}
				}

			case "raddish":
				switch rolldice() {
				case 1, 2, 3:
					enemies[checkblock].direction = 8
				case 4:
					enemies[checkblock].direction = 2
				case 5:
					enemies[checkblock].direction = 4
				case 6:
					enemies[checkblock].direction = 6
				}
				if enemies[checkblock].direction == 4 {
					enemies[checkblock].lr = false
				} else if enemies[checkblock].direction == 6 {
					enemies[checkblock].lr = true
				}

			case "bunny":
				if !world[checkblock-worldw].solid && !enemies[checkblock].switch1 {
					enemies[checkblock].direction = 2
				} else {
					if rolldice() == 6 {
						if flipcoin() {
							enemies[checkblock].direction = 4
						} else {
							enemies[checkblock].direction = 6
						}
					}
				}
				if rolldice()+rolldice() == 12 {
					enemies[checkblock].direction = 8
					enemies[checkblock].switch1 = true
					enemies[checkblock].count2 = 0
				}

				if enemies[checkblock].switch1 {
					enemies[checkblock].count2++
				}
				if enemies[checkblock].count2 > 20 {
					enemies[checkblock].switch1 = false
				}
				if enemies[checkblock].direction == 4 {
					enemies[checkblock].lr = false
				} else if enemies[checkblock].direction == 6 {
					enemies[checkblock].lr = true
				}
			case "chicken":
				if !world[checkblock-worldw].solid && !enemies[checkblock].switch1 {
					enemies[checkblock].direction = 2
				} else {

					playerv := getv(player.bloknumber)
					chickenv := getv(checkblock)

					if playerv > chickenv {
						enemies[checkblock].direction = 6
					} else {
						enemies[checkblock].direction = 4
					}

					if enemies[checkblock].direction == 6 && world[checkblock+1].solid {
						enemies[checkblock].direction = 8
						enemies[checkblock].switch1 = true
						enemies[checkblock].count2 = 0
					}
					if enemies[checkblock].direction == 4 && world[checkblock-1].solid {
						enemies[checkblock].direction = 8
						enemies[checkblock].switch1 = true
						enemies[checkblock].count2 = 0
					}
					if enemies[checkblock].switch1 {
						enemies[checkblock].count2++
					}
					if enemies[checkblock].count2 > 15 {
						enemies[checkblock].switch1 = false
					}

				}
				if enemies[checkblock].direction == 4 {
					enemies[checkblock].lr = false
				} else if enemies[checkblock].direction == 6 {
					enemies[checkblock].lr = true
				}
			case "ghost":
				if rolldice() < 4 {
					switch rolldice() {
					case 1, 2, 3:
						enemies[checkblock].direction = 8
					case 4:
						enemies[checkblock].direction = 2
					case 5:
						enemies[checkblock].direction = 4
					case 6:
						enemies[checkblock].direction = 6
					}
					if enemies[checkblock].direction == 4 {
						enemies[checkblock].lr = false
					} else if enemies[checkblock].direction == 6 {
						enemies[checkblock].lr = true
					}
				} else {
					playerv := getv(player.bloknumber)
					ghostv := getv(checkblock)
					playerh := geth(player.bloknumber)
					ghosth := geth(checkblock)

					if playerv > ghostv {
						enemies[checkblock].direction = 6
					} else if playerv < ghostv {
						enemies[checkblock].direction = 4
					} else if playerh < ghosth {
						enemies[checkblock].direction = 2
					} else if playerh > ghosth {
						enemies[checkblock].direction = 8
					}
				}
			}

		}
		checkblock++
		count++
		if count == drawwextend {
			count = 0
			checkblock -= drawwextend
			checkblock += worldw
		}
	}

}

func updatebojects() { // MARK: updatebojects

	count := 0
	checkblock := drawblocknextextend
	for a := 0; a < drawaextend; a++ {

		if objects[checkblock].activ && objects[checkblock].coin {
			if !world[checkblock-worldw].solid {
				objects[checkblock].nextblock = checkblock - worldw
			}
		}

		// bullet movements
		if bullets[checkblock].activ {
			if bullets[checkblock].direction {
				if !world[checkblock-1].solid {
					bullets[checkblock].nextblock = checkblock - 1
				} else {
					bullets[checkblock] = bullet{}
				}
			} else {
				if !world[checkblock+1].solid {
					bullets[checkblock].nextblock = checkblock + 1
				} else {
					bullets[checkblock] = bullet{}
				}
			}

		}

		checkblock++
		count++
		if count == drawwextend {
			count = 0
			checkblock -= drawwextend
			checkblock += worldw
		}
	}

	count = 0
	checkblock = drawblocknextextend
	for a := 0; a < drawaextend; a++ {

		if objects[checkblock].activ && objects[checkblock].coin {
			if objects[checkblock].nextblock != checkblock {
				objects[objects[checkblock].nextblock] = objects[checkblock]
				objects[checkblock] = blok{}
			}
		}
		// bullet nextblock update
		if bullets[checkblock].nextblock != checkblock {
			bullets[bullets[checkblock].nextblock] = bullets[checkblock]
			bullets[checkblock] = bullet{}

		}

		checkblock++
		count++
		if count == drawwextend {
			count = 0
			checkblock -= drawwextend
			checkblock += worldw
		}
	}

}
func updateplayer() { // MARK: updateplayer

	if !freemoveon {
		drawbloknext = player.bloknumber
		drawbloknext -= draww / 2
		drawbloknext -= (drawh / 3) * worldw
		drawblocknextextend = drawbloknext - draww
	}

	if jumpon {
		jump()
	}
	if !world[player.bloknumber-worldw].solid && !jumpon {
		fallon = true
	}
	if fallon {
		fall()
	}

}
func moveenemy(direction, blocknumber int) { // MARK: moveenemy

	switch direction {
	case 1:
		if !world[(blocknumber-worldw)-1].solid {
			enemies[blocknumber].nextblock = (blocknumber - worldw) - 1
		}
	case 2:
		if !world[(blocknumber - worldw)].solid {
			enemies[blocknumber].nextblock = (blocknumber - worldw)
		}
	case 3:
		if !world[(blocknumber-worldw)+1].solid {
			enemies[blocknumber].nextblock = (blocknumber - worldw) + 1
		}
	case 4:
		if !world[blocknumber-1].solid {
			enemies[blocknumber].nextblock = blocknumber - 1
		}
	case 6:
		if !world[blocknumber+1].solid {
			enemies[blocknumber].nextblock = blocknumber + 1
		}
	case 7:
		if !world[(blocknumber+worldw)-1].solid {
			enemies[blocknumber].nextblock = (blocknumber + worldw) - 1
		}
	case 8:
		if !world[(blocknumber + worldw)].solid {
			enemies[blocknumber].nextblock = (blocknumber + worldw)
		}
	case 9:
		if !world[(blocknumber+worldw)+1].solid {
			enemies[blocknumber].nextblock = (blocknumber + worldw) + 1
		}
	}

}
func playeraction() { // MARK: playeraction  ███████████████████████████████
	if player.direction {
		bullets[player.bloknumber-2].activ = true
		bullets[player.bloknumber-2].direction = true
		switch player.holding {
		case "uzzi":
			bullets[player.bloknumber-2].name = "uzzi"
			bullets[player.bloknumber-2].hp = 1
		case "shotgun":
			bullets[player.bloknumber-2].name = "shotgun"
		case "bazooka":
			bullets[player.bloknumber-2].name = "bazooka"
		}
	} else {
		bullets[player.bloknumber+2].activ = true
		bullets[player.bloknumber-2].direction = false
		switch player.holding {
		case "uzzi":
			bullets[player.bloknumber+2].name = "uzzi"
			bullets[player.bloknumber-2].hp = 1
		case "shotgun":
			bullets[player.bloknumber+2].name = "shotgun"
		case "bazooka":
			bullets[player.bloknumber+2].name = "bazooka"
		}
	}

}
func jump() { // MARK: jump

	if jumph > 0 {
		if world[player.bloknumber+worldw].powerup {
			objects[player.bloknumber+(worldw*2)].activ = true
			objects[player.bloknumber+(worldw*2)].pickup = true
			switch rolldice() {
			case 1:
				objects[player.bloknumber+(worldw*2)].img = axeimg
			case 2:
				objects[player.bloknumber+(worldw*2)].img = shotgunimg
			case 3:
				objects[player.bloknumber+(worldw*2)].img = bazookaimg
			case 4:
				objects[player.bloknumber+(worldw*2)].img = uzziimg
			case 5:
				objects[player.bloknumber+(worldw*2)].img = scorpionimg
				objects[player.bloknumber+(worldw*2)].pickup = false
			case 6:
				objects[player.bloknumber+(worldw*2)].img = bombimg
			}
			world[player.bloknumber+worldw].color1 = randomgrey()
			jumpon = false
			fallon = true

		} else {
			player.bloknumber += worldw
			jumph--
		}
	} else {
		jumpon = false
		fallon = true
	}

}
func fall() { // MARK: fall
	if !world[player.bloknumber-worldw].solid {
		player.bloknumber -= worldw
	} else {
		fallon = false
		player.moving = false

	}

}
func dropobj(playerbloknumber int) { // MARK: dropobj

	count := 0
	for {
		switch rolldice() {
		case 1:
			if !objects[playerbloknumber-1].activ && !world[playerbloknumber-1].solid {
				objects[playerbloknumber-1].activ = true
				objects[playerbloknumber-1].img = player.objimg
				count = 6
			}
		case 2:
			if !objects[(playerbloknumber-1)+worldw].activ && !world[(playerbloknumber-1)+worldw].solid {
				objects[(playerbloknumber-1)+worldw].activ = true
				objects[(playerbloknumber-1)+worldw].img = player.objimg
				count = 6
			}
		case 3:
			if !objects[(playerbloknumber)+worldw].activ && !world[(playerbloknumber)+worldw].solid {
				objects[(playerbloknumber)+worldw].activ = true
				objects[(playerbloknumber)+worldw].img = player.objimg
				count = 6
			}
		case 4:
			if !objects[(playerbloknumber+1)+worldw].activ && !world[(playerbloknumber+1)+worldw].solid {
				objects[(playerbloknumber+1)+worldw].activ = true
				objects[(playerbloknumber+1)+worldw].img = player.objimg
				count = 6
			}
		case 5:
			if !objects[playerbloknumber+1].activ && !world[playerbloknumber+1].solid {
				objects[playerbloknumber+1].activ = true
				objects[playerbloknumber+1].img = player.objimg
				count = 6
			}
		case 6:
			if flipcoin() {
				if !objects[(playerbloknumber+1)-worldw].activ && !world[(playerbloknumber+1)-worldw].solid {
					objects[(playerbloknumber+1)-worldw].activ = true
					objects[(playerbloknumber+1)-worldw].img = player.objimg
					count = 6
				}
			} else {
				if !objects[(playerbloknumber-1)-worldw].activ && !world[(playerbloknumber-1)-worldw].solid {
					objects[(playerbloknumber-1)-worldw].activ = true
					objects[(playerbloknumber-1)-worldw].img = player.objimg
					count = 6
				}
			}
		}

		count++
		if count >= 6 {
			break
		}
	}

}
func drawnocameraback() { // MARK: drawnocameraback
	if backgon {
		if lineson {
			for a := 0; a < len(background); a++ {
				//falling stars
				rl.DrawCircleGradient(int(background[a].v2.X), int(background[a].v2.Y), float32(background[a].w/2), rl.Fade(background[a].color1, background[a].opac), rl.Transparent)

				background[a].v2.X += float32(rInt(5, 10))
				background[a].v2.Y += float32(rInt(5, 10))
				if background[a].v2.X > float32(monw) {
					background[a].v2.X = 0
				}
				if background[a].v2.X < 0 {
					background[a].v2.X = float32(monw)
				}
				if background[a].v2.Y > float32(monh) {
					background[a].v2.Y = -100
				}
				background[a].w++

				if background[a].w >= 200 {
					background[a].w = rInt(20, 40)
				}
			}
			rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.DarkGray, 0.1))
		} else {
			// backg tiles
			rl.BeginMode2D(camerabackg)
			count := 0
			drawblok = drawbloknext
			x := 0
			y := monh - tilesize

			for a := 0; a < drawa; a++ {

				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
				rl.DrawTexturePro(imgs, backtiles[a], destrec, origin, 0, rl.Fade(rl.DarkPurple, fadeblink2))
				if ghosting {
					destrec.X += rFloat32(-6, 7)
					destrec.Y += rFloat32(-6, 7)
					rl.DrawTexturePro(imgs, backtiles[a], destrec, origin, 0, rl.Fade(rl.DarkPurple, 0.1))
				}
				x += tilesize
				count++

				if count == draww {
					count = 0
					x = 0
					y -= tilesize
				}
			}
			rl.EndMode2D()

		}

	}
}
func drawlayers() { // MARK: drawlayers ███████████████████████████████

	if !paused {
		// layer 1 bloks
		count := 0
		drawblok = drawbloknext
		x := 0
		y := monh - tilesize

		for a := 0; a < drawa; a++ {
			if world[drawblok].activ {
				if world[drawblok].solid && world[drawblok].ground && !world[drawblok].portal { //ground tiles

					if lineson {
						change := rInt(8, 21)
						if change%2 != 0 {
							change++
						}
						rl.DrawRectangleLines(x, y, tilesize, tilesize, world[drawblok].color1)
						rl.DrawRectangleLines(x+change/2, y+change/2, tilesize+change/2, tilesize-change, world[drawblok].color1)

						x += rInt(-2, 3)
						y += rInt(-2, 3)
						if ghosting {
							rl.DrawRectangleLines(x, y, tilesize, tilesize, rl.Fade(world[drawblok].color1, 0.5))
						}
					} else {

						origin := rl.NewVector2(float32(0), float32(0))
						destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))

						rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, world[drawblok].color1)

						if ghosting {
							destrec.X += rFloat32(-2, 3)
							destrec.Y += rFloat32(-2, 3)
							rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(rl.White, 0.1))
						}
					}
				} else if world[drawblok].solid && !world[drawblok].ground && !world[drawblok].portal { //blok tiles

					if lineson {
						if world[drawblok].powerup {
							origin := rl.NewVector2(float32(0), float32(0))
							destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
							if rolldice() == 6 {
								destrec.Y -= rFloat32(3, 7)
							}
							rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, world[drawblok].color1)
							rl.DrawRectangleLines(int(destrec.X), int(destrec.Y), tilesize, tilesize, rl.Black)
							if ghosting {
								destrec.X += rFloat32(-3, 4)
								destrec.Y += rFloat32(-3, 4)
								rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(world[drawblok].color1, 0.4))
							}
						} else {
							change := rInt(8, 21)
							if change%2 != 0 {
								change++
							}
							rl.DrawRectangleLines(x, y, tilesize, tilesize, world[drawblok].color1)
							rl.DrawRectangleLines(x+change/2, y+change/2, tilesize+change/2, tilesize-change, world[drawblok].color1)

							x += rInt(-2, 3)
							y += rInt(-2, 3)
							if ghosting {
								rl.DrawRectangleLines(x, y, tilesize, tilesize, rl.Fade(world[drawblok].color1, 0.5))
							}
						}
					} else {
						origin := rl.NewVector2(float32(0), float32(0))
						destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))

						if world[drawblok].powerup {
							if rolldice() == 6 {
								destrec.Y -= rFloat32(3, 7)
							}
							rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, world[drawblok].color1)
							rl.DrawRectangleLines(int(destrec.X), int(destrec.Y), tilesize, tilesize, rl.Black)
							if ghosting {
								destrec.X += rFloat32(-3, 4)
								destrec.Y += rFloat32(-3, 4)
								rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(world[drawblok].color1, 0.4))
							}
						} else {
							rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, world[drawblok].color1)
							if ghosting {
								destrec.X += rFloat32(-3, 4)
								destrec.Y += rFloat32(-3, 4)
								rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(world[drawblok].color1, 0.4))
							}
						}
					}

					// portal
				} else if world[drawblok].portal {

					if player.bloknumber == drawblok {
						levelend = true
					}

					origin := rl.NewVector2(float32(0), float32(0))
					destrec := rl.NewRectangle(float32(x-(tilesize/2+24)), float32(y-(tilesize/2+24)), float32(portal.Width*5), float32(portal.Height*5))

					rl.DrawTexturePro(imgs, portal, destrec, origin, 0, rl.White)

					rl.DrawRectangle(x, y, tilesize, tilesize, rl.Fade(rl.Magenta, 0.5))

					//	v2 := rl.NewVector2(float32(x), float32(y))
					//	rl.DrawTextureRec(imgs, portal, v2, rl.White)

				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}

		// layer 2 trees
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {
			if world[drawblok].activ {
				if world[drawblok].imgon {
					origin := rl.NewVector2(float32(0), float32(0))
					destrec := rl.NewRectangle(float32(x-170), float32(y-(218-tilesize)), 256, 256)
					rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(rl.Black, 0.5))
					destrec = rl.NewRectangle(float32(x-160), float32(y-(228-tilesize)), 256, 256)
					rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, world[drawblok].img, destrec, origin, 0, rl.Fade(rl.White, 0.5))
					}

				}
			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}
		// layer 3 objects
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {

			if objects[drawblok].activ {
				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
				if rolldice() == 6 {
					destrec.Y -= rFloat32(4, 11)
				}
				if objects[drawblok].coin {
					destrec = rl.NewRectangle(float32(x), float32(y+12), float32(tilesize/2)+16, float32(tilesize/2)+16)
					rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-5, 6)
						destrec.Y += rFloat32(-5, 6)
						rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				} else {
					rl.DrawTexturePro(imgs, objects[drawblok].img, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-3, 4)
						destrec.Y += rFloat32(-3, 4)
						rl.DrawTexturePro(imgs, objects[drawblok].img, destrec, origin, 0, rl.Fade(rl.White, 0.4))
					}
				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}
		// layer 4 monsters
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {

			if monsters[drawblok].activ {

				if player.bloknumber == drawblok || player.bloknumber-1 == drawblok || player.bloknumber-2 == drawblok || player.bloknumber+1 == drawblok || player.bloknumber+2 == drawblok {
					switch portaldirection {
					case 1:
						textlen := rl.MeasureText("down left", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("down left", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 2:
						textlen := rl.MeasureText("down", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("down", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 3:
						textlen := rl.MeasureText("down right", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("down right", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 4:
						textlen := rl.MeasureText("left", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("left", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 6:
						textlen := rl.MeasureText("right", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("right", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 7:
						textlen := rl.MeasureText("up left", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("up left", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 8:
						textlen := rl.MeasureText("up", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("up", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					case 9:
						textlen := rl.MeasureText("up right", 40)
						rl.DrawRectangle((x+tilesize/2)-((textlen/2)+15), y-55, textlen+30, 50, rl.White)
						rl.DrawText("up right", (x+tilesize/2)-(textlen/2), y-50, 40, rl.Black)
					}
				}

				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))

				switch monsters[drawblok].name {
				case "monster25":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster25.Width*4), float32(monster25.Height*4))
					rl.DrawTexturePro(imgs, monster25, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster25.Width*4), float32(monster25.Height*4))
					rl.DrawTexturePro(imgs, monster25, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster25, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				//	rl.DrawRectangle(x, y, tilesize, tilesize, rl.Fade(rl.Magenta, 0.5))
				case "monster24":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster24.Width*4), float32(monster24.Height*4))
					rl.DrawTexturePro(imgs, monster24, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster24.Width*4), float32(monster24.Height*4))
					rl.DrawTexturePro(imgs, monster24, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster24, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster23":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster23.Width*4), float32(monster23.Height*4))
					rl.DrawTexturePro(imgs, monster23, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster23.Width*4), float32(monster23.Height*4))
					rl.DrawTexturePro(imgs, monster23, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster23, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster22":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster22.Width*4), float32(monster22.Height*4))
					rl.DrawTexturePro(imgs, monster22, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster22.Width*4), float32(monster22.Height*4))
					rl.DrawTexturePro(imgs, monster22, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster22, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster21":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster21.Width*4), float32(monster21.Height*4))
					rl.DrawTexturePro(imgs, monster21, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster21.Width*4), float32(monster21.Height*4))
					rl.DrawTexturePro(imgs, monster21, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster21, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster20":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster20.Width*4), float32(monster20.Height*4))
					rl.DrawTexturePro(imgs, monster20, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster20.Width*4), float32(monster20.Height*4))
					rl.DrawTexturePro(imgs, monster20, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster20, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster19":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster19.Width*4), float32(monster19.Height*4))
					rl.DrawTexturePro(imgs, monster19, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster19.Width*4), float32(monster19.Height*4))
					rl.DrawTexturePro(imgs, monster19, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster19, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster18":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster18.Width*4), float32(monster18.Height*4))
					rl.DrawTexturePro(imgs, monster18, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster18.Width*4), float32(monster18.Height*4))
					rl.DrawTexturePro(imgs, monster18, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster18, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster17":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster17.Width*4), float32(monster17.Height*4))
					rl.DrawTexturePro(imgs, monster17, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster17.Width*4), float32(monster17.Height*4))
					rl.DrawTexturePro(imgs, monster17, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster17, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster16":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster16.Width*4), float32(monster16.Height*4))
					rl.DrawTexturePro(imgs, monster16, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster16.Width*4), float32(monster16.Height*4))
					rl.DrawTexturePro(imgs, monster16, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster16, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster15":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster15.Width*4), float32(monster15.Height*4))
					rl.DrawTexturePro(imgs, monster15, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster15.Width*4), float32(monster15.Height*4))
					rl.DrawTexturePro(imgs, monster15, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster15, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster14":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster14.Width*4), float32(monster14.Height*4))
					rl.DrawTexturePro(imgs, monster14, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster14.Width*4), float32(monster14.Height*4))
					rl.DrawTexturePro(imgs, monster14, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster14, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster13":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster13.Width*4), float32(monster13.Height*4))
					rl.DrawTexturePro(imgs, monster13, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster13.Width*4), float32(monster13.Height*4))
					rl.DrawTexturePro(imgs, monster13, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster13, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster12":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster12.Width*4), float32(monster12.Height*4))
					rl.DrawTexturePro(imgs, monster12, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster12.Width*4), float32(monster12.Height*4))
					rl.DrawTexturePro(imgs, monster12, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster12, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster11":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster11.Width*4), float32(monster11.Height*4))
					rl.DrawTexturePro(imgs, monster11, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster11.Width*4), float32(monster11.Height*4))
					rl.DrawTexturePro(imgs, monster11, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster11, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster10":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster10.Width*4), float32(monster10.Height*4))
					rl.DrawTexturePro(imgs, monster10, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster10.Width*4), float32(monster10.Height*4))
					rl.DrawTexturePro(imgs, monster10, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster10, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster9":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster9.Width*4), float32(monster9.Height*4))
					rl.DrawTexturePro(imgs, monster9, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster9.Width*4), float32(monster9.Height*4))
					rl.DrawTexturePro(imgs, monster9, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster9, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster8":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster8.Width*4), float32(monster8.Height*4))
					rl.DrawTexturePro(imgs, monster8, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster8.Width*4), float32(monster8.Height*4))
					rl.DrawTexturePro(imgs, monster8, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster8, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster7":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster7.Width*4), float32(monster7.Height*4))
					rl.DrawTexturePro(imgs, monster7, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster7.Width*4), float32(monster7.Height*4))
					rl.DrawTexturePro(imgs, monster7, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster7, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster6":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster6.Width*4), float32(monster6.Height*4))
					rl.DrawTexturePro(imgs, monster6, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster6.Width*4), float32(monster6.Height*4))
					rl.DrawTexturePro(imgs, monster6, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster6, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster5":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster5.Width*4), float32(monster5.Height*4))
					rl.DrawTexturePro(imgs, monster5, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster5.Width*4), float32(monster5.Height*4))
					rl.DrawTexturePro(imgs, monster5, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster5, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster4":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster4.Width*4), float32(monster4.Height*4))
					rl.DrawTexturePro(imgs, monster4, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster4.Width*4), float32(monster4.Height*4))
					rl.DrawTexturePro(imgs, monster4, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster4, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster3":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster3.Width*4), float32(monster3.Height*4))
					rl.DrawTexturePro(imgs, monster3, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster3.Width*4), float32(monster3.Height*4))
					rl.DrawTexturePro(imgs, monster3, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster3, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster2":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster2.Width*4), float32(monster2.Height*4))
					rl.DrawTexturePro(imgs, monster2, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster2.Width*4), float32(monster2.Height*4))
					rl.DrawTexturePro(imgs, monster2, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster2, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}
				case "monster1":
					destrec = rl.NewRectangle(float32(x+8), float32(y+32), float32(monster1.Width*4), float32(monster1.Height*4))
					rl.DrawTexturePro(imgs, monster1, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
					destrec = rl.NewRectangle(float32(x+12), float32(y+28), float32(monster1.Width*4), float32(monster1.Height*4))
					rl.DrawTexturePro(imgs, monster1, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-4, 5)
						destrec.Y += rFloat32(-4, 5)
						rl.DrawTexturePro(imgs, monster1, destrec, origin, 0, rl.Fade(rl.White, 0.7))
					}

				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}

		// layer 5 bullets
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {

			if bullets[drawblok].activ {

				switch bullets[drawblok].name {
				case "uzzi":

					rl.DrawCircle(x, y+tilesize/2, 12, rl.Black)
					rl.DrawCircle(x, y+tilesize/2, 10, rl.White)
					rl.DrawCircle(x+tilesize/3, y+tilesize/2, 12, rl.Black)
					rl.DrawCircle(x+tilesize/3, y+tilesize/2, 10, rl.White)
					rl.DrawCircle(x+((tilesize/3)*2), y+tilesize/2, 12, rl.Black)
					rl.DrawCircle(x+((tilesize/3)*2), y+tilesize/2, 10, rl.White)

				case "shotgun":
					rl.DrawCircle(x, y, 50, rl.Yellow)
				case "bazooka":
					rl.DrawCircle(x, y, 50, rl.Orange)
				}

				if onoff3 {

					if bullets[drawblok].pause {
						bullets[drawblok].pause = false
					}
				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}

		// layer 6 enemies
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {

			if enemies[drawblok].activ {

				if enemies[drawblok].hp <= 0 && bullets[drawblok].activ {
					enemies[drawblok] = enemyblok{}
					destroyx = x
					destroyy = y
					destroyon = true
					destroytimer = 1
				}

				enemies[drawblok].x = x
				enemies[drawblok].y = y
				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))

				switch enemies[drawblok].name {
				case "raddish":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-12), float32(y-tilesize/2+4), float32(raddishimg.Width*4), float32(raddishimg.Height*4))
						rl.DrawTexturePro(imgs, raddishimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-8), float32(y-tilesize/2), float32(raddishimg.Width*4), float32(raddishimg.Height*4))
						rl.DrawTexturePro(imgs, raddishimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, raddishimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-4), float32(y-tilesize/2+4), float32(raddishimg.Width*4), float32(raddishimg.Height*4))
						rl.DrawTexturePro(imgs, raddishlimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-8), float32(y-tilesize/2), float32(raddishimg.Width*4), float32(raddishimg.Height*4))
						rl.DrawTexturePro(imgs, raddishlimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, raddishlimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					}
				case "ghost":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-tilesize/2+8), float32(y-14), float32(ghostimg.Width*4), float32(ghostimg.Height*4))
						rl.DrawTexturePro(imgs, ghostimg, destrec, origin, 0, rl.Fade(rl.White, rF32(0.2, 0.9)))
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, ghostimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-tilesize/2+8), float32(y-14), float32(ghostimg.Width*4), float32(ghostimg.Height*4))
						rl.DrawTexturePro(imgs, ghostlimg, destrec, origin, 0, rl.Fade(rl.White, rF32(0.2, 0.9)))
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, ghostlimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					}
				case "chicken":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-tilesize-4), float32(y-(tilesize)-40+4), float32(chickenimg.Width*8), float32(chickenimg.Height*8))
						rl.DrawTexturePro(imgs, chickenimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize), float32(y-(tilesize)-40), float32(chickenimg.Width*8), float32(chickenimg.Height*8))
						rl.DrawTexturePro(imgs, chickenimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, chickenimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-tilesize+4), float32(y-(tilesize)-40+4), float32(chickenimg.Width*8), float32(chickenimg.Height*8))
						rl.DrawTexturePro(imgs, chickenlimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize), float32(y-(tilesize)-40), float32(chickenimg.Width*8), float32(chickenimg.Height*8))
						rl.DrawTexturePro(imgs, chickenlimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, chickenlimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					}
				case "spikes":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-6), float32(y-6), float32(spikesimg.Width*4), float32(spikesimg.Height*4))
						rl.DrawTexturePro(imgs, spikesimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x), float32(y), float32(spikesimg.Width*4), float32(spikesimg.Height*4))
						rl.DrawTexturePro(imgs, spikesimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, spikesimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-6), float32(y-6), float32(spikesimg.Width*4), float32(spikesimg.Height*4))
						rl.DrawTexturePro(imgs, spikeslimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x), float32(y), float32(spikesimg.Width*4), float32(spikesimg.Height*4))
						rl.DrawTexturePro(imgs, spikeslimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, spikeslimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					}
				case "mushroom":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-tilesize/2-4), float32(y-tilesize+4), float32(mushroomimg.Width*6), float32(mushroomimg.Height*6))
						rl.DrawTexturePro(imgs, mushroomimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize/2), float32(y-tilesize), float32(mushroomimg.Width*6), float32(mushroomimg.Height*6))
						rl.DrawTexturePro(imgs, mushroomimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, mushroomimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-tilesize/2+4), float32(y-tilesize+4), float32(mushroomimg.Width*6), float32(mushroomimg.Height*6))
						rl.DrawTexturePro(imgs, mushroomlimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize/2), float32(y-tilesize), float32(mushroomimg.Width*6), float32(mushroomimg.Height*6))
						rl.DrawTexturePro(imgs, mushroomlimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, mushroomlimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					}
				case "bunny":
					if enemies[drawblok].lr {
						destrec = rl.NewRectangle(float32(x-tilesize/2-4), float32(y-tilesize-tilesize/2+4), float32(bunnyimg.Width*6), float32(bunnyimg.Height*6))
						rl.DrawTexturePro(imgs, bunnyimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize/2), float32(y-tilesize-tilesize/2), float32(bunnyimg.Width*6), float32(bunnyimg.Height*6))
						rl.DrawTexturePro(imgs, bunnyimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, bunnyimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}
					} else {
						destrec = rl.NewRectangle(float32(x-tilesize/2+4), float32(y-tilesize-tilesize/2+4), float32(bunnyimg.Width*6), float32(bunnyimg.Height*6))
						rl.DrawTexturePro(imgs, bunnylimg, destrec, origin, 0, rl.Fade(rl.Black, 0.8))
						destrec = rl.NewRectangle(float32(x-tilesize/2), float32(y-tilesize-tilesize/2), float32(bunnyimg.Width*6), float32(bunnyimg.Height*6))
						rl.DrawTexturePro(imgs, bunnylimg, destrec, origin, 0, rl.White)
						if ghosting {
							destrec.X += rFloat32(-4, 5)
							destrec.Y += rFloat32(-4, 5)
							rl.DrawTexturePro(imgs, bunnylimg, destrec, origin, 0, rl.Fade(rl.White, 0.7))
						}

					}
				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}

		// layer 7 player
		count = 0
		drawblok = drawbloknext
		x = 0
		y = monh - tilesize

		for a := 0; a < drawa; a++ {

			if player.bloknumber == drawblok {

				if objects[drawblok].activ {
					if objects[drawblok].coin {
						objects[drawblok].coin = false
						objects[drawblok].activ = false
						objects[drawblok] = blok{}
						coinstotal++
					}
					if objects[drawblok].pickup {
						if player.isholding {
							dropobj(drawblok)
							player.objimg = objects[drawblok].img
							switch player.objimg {
							case axeimg:
								player.holding = "axe"
								player.objlimg = axelimg
							case shotgunimg:
								player.holding = "shotgun"
								player.objlimg = shotgunlimg
							case bazookaimg:
								player.holding = "bazooka"
								player.objlimg = bazookalimg
							case uzziimg:
								player.holding = "uzzi"
								player.objlimg = uzzilimg
							case bombimg:
								player.holding = "bomb"
								player.objlimg = bomblimg
							}
							objects[drawblok] = blok{}
						} else {
							player.isholding = true
							player.objimg = objects[drawblok].img
							switch player.objimg {
							case axeimg:
								player.holding = "axe"
								player.objlimg = axelimg
							case shotgunimg:
								player.holding = "shotgun"
								player.objlimg = shotgunlimg
							case bazookaimg:
								player.holding = "bazooka"
								player.objlimg = bazookalimg
							case uzziimg:
								player.holding = "uzzi"
								player.objlimg = uzzilimg
							case bombimg:
								player.holding = "bomb"
								player.objlimg = bomblimg
							}
							objects[drawblok] = blok{}
						}

						currentpickuptext = player.holding
						pickuptexttimer = 3
						pickuptextychange = 70
						pickuptexttimercount = 0
						pickuptexton = true
					}

				}

				//	rl.DrawCircle(x+24, y+24, 24, randomcolor())
				player.x = x
				player.y = y

				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))

				// player object
				if player.isholding {
					playobjimg := player.objimg
					switch player.holding {
					case "bazooka":
						destrec2 := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
						if player.direction {
							destrec2.X -= 20
							destrec2.Y -= 35
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
							playobjimg = player.objlimg
						} else {
							destrec2.X += 20
							destrec2.Y -= 35
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
						}
						rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.White)
						if ghosting {
							destrec2.X += rFloat32(-3, 4)
							destrec2.Y += rFloat32(-3, 4)
							rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.Fade(rl.White, 0.3))
						}
					case "uzzi":
						destrec2 := rl.NewRectangle(float32(x), float32(y), float32(tilesize-10), float32(tilesize-10))
						if player.direction {
							destrec2.X -= 34
							destrec2.Y += 7
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
							playobjimg = player.objlimg
						} else {
							destrec2.X += 44
							destrec2.Y += 7
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
						}
						rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.White)
						if ghosting {
							destrec2.X += rFloat32(-3, 4)
							destrec2.Y += rFloat32(-3, 4)
							rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.Fade(rl.White, 0.3))
						}
					case "bomb":
						destrec2 := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
						if player.direction {
							destrec2.X -= 50
							destrec2.Y -= 30
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
							playobjimg = player.objlimg
						} else {
							destrec2.X += 60
							destrec2.Y -= 30
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
						}
						rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.White)
						if ghosting {
							destrec2.X += rFloat32(-3, 4)
							destrec2.Y += rFloat32(-3, 4)
							rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.Fade(rl.White, 0.3))
						}
					case "shotgun":
						destrec2 := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
						if player.direction {
							destrec2.X -= 54
							destrec2.Y += 2
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
							playobjimg = player.objlimg
						} else {
							destrec2.X += 57
							destrec2.Y += 2
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
						}
						rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.White)
						if ghosting {
							destrec2.X += rFloat32(-3, 4)
							destrec2.Y += rFloat32(-3, 4)
							rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, 0, rl.Fade(rl.White, 0.3))
						}
					case "axe":
						origin = rl.NewVector2(float32(0), float32(tilesize))
						destrec2 := rl.NewRectangle(float32(x), float32(y), float32(tilesize), float32(tilesize))
						if player.direction {
							if axerotation == 45 {

								axerotation = -45
							}
							destrec2.X -= 54
							destrec2.Y -= 30
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
							playobjimg = player.objlimg
						} else {
							destrec2.X += 60
							destrec2.Y -= 30
							if rolldice() == 6 {
								destrec2.Y += rFloat32(-4, 5)
							}
						}
						rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, axerotation, rl.White)
						if ghosting {
							destrec2.X += rFloat32(-3, 4)
							destrec2.Y += rFloat32(-3, 4)
							rl.DrawTexturePro(imgs, playobjimg, destrec2, origin, axerotation, rl.Fade(rl.White, 0.3))
						}
					}
				}

				// player emote
				if emoteon {
					emoterec := rl.NewRectangle(float32(x+20), float32(y-tilesize/2)+4, float32(tilesize/2), float32(tilesize/2))
					if rolldice() == 6 {
						emoterec.Y += rFloat32(-8, 9)
					}
					rl.DrawTexturePro(imgs, currentemote, emoterec, origin, 0, rl.Fade(rl.White, rF32(0.4, 0.8)))
				}
				// player img
				if player.direction {
					if rolldice() == 6 {
						destrec.Y -= rFloat32(2, 5)
					}
					rl.DrawTexturePro(imgs, dinol, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-3, 4)
						destrec.Y += rFloat32(-3, 4)
						rl.DrawTexturePro(imgs, dinol, destrec, origin, 0, rl.Fade(rl.White, 0.3))
					}
				} else {
					if rolldice() == 6 {
						destrec.Y -= rFloat32(2, 5)
					}
					rl.DrawTexturePro(imgs, dinor, destrec, origin, 0, rl.White)
					if ghosting {
						destrec.X += rFloat32(-3, 4)
						destrec.Y += rFloat32(-3, 4)
						rl.DrawTexturePro(imgs, dinor, destrec, origin, 0, rl.Fade(rl.White, 0.3))
					}
				}

			}

			x += tilesize
			count++
			drawblok++
			if count == draww {
				count = 0
				drawblok += worldw
				drawblok -= draww
				x = 0
				y -= tilesize
			}
		}

	}
}
func drawnocamera() { // MARK: drawnocamera

	if introon {
		drawintro()
	}
	if startnewgameon {
		drawstartnewgame()
	}
	if options {
		drawoptionsmenu()
	}

	if weather {
		drawweather()
	}

	if !options && !paused {
		drawgamebar()
	}

	// screen fx
	if scanlines {
		for a := 0; a < monh; a++ {
			rl.DrawLine(0, a, monw, a, rl.Fade(rl.Black, 0.2))
			a += 2
		}
	}
	if pixelnoise {
		for a := 0; a < 100; a++ {
			width := rFloat32(1, 3)
			rec := rl.NewRectangle(rFloat32(0, monw), rFloat32(0, monh), width, width)
			rl.DrawRectangleRec(rec, rl.Fade(rl.Black, rF32(0.4, 1.1)))
		}

	}

}
func drawdestroy() { // MARK: drawdestroy

	number := rInt(8, 13)
	if destroytimer > 0 {
		for a := 0; a < number; a++ {
			x := destroyx + rInt(-100, 101)
			y := destroyy + rInt(-100, 101)
			radius := rFloat32(10, 40)
			rl.DrawCircle(x, y, radius, rl.Fade(randomcolor(), rF32(0.4, 1.0)))
		}

	}

}
func drawgamebar() { // MARK: drawgamebar
	rl.DrawRectangle(0, monh-(tilesize+16), monw, tilesize+16, rl.Fade(rl.Black, 0.8))

	// coins icon
	menubarx := 20
	menubary := monh - (tilesize)
	origin := rl.NewVector2(float32(0), float32(0))
	destrec := rl.NewRectangle(float32(menubarx), float32(menubary+4), coinimg.Width*4, coinimg.Height*4)
	rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.White)
	if ghosting {
		destrec.X += rFloat32(-3, 4)
		destrec.Y += rFloat32(-3, 4)
		rl.DrawTexturePro(imgs, coinimg, destrec, origin, 0, rl.Fade(rl.White, 0.4))
	}
	// coins text
	menubarx += 80
	coinstext := strconv.Itoa(coinstotal)
	rl.DrawText(coinstext, menubarx, menubary, textsize1, rl.White)
	// time icon
	origin = rl.NewVector2(float32(0), float32(0))
	destrec = rl.NewRectangle(float32(monw-210), float32(menubary+4), timeimg.Width*4, timeimg.Height*4)
	if gametime > 99 {
		destrec.X -= 30
	}
	rl.DrawTexturePro(imgs, timeimg, destrec, origin, 0, rl.White)
	// time text
	gametimetext := strconv.Itoa(gametime)
	if gametime > 99 {
		rl.DrawText(gametimetext, monw-140, menubary, textsize1, rl.White)
	} else {
		rl.DrawText(gametimetext, monw-120, menubary, textsize1, rl.White)
	}
	// hp icons
	menubarx = monw / 2
	menubarx -= player.hp / 2 * 100

	origin = rl.NewVector2(float32(0), float32(0))
	destrec = rl.NewRectangle(float32(menubarx), float32(menubary+4), hpimg.Width*2, hpimg.Height*2)
	for a := 0; a < player.hp; a++ {
		destrec.X -= 8
		destrec.Y += 4
		rl.DrawTexturePro(imgs, hpimg, destrec, origin, 0, darkred())
		destrec.X += 8
		destrec.Y -= 4
		rl.DrawTexturePro(imgs, hpimg, destrec, origin, 0, randomred())
		rl.DrawTexturePro(imgs, hpimg, destrec, origin, 0, rl.Fade(brightred(), fadeblink))
		destrec.X += (hpimg.Width * 2) + 16
	}
}
func drawintro() { // MARK: drawintro
	paused = true
	rl.DrawRectangle(0, 0, monw, monh, rl.Black)

	y := float32(0)
	x := float32(0)

	for a := 0; a < len(dinomultiplier); a++ {

		destrec := rl.NewRectangle(x, y, float32(dinomultiplier[a]*int(dinor.Height)), float32(dinomultiplier[a]*int(dinor.Height)))
		origin := rl.NewVector2(0, 0)
		switch dinotype[a] {
		case 1:
			rl.DrawTexturePro(imgs, dinor, destrec, origin, 0, rl.Fade(rl.White, fadeblink))
		case 2:
			rl.DrawTexturePro(imgs, dino2r, destrec, origin, 0, rl.Fade(rl.White, fadeblink))
		case 3:
			rl.DrawTexturePro(imgs, dino3r, destrec, origin, 0, rl.Fade(rl.White, fadeblink))
		case 4:
			rl.DrawTexturePro(imgs, dino4r, destrec, origin, 0, rl.Fade(rl.White, fadeblink))
		}

		x += float32(dinomultiplier[a]) * dinor.Height

		if x > float32(monw) {
			x = 0
			y += dinor.Height * 4
		}

		if y > float32(monh) {
			break
		}

	}
	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, float32(backfade)))

	textlen := rl.MeasureText("dino", 200)

	rl.DrawText("dino", monw/2-(textlen/2)+rInt(-5, 6), monh/2-250+rInt(-5, 6), 200, rl.Fade(randomcolor(), float32(introtextfade)))
	rl.DrawText("dino", monw/2-(textlen/2), monh/2-250, 200, rl.Fade(rl.White, float32(introtextfade)))
	textlen = rl.MeasureText("9T", 200)
	rl.DrawText("9T", monw/2-(textlen/2)+rInt(-5, 6), monh/2-50+rInt(-5, 6), 200, rl.Fade(randomcolor(), float32(introtextfade)))
	rl.DrawText("9T", monw/2-(textlen/2), monh/2-50, 200, rl.Fade(rl.White, float32(introtextfade)))

	if introtextfade >= 1.0 {
		rl.DrawText("© 2021 nicholasimon", scrolltextx, 20, 40, rl.White)
		rl.DrawText("www.golang.org  |  www.raylib.com", scrolltext2x, monh-60, 40, rl.White)
	}

	if rl.IsKeyPressed(rl.KeySpace) {
		startnewgameon = true

	}

}
func drawstartnewgame() { // MARK: drawstartnewgame
	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, float32(backfade2)))
	if backfade2 >= 1.0 {
		rl.DrawRectangle(monw/2-300, 0, 600, monh, rl.Fade(rl.DarkPurple, 0.8))
		textlen := rl.MeasureText("new game", 80)
		rl.DrawText("new game", (monw/2)-(textlen/2)+rInt(-10, 11), 40+rInt(-10, 11), 80, rl.Black)
		rl.DrawText("new game", (monw/2)-(textlen/2), 40, 80, rl.White)

	}

}
func drawweather() { // MARK: drawweather

	// clouds
	if cloudson && !paused {
		for a := 0; a < len(cloudsl); a++ {
			rl.DrawCircle(int(cloudsv2[a].X), int(cloudsv2[a].Y), float32(cloudsl[a]), rl.Fade(rl.White, rF32(0.3, 0.7)))
			if rolldice() == 6 {
				cloudsv2[a].Y += rFloat32(-4, 5)
			}

			if rolldice() == 6 {
				cloudsv2[a].X += rFloat32(-2, 3)
			}

			if cloudlr {
				cloudsv2[a].X += float32(cloudspeed)
			} else {
				cloudsv2[a].X -= float32(cloudspeed)
			}
		}
	}

	if !options && !paused {

		if snow && rain {
			if flipcoin() {
				snow = true
				rain = false
			} else {
				snow = false
				rain = true
			}

		}

		// snow
		if snow {
			for a := 0; a < len(snowl); a++ {
				origin := rl.NewVector2(float32(0), float32(0))
				destrec := rl.NewRectangle(snowv2[a].X, snowv2[a].Y, float32(snowl[a]), float32(snowl[a]))
				rl.DrawTexturePro(imgs, snowimg[a], destrec, origin, 0, rl.Fade(rl.White, rF32(0.5, 0.9)))

				snowv2[a].Y += rFloat32(6, 13)
				if rolldice() == 6 {
					snowv2[a].Y += rFloat32(-5, 6)
				}
				if snowv2[a].Y > float32(monh) {
					snowv2[a].Y = rFloat32(-15, -5)
					snowv2[a].X = rFloat32(0, monw)
				}
			}
		}
		//rain
		if rain {
			for a := 0; a < len(rainl); a++ {
				rl.DrawCircle(int(rainv2[a].X), int(rainv2[a].Y-3), rainl[a]-1, rl.Fade(randombluelight(), rF32(0.7, 1.0)))
				rl.DrawCircle(int(rainv2[a].X), int(rainv2[a].Y), rainl[a], rl.Fade(randombluelight(), rF32(0.7, 1.0)))

				rainv2[a].Y += rFloat32(28, 31)
				if rolldice() == 6 {
					rainv2[a].X += rFloat32(-5, 6)
				}

				if rainv2[a].Y > float32(monh) {
					rainv2[a].Y = rFloat32(-15, -5)
					rainv2[a].X = rFloat32(0, monw)
				}
			}

		}

	}

}
func drawoptionsmenu() { // MARK: drawoptionsmenu
	rl.DrawRectangle(monw/2-300, 0, 600, monh, rl.Fade(rl.DarkPurple, 0.8))
	textlen := rl.MeasureText("options", 80)
	rl.DrawText("options", (monw/2)-(textlen/2)+rInt(-10, 11), 40+rInt(-10, 11), 80, rl.Black)
	rl.DrawText("options", (monw/2)-(textlen/2), 40, 80, rl.White)

	optiony := 145
	switch optionselect {
	case 0:
		rl.DrawRectangle(monw/2-300, optiony, 600, 50, rl.Fade(randombluedark(), 0.8))
		if rl.IsKeyPressed(rl.KeySpace) {
			if ghosting {
				ghosting = false
			} else {
				ghosting = true
			}
		}
	case 1:
		rl.DrawRectangle(monw/2-300, optiony+50, 600, 50, rl.Fade(randombluedark(), 0.8))
		if rl.IsKeyPressed(rl.KeySpace) {
			if scanlines {
				scanlines = false
			} else {
				scanlines = true
			}
		}
	case 2:
		rl.DrawRectangle(monw/2-300, optiony+100, 600, 50, rl.Fade(randombluedark(), 0.8))
		if rl.IsKeyPressed(rl.KeySpace) {
			if pixelnoise {
				pixelnoise = false
			} else {
				pixelnoise = true
			}
		}
	case 3:
		rl.DrawRectangle(monw/2-300, optiony+150, 600, 50, rl.Fade(randombluedark(), 0.8))
		if rl.IsKeyPressed(rl.KeySpace) {
			if weather {
				weather = false
			} else {
				weather = true
			}
		}
	case 4:
		rl.DrawRectangle(monw/2-300, optiony+200, 600, 50, rl.Fade(randombluedark(), 0.8))
		if rl.IsKeyPressed(rl.KeySpace) {
			if backgon {
				backgon = false
			} else {
				backgon = true
			}
		}
	}

	optiony = 150

	rl.DrawText("ghosting             ", (monw/2)-220, optiony, 40, rl.White)
	rl.DrawRectangle((monw/2)+180, optiony, 40, 40, rl.Black)
	if ghosting {
		rl.DrawRectangle((monw/2)+188, optiony+8, 24, 24, rl.White)
	}
	optiony += 50
	rl.DrawText("scanlines             ", (monw/2)-220, optiony, 40, rl.White)
	rl.DrawRectangle((monw/2)+180, optiony, 40, 40, rl.Black)
	if scanlines {
		rl.DrawRectangle((monw/2)+188, optiony+8, 24, 24, rl.White)
	}
	optiony += 50
	rl.DrawText("pixel noise             ", (monw/2)-220, optiony, 40, rl.White)
	rl.DrawRectangle((monw/2)+180, optiony, 40, 40, rl.Black)
	if pixelnoise {
		rl.DrawRectangle((monw/2)+188, optiony+8, 24, 24, rl.White)
	}
	optiony += 50
	rl.DrawText("weather             ", (monw/2)-220, optiony, 40, rl.White)
	rl.DrawRectangle((monw/2)+180, optiony, 40, 40, rl.Black)
	if weather {
		rl.DrawRectangle((monw/2)+188, optiony+8, 24, 24, rl.White)
	}
	optiony += 50
	rl.DrawText("backgrounds             ", (monw/2)-220, optiony, 40, rl.White)
	rl.DrawRectangle((monw/2)+180, optiony, 40, 40, rl.Black)
	if backgon {
		rl.DrawRectangle((monw/2)+188, optiony+8, 24, 24, rl.White)
	}
	optiony += 50

}
func drawlevelend() { // MARK: drawlevelend

	rl.DrawRectangle(0, 0, monw, monh, rl.Fade(rl.Black, 0.5))

	if rl.IsKeyPressed(rl.KeySpace) {

		newlevel()
	}

}
func drawflying() { // MARK: drawflying

	origin := rl.NewVector2(float32(0), float32(0))
	destrec := rl.NewRectangle(float32(flyingx), float32(flyingy), float32(helicopterl.Width*1.5), float32(helicopterl.Height*1.5))
	rl.DrawTexturePro(imgs, helicopterr, destrec, origin, flyingrotation, rl.White)

	destrec = rl.NewRectangle(float32(flyingx+80), float32(flyingy+27), float32(propellor.Width*1.5), float32(propellor.Height*1.5))
	rl.DrawTexturePro(imgs, propellor, destrec, origin, flyingrotation, rl.White)

	destrec = rl.NewRectangle(float32(flyingx+10), float32(flyingy+50), float32(propellor2.Width*1.5), float32(propellor2.Height*1.5))
	rl.DrawTexturePro(imgs, propellor2, destrec, origin, flyingrotation, rl.White)

	if onoff2 {
		propellor.Y = 586
		propellor2.X = 1328
	} else {
		propellor.Y = 575
		propellor2.X = 1380
	}

	//rl.DrawTextureRec(imgs, helicopterr, v2, rl.White)

	flyingx += flyingspeed

	if rolldice()+rolldice() == 12 && !flyingdrop {
		flyingh, flyingv = (monh-flyingy)/tilesize, flyingx/tilesize
		flyingobjectblok = drawbloknext + ((flyingh * worldw) + flyingv)
		objects[flyingobjectblok].activ = true
		objects[flyingobjectblok].coin = true
		objects[flyingobjectblok].nextblock = flyingobjectblok
		flyingdrop = true
	}
	if rolldice() == 6 {
		flyingy += rInt(-20, 21)
	}
	if rolldice() == 6 {
		flyingrotation += rFloat32(-2, 3)
	}
	if flyingrotation > 10.0 {
		flyingrotation -= 2.0
	}
	if flyingrotation < -10.0 {
		flyingrotation += 2.0
	}
	if flyingx > monw+10 {
		flyingon = false
		flyingdrop = false
	}

}
func drawpickuptext() { // MARK: drawpickuptext

	rl.DrawText(currentpickuptext, player.x-10, player.y-pickuptextychange, 40, rl.White)

}
func newlevel() { // MARK: newlevel

	//lineson = flipcoin()
	levelnumber++
	gametime = 90
	weathertimer = rInt(30, 60)
	rain = flipcoin()

	for a := 0; a < len(world); a++ {
		world[a] = blok{}
		objects[a] = blok{}
	}
	for a := 0; a < len(enemies); a++ {
		enemies[a] = enemyblok{}
	}

	createmap()
	levelon = true
	levelend = false
	paused = false

}
func createmap() { // MARK: createmap
	treetype = rInt(0, 3)
	chooseimg := tiles[rInt(0, len(tiles))]

	//floor
	for a := 0; a < worldw*20; a++ {
		world[a].activ = true
		world[a].ground = true
		world[a].color1 = randomgreen()
		world[a].solid = true
		world[a].img = chooseimg

		if a > worldw*19 {
			if rolldice()+rolldice() == 12 {

				length := rInt(1, 6)
				newblok := a
				newblok += worldw

				for {

					for b := 0; b < length; b++ {
						world[newblok].activ = true
						world[newblok].ground = true
						world[newblok].color1 = randomgreen()
						world[newblok].solid = true
						world[newblok].img = chooseimg
						newblok++
					}
					newblok -= length
					newblok++
					newblok += worldw
					length--

					if length <= 0 {
						break
					}
				}

			}
		}
	}
	//trees
	for a := worldw * 20; a < worldw*21; a++ {
		if rolldice() == 6 {
			world[a].activ = true
			world[a].imgon = true
			switch treetype {
			case 0:
				world[a].img = trees[rInt(0, len(trees))]
			case 1:
				world[a].img = trees2[rInt(0, len(trees2))]
			case 2:
				world[a].img = trees3[rInt(0, len(trees3))]
			}
		}
	}
	createplatforms()
	//back tiles
	backtile1 = tiles[rInt(0, len(tiles))]
	backtile2 = tiles[rInt(0, len(tiles))]
	backcolor1 = randomcolor()
	backcolor2 = randomcolor()

	for a := 0; a < len(backtiles); a++ {
		backtiles[a] = backtile1
		if rolldice() > 4 {
			backtiles[a] = backtile2
		}
	}

	for a := 0; a < enemiesnumber; a++ {

		choose := rInt(worldw*20, worlda-(worldw*20))
		for {
			choose = rInt(worldw*20, worlda-(worldw*20))
			if !world[choose].solid {
				break
			}
		}
		enemies[choose].nextblock = choose
		enemies[choose].activ = true
		if flipcoin() {
			enemies[player.bloknumber+4].direction = 4
		} else {
			enemies[player.bloknumber+4].direction = 6
		}
		switch rolldice() {
		case 1:
			enemies[choose].name = "spikes"
			enemies[choose].hp = 5
		case 2:
			enemies[choose].name = "bunny"
			enemies[choose].hp = 8
		case 3:
			enemies[choose].name = "chicken"
			enemies[choose].hp = 7
		case 4:
			enemies[choose].name = "ghost"
			enemies[choose].hp = 2
		case 5:
			enemies[choose].name = "mushroom"
			enemies[choose].hp = 2
		case 6:
			enemies[choose].name = "raddish"
			enemies[choose].hp = 3
		}

	}

	for a := 0; a < monsternumber; a++ {

		choose := rInt(worldw*20, worlda-(worldw*20))
		for {
			choose = rInt(worldw*20, worlda-(worldw*20))
			if !world[choose].solid {
				break
			}
		}

		monsters[choose].activ = true
		switch rInt(1, 26) {
		case 1:
			monsters[choose].name = "monster1"
		case 2:
			monsters[choose].name = "monster2"
		case 3:
			monsters[choose].name = "monster3"
		case 4:
			monsters[choose].name = "monster4"
		case 5:
			monsters[choose].name = "monster5"
		case 6:
			monsters[choose].name = "monster6"
		case 7:
			monsters[choose].name = "monster7"
		case 8:
			monsters[choose].name = "monster8"
		case 9:
			monsters[choose].name = "monster9"
		case 10:
			monsters[choose].name = "monster10"
		case 11:
			monsters[choose].name = "monster11"
		case 12:
			monsters[choose].name = "monster12"
		case 13:
			monsters[choose].name = "monster13"
		case 14:
			monsters[choose].name = "monster14"
		case 15:
			monsters[choose].name = "monster15"
		case 16:
			monsters[choose].name = "monster16"
		case 17:
			monsters[choose].name = "monster17"
		case 18:
			monsters[choose].name = "monster18"
		case 19:
			monsters[choose].name = "monster19"
		case 20:
			monsters[choose].name = "monster20"
		case 21:
			monsters[choose].name = "monster21"
		case 22:
			monsters[choose].name = "monster22"
		case 23:
			monsters[choose].name = "monster23"
		case 24:
			monsters[choose].name = "monster24"
		case 25:
			monsters[choose].name = "monster25"
		}

	}

	for {
		portalblock = rInt(worldw*20, worlda-(worldw*20))
		if !world[portalblock].solid && !world[portalblock-worldw].solid && world[portalblock-(worldw*2)].solid {
			break
		}
	}

	world[portalblock].activ = true
	world[portalblock].portal = true

	levelon = true
}
func createplatforms() { // MARK: createplatforms

	platblok := worldw * 23

	currentplath := platblok / worldw

	plattile1 = tiles[rInt(0, len(tiles))]
	plattile2 = tiles[rInt(0, len(tiles))]
	plattile3 = tiles[rInt(0, len(tiles))]

	plat1color = randomorange()
	plat2color = brightred()
	plat3color = brightyellow()

	for {

		if rolldice()+rolldice() == 12 {

			platlen := rInt(4, 15)
			platcolor := plat1color
			platimg := plattile1
			if flipcoin() {
				platimg = plattile2
			}
			if flipcoin() {
				platcolor = plat2color
			}

			//draw platform
			for a := 0; a < platlen; a++ {
				if rolldice()+rolldice() == 12 {
					world[platblok].activ = true
					world[platblok].solid = true
					world[platblok].powerup = true
					world[platblok].color1 = brightyellow()
					world[platblok].img = powerupblokimg
				} else {
					world[platblok].activ = true
					world[platblok].solid = true
					world[platblok].color1 = platcolor
					world[platblok].img = platimg
				}
				platblok++
			}
			platblok += rInt(8, 24)
		}

		//draw shape
		if rolldice()+rolldice()+rolldice() > 16 {

			choose := rInt(1, 4)

			shapeblok := platblok
			switch choose {
			case 3: // solid block

				length := rInt(3, 11)
				area := length * length
				count := 0
				newcolor := randomcolor()
				for a := 0; a < area; a++ {
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok++
					count++

					if count == length {
						count = 0
						shapeblok -= length
						shapeblok += worldw
					}

				}

			case 2: // pyramid steps

				number := rInt(3, 6)
				orignumber := number
				newcolor := randomcolor()

				for {
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok++
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok += worldw
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3

					number--
					if number <= 0 {
						break
					}
				}
				number = orignumber
				for {
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok++
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok -= worldw
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3

					number--
					if number <= 0 {
						shapeblok++
						world[shapeblok].activ = true
						world[shapeblok].solid = true
						world[shapeblok].color1 = newcolor
						world[shapeblok].img = plattile3
						break
					}
				}

			case 1: // steps
				number := rInt(3, 11)
				newcolor := randomcolor()

				for {

					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok++
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3
					shapeblok += worldw
					world[shapeblok].activ = true
					world[shapeblok].solid = true
					world[shapeblok].color1 = newcolor
					world[shapeblok].img = plattile3

					number--
					if number <= 0 {
						break
					}
				}

			}

		}

		platblok++

		plath := platblok / worldw

		if plath > currentplath {
			platblok += rInt(1, 3) * worldw
			currentplath = platblok / worldw
			plath = platblok / worldw

		}

		if platblok > worlda-(worldw*50) {
			break
		}
	}

}
func createcloud() { // MARK: createcloud

	cloudspeed = rInt(10, 25)
	cloudlr = flipcoin()
	sizemap := rInt(50, 250)
	cloudsl = make([]int, sizemap)
	cloudsv2 = make([]rl.Vector2, sizemap)

	startv2 := rl.NewVector2(float32(monw)+500, rFloat32(-100, 100))
	if cloudlr {
		startv2.X = -500
	}

	length := rInt(20, 40)

	origstartv2 := startv2
	count := 0
	for a := 0; a < len(cloudsl); a++ {

		cloudsl[a] = rInt(15, 35)
		cloudsv2[a] = startv2
		startv2.X += rFloat32(15, 25)
		if rolldice() == 6 {
			cloudsv2[a].Y += rFloat32(-2, 3)
		}
		if rolldice() == 6 {
			cloudsv2[a].Y -= rFloat32(10, 30)
		}
		if rolldice() == 6 {
			cloudsv2[a].Y += rFloat32(10, 30)
		}
		count++
		if count == length {
			count = 0
			startv2.X = origstartv2.X
			startv2.Y += rFloat32(20, 40)
			change := rInt(20, 40)
			startv2.X -= float32((change / 2) * rInt(-4, 5))
			if rolldice() == 6 {
				startv2.X -= rFloat32(10, 30)
			}
			length += rInt(-10, 11)

		}

	}

}

func geth(blocknumber int) int { // MARK: geth
	h := blocknumber / worldw
	return h
}
func getv(blocknumber int) int { // MARK: getv
	v := blocknumber % worldw
	return v
}
func checkhv(blocknumber int) bool { // MARK: checkhv

	end := false
	h, v := blocknumber/worldw, blocknumber%worldw

	if h > worlda-(worldw*10) {
		end = true
	}
	if h < (worldw * 10) {
		end = true
	}
	if v > worldw-20 {
		end = true
	}
	if v < 20 {
		end = true
	}

	return end

}
func updatehv() { // MARK: updatehv

	portalh, portalv = portalblock/worldh, portalblock%worldw
	player.h, player.v = player.bloknumber/worldh, player.bloknumber%worldw

	vdiff := math.Abs(float64(portalv) - float64(player.v))
	hdiff := math.Abs(float64(portalh) - float64(player.h))

	if portalh > player.h && portalv < player.v && vdiff > 4 {
		portaldirection = 7
	} else if portalh > player.h && vdiff < 5 {
		portaldirection = 8
	} else if portalh > player.h && portalv > player.v && vdiff > 4 {
		portaldirection = 9
	} else if hdiff < 5 && portalv > player.v {
		portaldirection = 6
	} else if hdiff < 5 && portalv < player.v {
		portaldirection = 4
	} else if portalh < player.h && portalv > player.v && vdiff > 4 {
		portaldirection = 3
	} else if portalh < player.h && portalv < player.v && vdiff > 4 {
		portaldirection = 1
	} else if portalh < player.h && vdiff < 5 {
		portaldirection = 2
	}

}

// MARK: core	core	core	core	core	core	core	core	core	core	core
func main() { // MARK: main
	rand.Seed(time.Now().UnixNano()) // random numbers
	rl.SetTraceLogLevel(rl.LogError) // hides info window
	rl.InitWindow(monw, monh, "setres")
	setres(0, 0)
	rl.CloseWindow()
	setinitialvalues()
	raylib()

}
func input() { // MARK: input

	if rl.IsKeyPressed(rl.KeyEscape) {
		if options {
			options = false
			paused = false
		} else {
			options = true
			paused = true
			optionselect = 0
		}

	}

	if rl.IsKeyPressed(rl.KeyPause) {
		if paused {
			paused = false
		} else {
			paused = true
		}
	}

	if freemoveon && !options {
		if rl.IsKeyDown(rl.KeyUp) {
			drawbloknext += worldw
		}
		if rl.IsKeyDown(rl.KeyDown) {
			drawbloknext -= worldw
		}
		if rl.IsKeyDown(rl.KeyRight) {
			drawbloknext++
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			drawbloknext--
		}

	} else if options {
		if rl.IsKeyPressed(rl.KeyUp) {
			optionselect--
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			optionselect++
		}
	} else {
		if rl.IsKeyPressed(rl.KeySpace) {
			playeraction()

		}
		if rl.IsKeyDown(rl.KeyUp) {
			jumph = 5
			jumpon = true
		}
		if rl.IsKeyDown(rl.KeyDown) {
			drawbloknext -= worldw
		}
		if rl.IsKeyDown(rl.KeyRight) {
			if !world[player.bloknumber+1].solid {
				player.bloknumber++
			}
			player.direction = false
			player.moving = true
		} else if rl.IsKeyReleased(rl.KeyRight) {
			player.moving = false
		}

		if rl.IsKeyDown(rl.KeyLeft) {
			if !world[player.bloknumber-1].solid {
				player.bloknumber--
			}
			player.direction = true
			player.moving = true
		} else if rl.IsKeyReleased(rl.KeyLeft) {
			player.moving = false
		}

	}

	// DEV KEYS DELETE
	if rl.IsKeyPressed(rl.KeyF3) {
		if introon {
			introon = false
		} else {
			introon = true
		}

	}
	if rl.IsKeyPressed(rl.KeyF2) {
		levelend = true
	}
	if rl.IsKeyPressed(rl.KeyF1) {
		if freemoveon {
			freemoveon = false
		} else {
			freemoveon = true
		}
	}

	// DEV KEYS DELETE

	if rl.IsKeyPressed(rl.KeyKpAdd) {
		camera.Zoom += 0.2
	}
	if rl.IsKeyPressed(rl.KeyKpSubtract) {
		camera.Zoom -= 0.2
	}

	if rl.IsKeyPressed(rl.KeyKpDecimal) {
		if debugon {
			debugon = false
		} else {
			debugon = true
		}
	}

	if rl.IsKeyPressed(rl.KeyKp0) {
		if gridon {
			gridon = false
		} else {
			gridon = true
		}
	}
}
func createimgs() { // MARK: createimgs

	x := float32(0)
	y := float32(0)
	count := 0
	for a := 0; a < len(tiles); a++ {

		tiles[a] = rl.NewRectangle(x, y, 16, 16)
		x += 16
		count++
		if count == 5 {
			count = 0
			x = 0
			y += 16
		}
	}

	x = float32(193)
	y = float32(0)
	count = 0
	for a := 0; a < len(trees); a++ {
		trees[a] = rl.NewRectangle(x, y, 32, 32)
		x += 32
		count++
		if count == 4 {
			count = 0
			x = 193
			y += 32
		}
	}

	x = float32(214)
	y = float32(151)
	count = 0
	for a := 0; a < len(trees2); a++ {
		trees2[a] = rl.NewRectangle(x, y, 32, 32)
		x += 32
		count++
		if count == 4 {
			count = 0
			x = 214
			y += 32
		}
	}

	x = float32(390)
	y = float32(134)
	count = 0
	for a := 0; a < len(trees3); a++ {
		trees3[a] = rl.NewRectangle(x, y, 32, 32)
		x += 32
		count++
		if count == 6 {
			count = 0
			x = 390
			y += 32
		}
	}

	x = float32(640)
	y = float32(847)

	for a := 0; a < len(emotes); a++ {
		emotes[a] = rl.NewRectangle(x, y, 32, 32)
		x += 96
		if x > 1536 {
			x = 640
			y += 32
		}

	}

}

func drawdebug() { // MARK: DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG DEBUG

	//centerlines
	rl.DrawLine(monw/2, 0, monw/2, monh, rl.Magenta)
	rl.DrawLine(0, monh/2, monw, monh/2, rl.Magenta)

	rl.DrawRectangle(monw-300, 0, 300, monh, rl.Fade(rl.Black, 0.8))
	textx := monw - 290
	textx2 := monw - 145
	texty := 10

	camerazoomtext := fmt.Sprintf("%g", camera.Zoom)
	worldwtext := strconv.Itoa(worldw)
	drawatext := strconv.Itoa(drawa)
	drawwtext := strconv.Itoa(draww)
	drawhtext := strconv.Itoa(drawh)
	playervtext := strconv.Itoa(player.v)
	playerhtext := strconv.Itoa(player.h)
	drawblocknextextendtext := strconv.Itoa(drawblocknextextend)
	drawbloknexttext := strconv.Itoa(drawbloknext)
	playermovingtext := strconv.FormatBool(player.moving)
	playerisholdingtext := strconv.FormatBool(player.isholding)
	playerbloknumbertext := strconv.Itoa(player.bloknumber)

	rl.DrawText("drawa", textx, texty, 10, rl.White)
	rl.DrawText(drawatext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("draww", textx, texty, 10, rl.White)
	rl.DrawText(drawwtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("drawh", textx, texty, 10, rl.White)
	rl.DrawText(drawhtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("camerazoomtext", textx, texty, 10, rl.White)
	rl.DrawText(camerazoomtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("worldwtext", textx, texty, 10, rl.White)
	rl.DrawText(worldwtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("playervtext", textx, texty, 10, rl.White)
	rl.DrawText(playervtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("playerhtext", textx, texty, 10, rl.White)
	rl.DrawText(playerhtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("drawblocknextextend", textx, texty, 10, rl.White)
	rl.DrawText(drawblocknextextendtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("drawbloknext", textx, texty, 10, rl.White)
	rl.DrawText(drawbloknexttext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("playermovingtext", textx, texty, 10, rl.White)
	rl.DrawText(playermovingtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("playerisholdingtext", textx, texty, 10, rl.White)
	rl.DrawText(playerisholdingtext, textx2, texty, 10, rl.White)
	texty += 12
	rl.DrawText("playerbloknumbertext", textx, texty, 10, rl.White)
	rl.DrawText(playerbloknumbertext, textx2, texty, 10, rl.White)
	texty += 12

	// fps
	rl.DrawRectangle(monw-110, monh-110, 100, 40, rl.Black)
	rl.DrawFPS(monw-100, monh-100)

}
func timers() { // MARK: timers

	if destroytimer != 0 {
		destroytimercount++
		if destroytimercount%30 == 0 {
			destroytimer--
		}
	} else if destroytimer <= 0 {
		destroyon = false
	}

	if player.moving {
		if player.direction {
			dinol.X += 24
			if dinol.X > 318 {
				dinol.X = 0
			}
		} else {
			dinor.X -= 24
			if dinor.X < 0 {
				dinor.X = 314
			}
		}
	} else {
		if onoff6 {
			if player.direction {
				dinol.X += 24
				if dinol.X > 56 {
					dinol.X = 0
				}
			} else {
				dinor.X -= 24
				if dinor.X < 260 {
					dinor.X = 314
				}
			}
		}

	}

	if onoff2 {
		monster25.X += 16
		if monster25.X > 448 {
			monster25.X = 380
		}
		monster24.X += 16
		if monster24.X > 548 {
			monster24.X = 481
		}
		monster23.X += 16
		if monster23.X > 448 {
			monster23.X = 382
		}
		monster22.X += 16
		if monster22.X > 464 {
			monster22.X = 382
		}
		monster21.X += 16
		if monster21.X > 552 {
			monster21.X = 501
		}
		monster20.X += 16
		if monster20.X > 465 {
			monster20.X = 382
		}
		monster19.X += 16
		if monster19.X > 552 {
			monster19.X = 501
		}
		monster18.X += 16
		if monster18.X > 464 {
			monster18.X = 383
		}
		monster17.X += 16
		if monster17.X > 533 {
			monster17.X = 483
		}
		monster16.X += 16
		if monster16.X > 448 {
			monster16.X = 383
		}
		monster15.X += 16
		if monster15.X > 536 {
			monster15.X = 502
		}
		monster14.X += 16
		if monster14.X > 466 {
			monster14.X = 384
		}
		monster13.X += 16
		if monster13.X > 520 {
			monster13.X = 502
		}
		monster12.X += 16
		if monster12.X > 464 {
			monster12.X = 382
		}
		monster11.X += 16
		if monster11.X > 518 {
			monster11.X = 468
		}
		monster10.X += 16
		if monster10.X > 448 {
			monster10.X = 382
		}
		monster9.X += 16
		if monster9.X > 510 {
			monster9.X = 460
		}
		monster8.X += 16
		if monster8.X > 434 {
			monster8.X = 382
		}
		monster7.X += 16
		if monster7.X > 512 {
			monster7.X = 462
		}
		monster6.X += 16
		if monster6.X > 432 {
			monster6.X = 382
		}
		monster5.X += 16
		if monster5.X > 514 {
			monster5.X = 463
		}
		monster4.X += 16
		if monster4.X > 432 {
			monster4.X = 382
		}
		monster3.X += 16
		if monster3.X > 572 {
			monster3.X = 523
		}
		monster1.X += 16
		if monster1.X > 434 {
			monster1.X = 383
		}
		monster2.X += 16
		if monster2.X > 496 {
			monster2.X = 463
		}
	}

	if cloudtimer > 0 {
		cloudtimercount++
		if cloudtimercount%30 == 0 {
			cloudtimer--
		}
	} else if cloudtimer <= 0 {

		if cloudson {
			cloudson = false
			cloudtimer = rInt(2, 5)
		} else {
			for a := 0; a < len(cloudsl); a++ {
				cloudsl[a] = 0
				cloudsv2[a] = rl.NewVector2(0, 0)
			}

			cloudtimer = rInt(10, 15)
			createcloud()
			cloudson = true
		}

	}

	if emotetimer != 0 {
		emotetimercount++
		if emotetimercount%30 == 0 {
			emotetimer--
		}
	} else if emotetimer <= 0 {
		if emoteon {
			emoteon = false
			emotetimercount = 0
			emotetimer = rInt(20, 40)
		} else {
			emotetimercount = 0
			emotetimer = rInt(5, 11)
			currentemote = emotes[rInt(0, len(emotes))]
			emotex = currentemote.X
			emoteon = true
		}
	}
	if emoteon {
		if onoff2 {
			currentemote.X += 32
			emotecount++
			if emotecount == 3 {
				emotecount = 0
				currentemote.X = emotex
			}
		}
	}

	if levelon {
		gametimecount++

		if gametimecount%30 == 0 {
			gametime--
		}
		if gametime <= 0 {
			levelend = true
		}

	}

	if weather {
		if rain || snow {
			weathertimercount++
			if weathertimercount%30 == 0 {
				weathertimer--
			}
			if weathertimer <= 0 {
				weathertimer = rInt(30, 60)
				weathertimercount = 0
				if rain {
					rain = false
				} else if snow {
					snow = false
				}
			}

		} else {
			weathertimercount++
			if weathertimercount%30 == 0 {
				weathertimer--
			}
			if weathertimer <= 0 {
				weathertimer = rInt(30, 60)
				weathertimercount = 0
				if flipcoin() {
					rain = true
				} else {
					snow = true
				}
			}

		}
	}
	if onoff2 {
		raddishimg.X += 30
		if raddishimg.X > 162 {
			raddishimg.X = 9
		}
		raddishlimg.X -= 30
		if raddishlimg.X < 0 {
			raddishlimg.X = 160
		}
	}
	ghostimg.X += 44
	if ghostimg.X > 414 {
		ghostimg.X = 6
	}
	ghostlimg.X -= 44
	if ghostlimg.X < 0 {
		ghostlimg.X = 403
	}
	chickenimg.X += 32
	if chickenimg.X > 424 {
		chickenimg.X = 3
	}
	chickenlimg.X -= 32
	if chickenlimg.X < 0 {
		chickenlimg.X = 422
	}
	if onoff3 {
		spikesimg.X += 44
		if spikesimg.X > 320 {
			spikesimg.X = 13
		}
		spikeslimg.X -= 44
		if spikeslimg.X < 0 {
			spikeslimg.X = 321
		}
	}
	mushroomimg.X += 32
	if mushroomimg.X > 490 {
		mushroomimg.X = 5
	}
	mushroomlimg.X -= 32
	if mushroomlimg.X < 0 {
		mushroomlimg.X = 488
	}
	bunnyimg.X += 34
	if bunnyimg.X > 382 {
		bunnyimg.X = 2
	}
	bunnylimg.X -= 34
	if bunnylimg.X < 0 {
		bunnylimg.X = 378
	}

	if backgzoomon {
		camerabackg.Zoom -= 0.01
		if camerabackg.Zoom <= 1.0 {
			backgzoomon = false
		}
	} else {
		camerabackg.Zoom += 0.01
		if camerabackg.Zoom >= 1.2 {
			backgzoomon = true
		}
	}

	if pickuptexton {
		drawpickuptext()
		pickuptextychange += 10
		pickuptexttimercount++
		if pickuptexttimercount%fps == 0 {
			pickuptexttimer--
			if pickuptexttimer <= 0 {
				pickuptexttimercount = 0
				pickuptextychange = 70
				pickuptexton = false
			}
		}
	}

	if onoff3 {
		coinimg.X += 16
		if coinimg.X > 70 {
			coinimg.X = 0
		}
	}

	if flyingtimer > 0 && !flyingon {
		flyingtimercount++
		if flyingtimercount%30 == 0 {
			flyingtimer--
		}
	} else if flyingtimer <= 0 && !flyingon {
		flyingy = rInt(20, 200)
		flyingx = -250
		flyingspeed = rInt(15, 25)
		flyingon = true
		flyingtimercount = 0
		flyingtimer = 0
	}

	if framecount%2 == 0 {
		if onoff2 {
			onoff2 = false
		} else {
			onoff2 = true
		}
	}
	if framecount%3 == 0 {
		if onoff3 {
			onoff3 = false
		} else {
			onoff3 = true
		}
	}
	if framecount%6 == 0 {
		if onoff6 {
			onoff6 = false
		} else {
			onoff6 = true
		}
	}
	if framecount%10 == 0 {
		if onoff10 {
			onoff10 = false
		} else {
			onoff10 = true
		}
	}
	if framecount%15 == 0 {
		if onoff15 {
			onoff15 = false
		} else {
			onoff15 = true
		}
	}
	if framecount%30 == 0 {
		if onoff30 {
			onoff30 = false
		} else {
			onoff30 = true
		}
	}
	if framecount%60 == 0 {
		if onoff60 {
			onoff60 = false
		} else {
			onoff60 = true
		}
	}
	if fadeblinkon {
		if fadeblink > 0.2 {
			fadeblink -= 0.05
		} else {
			fadeblinkon = false
		}
	} else {
		if fadeblink < 0.6 {
			fadeblink += 0.05
		} else {
			fadeblinkon = true
		}
	}
	if onoff3 {
		if fadeblink2on {
			if fadeblink2 > 0.1 {
				fadeblink2 -= 0.01
			} else {
				fadeblink2on = false
			}
		} else {
			if fadeblink2 < 0.2 {
				fadeblink2 += 0.01
			} else {
				fadeblink2on = true
			}
		}
	}
}

func setres(w, h int) { // MARK: setres

	if w == 0 {

		monw = rl.GetMonitorWidth(0)
		monh = rl.GetMonitorHeight(0)
		camera.Zoom = 1.0
		camerabackg.Zoom = 1.0

		if monw >= 1600 {
			tilesize = 96
			textsize1 = 80
		} else if monw < 1600 {
			tilesize = 72
			textsize1 = 40
		}

	} else {
		monw = w
		monh = h
		camera.Zoom = 1.0
		camerabackg.Zoom = 1.0

		if monw >= 1600 {
			tilesize = 96
			textsize1 = 80
		} else if monw < 1600 {
			tilesize = 72
			textsize1 = 40
		}
	}

}

func setinitialvalues() { // MARK: setinitialvalues

	player.isholding = true
	player.holding = "uzzi"
	player.objimg = uzziimg
	player.objlimg = uzzilimg

	emotetimer = rInt(20, 40)

	for a := 0; a < len(dinomultiplier); a++ {
		dinomultiplier[a] = rInt(2, 6)
		dinotype[a] = rInt(1, 5)
	}

	cloudson = true
	cloudtimer = rInt(10, 20)
	createcloud()

	emotetimer = rInt(3, 5)
	backgon = true
	lineson = flipcoin()
	player.hp = 3
	weather = true
	levelnumber = 1
	gametime = 90
	weathertimer = rInt(30, 60)
	rain = flipcoin()
	freemoveon = false
	flyingtimer = rInt(2, 5)
	//backgon = true
	ghosting = true
	scanlines = true
	pixelnoise = true

	draww = (monw / tilesize) + 1
	drawh = (monh / tilesize) + 1
	drawa = draww * drawh

	drawwextend = draww * 3
	drawhextend = drawh * 3
	drawaextend = drawwextend * drawhextend

	backtiles = make([]rl.Rectangle, drawa)

	player.bloknumber = (worldw * 21) + 10
	drawbloknext = player.bloknumber
	drawbloknext -= draww / 2
	drawbloknext -= ((drawh / 3) * 2) * worldw

	for a := 0; a < len(background); a++ {
		background[a].v2.X = float32(rInt(0, monw))
		background[a].v2.Y = float32(rInt(0, monw))
		background[a].w = rInt(20, 40)
		background[a].h = rInt(20, 40)
		background[a].color1 = randomorange()
		background[a].color2 = randomcolor()
		background[a].opac = rF32(0.4, 0.9)
	}

	for a := 0; a < enemiesnumber; a++ {

		choose := rInt(worldw*20, worlda-(worldw*20))
		for {
			choose = rInt(worldw*20, worlda-(worldw*20))
			if !world[choose].solid {
				break
			}
		}
		enemies[choose].nextblock = choose
		enemies[choose].activ = true
		if flipcoin() {
			enemies[player.bloknumber+4].direction = 4
		} else {
			enemies[player.bloknumber+4].direction = 6
		}
		switch rolldice() {
		case 1:
			enemies[choose].name = "spikes"
		case 2:
			enemies[choose].name = "bunny"
		case 3:
			enemies[choose].name = "chicken"
		case 4:
			enemies[choose].name = "ghost"
		case 5:
			enemies[choose].name = "mushroom"
		case 6:
			enemies[choose].name = "raddish"
		}

	}

	if flipcoin() {
		enemies[player.bloknumber+4].direction = 4
	} else {
		enemies[player.bloknumber+4].direction = 6
	}

	// create rain
	for a := 0; a < len(rainl); a++ {
		rainl[a] = rFloat32(2, 5)
		rainv2[a].X = rFloat32(0, monw)
		rainv2[a].Y = rFloat32(0, monh)
	}
	// create snow
	for a := 0; a < len(snowl); a++ {
		snowl[a] = rInt(14, 49)
		snowv2[a].X = rFloat32(0, monw)
		snowv2[a].Y = rFloat32(0, monh)
		switch rInt(1, 5) {
		case 1:
			snowimg[a] = snowflake1
		case 2:
			snowimg[a] = snowflake2
		case 3:
			snowimg[a] = snowflake3
		case 4:
			snowimg[a] = snowflake4
		}
	}

}

// MARK:  █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █ █
func drawgrid() { // MARK: drawgrid

	x := 16
	for {
		rl.DrawLine(x, 0, x, monh, rl.Fade(rl.Magenta, 0.1))
		x += 16
		if x > monw {
			break
		}
	}
	y := 16
	for {
		rl.DrawLine(0, y, monw, y, rl.Fade(rl.Magenta, 0.1))
		y += 16
		if y > monh {
			break
		}
	}

}

// MARK: colors
// https://www.rapidtables.com/web/color/RGB_Color.html
func darkred() rl.Color {
	color := rl.NewColor(55, 0, 0, 255)
	return color
}
func semidarkred() rl.Color {
	color := rl.NewColor(70, 0, 0, 255)
	return color
}
func brightred() rl.Color {
	color := rl.NewColor(230, 0, 0, 255)
	return color
}
func randomgrey() rl.Color {
	color := rl.NewColor(uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(160, 193)), uint8(rInt(0, 255)))
	return color
}
func randombluelight() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 180)), uint8(rInt(120, 256)), uint8(rInt(120, 256)), 255)
	return color
}
func randombluedark() rl.Color {
	color := rl.NewColor(0, 0, uint8(rInt(120, 250)), 255)
	return color
}
func randomyellow() rl.Color {
	color := rl.NewColor(255, uint8(rInt(150, 256)), 0, 255)
	return color
}
func randomorange() rl.Color {
	color := rl.NewColor(uint8(rInt(250, 256)), uint8(rInt(60, 210)), 0, 255)
	return color
}
func randomred() rl.Color {
	color := rl.NewColor(uint8(rInt(128, 256)), uint8(rInt(0, 129)), uint8(rInt(0, 129)), 255)
	return color
}
func randomgreen() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 170)), uint8(rInt(100, 256)), uint8(rInt(0, 50)), 255)
	return color
}
func randomcolor() rl.Color {
	color := rl.NewColor(uint8(rInt(0, 256)), uint8(rInt(0, 256)), uint8(rInt(0, 256)), 255)
	return color
}
func brightyellow() rl.Color {
	color := rl.NewColor(uint8(255), uint8(255), uint8(0), 255)
	return color
}
func brightbrown() rl.Color {
	color := rl.NewColor(uint8(218), uint8(165), uint8(32), 255)
	return color
}
func brightgrey() rl.Color {
	color := rl.NewColor(uint8(212), uint8(212), uint8(213), 255)
	return color
}

// random numbers
func rF32(min, max float32) float32 {
	return (rand.Float32() * (max - min)) + min
}
func rInt(min, max int) int {
	return rand.Intn(max-min) + min
}
func rInt32(min, max int) int32 {
	a := int32(rand.Intn(max-min) + min)
	return a
}
func rFloat32(min, max int) float32 {
	a := float32(rand.Intn(max-min) + min)
	return a
}
func flipcoin() bool {
	var b bool
	a := rInt(0, 10001)
	if a < 5000 {
		b = true
	}
	return b
}
func rolldice() int {
	a := rInt(1, 7)
	return a
}
