package main

import (
	"image"

	//needed because of ebiten
	_ "image/png"
	"log"

	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type direction int

const (
	up direction = iota
	down
	left
	right
	noDirection
)

// Position bla
type Position struct {
	x int
	y int
}

// Character bla
type Character struct {
	position      Position
	dir           direction
	counter       int
	walkSprites   map[direction][]*ebiten.Image
	shieldSprites map[direction][]*ebiten.Image
	attackSprites map[direction][]*ebiten.Image
	lastSprite    *ebiten.Image
	framesPerStep int
}

// Walk bla
func (character *Character) Walk(dir direction) {

}

// Shield bla
func (character *Character) Shield(dir direction) {

}

// Attack bla
func (character *Character) Attack(dir direction) {

}

// Draw bla
func (character *Character) Draw(screen *ebiten.Image) {
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(character.position.x), float64(character.position.y))
	var sprite *ebiten.Image
	if character.lastSprite == nil {
		sprite = character.walkSprites[character.dir][(character.counter/10)%2]
	} else {
		sprite = character.lastSprite
	}

	screen.DrawImage(sprite, &options)
}

// Move bla
func (character *Character) Move(dx, dy int) {
	character.position.x += dx
	character.position.y += dy
}

//Update bla
func (character *Character) Update(dir direction) {
	if dir != noDirection {
		character.dir = dir
		character.counter++
		switch dir {
		case up:
			character.Move(0, -1)
		case down:
			character.Move(0, 1)
		case left:
			character.Move(-1, 0)
		case right:
			character.Move(1, 0)
		}
	}
}

func imgFromFile(filename string) image.Image {
	file, err := ebitenutil.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	img, format, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Format:", format)
	return img
}

func getWalkingSprites() map[direction][]*image.NRGBA {
	walkSprites := make(map[direction][]*image.NRGBA)
	spriteImage := imgFromFile("sprites.png").(*image.NRGBA)

	walkLeft1 := spriteImage.SubImage(image.Rect(35, 11, 51, 27)).(*image.NRGBA)
	walkLeft2 := spriteImage.SubImage(image.Rect(52, 11, 68, 27)).(*image.NRGBA)

	walkRight1 := imaging.FlipH(walkLeft1)
	walkRight2 := imaging.FlipH(walkLeft2)

	walkUp1 := spriteImage.SubImage(image.Rect(18, 11, 34, 27)).(*image.NRGBA)
	walkUp2 := imaging.FlipH(walkUp1)

	walkDown1 := spriteImage.SubImage(image.Rect(1, 11, 17, 27)).(*image.NRGBA)
	walkDown2 := imaging.FlipH(walkDown1)

	walkSprites[left] = []*image.NRGBA{walkLeft1, walkLeft2}
	walkSprites[right] = []*image.NRGBA{walkRight1, walkRight2}
	walkSprites[up] = []*image.NRGBA{walkUp1, walkUp2}
	walkSprites[down] = []*image.NRGBA{walkDown1, walkDown2}

	return walkSprites
}

func getShieldSprites() map[direction][]*image.NRGBA {
	spriteImage := imgFromFile("sprites.png").(*image.NRGBA)
	shieldSprites := make(map[direction][]*image.NRGBA)

	shieldRight1 := spriteImage.SubImage(image.Rect(69, 42, 85, 58)).(*image.NRGBA)
	shieldRight2 := spriteImage.SubImage(image.Rect(86, 42, 102, 58)).(*image.NRGBA)

	shieldLeft1 := spriteImage.SubImage(image.Rect(35, 11, 51, 27)).(*image.NRGBA)
	shieldLeft2 := spriteImage.SubImage(image.Rect(52, 11, 68, 27)).(*image.NRGBA)

	shieldUp1 := spriteImage.SubImage(image.Rect(35, 42, 51, 58)).(*image.NRGBA)
	shieldUp2 := spriteImage.SubImage(image.Rect(52, 42, 68, 58)).(*image.NRGBA)

	shieldDown1 := spriteImage.SubImage(image.Rect(1, 42, 17, 58)).(*image.NRGBA)
	shieldDown2 := spriteImage.SubImage(image.Rect(18, 42, 34, 58)).(*image.NRGBA)

	shieldSprites[left] = []*image.NRGBA{shieldLeft1, shieldLeft2}
	shieldSprites[right] = []*image.NRGBA{shieldRight1, shieldRight2}
	shieldSprites[up] = []*image.NRGBA{shieldUp1, shieldUp2}
	shieldSprites[down] = []*image.NRGBA{shieldDown1, shieldDown2}

	return shieldSprites
}

// NewCharacter returns an initialized Character struct
func NewCharacter(pos Position) *Character {
	walkSprites := getWalkingSprites()
	shieldSprites := getShieldSprites()
	walkSpritesEb := make(map[direction][]*ebiten.Image)
	shieldSpritesEb := make(map[direction][]*ebiten.Image)

	for _, dir := range []direction{left, right, up, down} {
		for _, sprite := range walkSprites[dir] {
			walkSpritesEb[dir] = append(walkSpritesEb[dir], ebiten.NewImageFromImage(sprite))
		}
		for _, sprite := range shieldSprites[dir] {
			shieldSpritesEb[dir] = append(shieldSpritesEb[dir], ebiten.NewImageFromImage(sprite))
		}
	}

	return &Character{
		position:      pos,
		dir:           up,
		walkSprites:   walkSpritesEb,
		shieldSprites: shieldSpritesEb,
		framesPerStep: 8,
	}
}
