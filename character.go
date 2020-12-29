package main

import (
	"fmt"
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
	spriteImg     *ebiten.Image
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
func (character *Character) Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions) {
	dir := direction((character.counter / 200) % 4)
	subimg := character.walkSprites[dir][(character.counter/character.framesPerStep)%2]
	fmt.Printf("%v\n", subimg.Bounds())
	// minX, minY := subimg.Bounds().Min.X, subimg.Bounds().Min.Y
	// maxX, maxY := subimg.Bounds().Max.X, subimg.Bounds().Max.Y
	screen.DrawImage(subimg, options)
}

// Move bla
func (character *Character) Move(dx, dy int) {
	character.position.x += dx
	character.position.y += dy
}

//Update bla
func (character *Character) Update() {
	character.counter++
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

	walkLeft1 := spriteImage.SubImage(image.Rect(35, 11, 50, 26)).(*image.NRGBA)
	walkLeft2 := spriteImage.SubImage(image.Rect(52, 11, 67, 26)).(*image.NRGBA)

	walkRight1 := imaging.FlipH(walkLeft1)
	walkRight2 := imaging.FlipH(walkLeft2)

	walkUp1 := spriteImage.SubImage(image.Rect(18, 11, 33, 26)).(*image.NRGBA)
	walkUp2 := imaging.FlipH(walkUp1)

	walkDown1 := spriteImage.SubImage(image.Rect(1, 11, 16, 26)).(*image.NRGBA)
	walkDown2 := imaging.FlipH(walkDown1)

	walkSprites[left] = []*image.NRGBA{walkLeft1, walkLeft2}
	walkSprites[right] = []*image.NRGBA{walkRight1, walkRight2}
	walkSprites[up] = []*image.NRGBA{walkUp1, walkUp2}
	walkSprites[down] = []*image.NRGBA{walkDown1, walkDown2}

	return walkSprites
}

// NewCharacter bla
func NewCharacter() *Character {
	walkSprites := getWalkingSprites()
	walkSpritesEb := make(map[direction][]*ebiten.Image)

	for _, dir := range []direction{left, right, up, down} {
		for _, sprite := range walkSprites[dir] {
			walkSpritesEb[dir] = append(walkSpritesEb[dir], ebiten.NewImageFromImage(sprite))
		}
	}

	return &Character{
		position:      Position{0, 0},
		dir:           up,
		walkSprites:   walkSpritesEb,
		framesPerStep: 8,
	}
}
