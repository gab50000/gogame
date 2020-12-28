package game

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
)

type Position struct {
	x int
	y int
}

type mainCharacter struct {
	walkSprites   map[direction][]*ebiten.Image
	shieldSprites map[direction][]*ebiten.Image
	attackSprites map[direction][]*ebiten.Image
}

func (character *mainCharacter) Walk(dir direction) {

}

func (character *mainCharacter) Shield(dir direction) {

}

func (character *mainCharacter) Attack(dir direction) {

}

func (character *mainCharacter) Draw(screen *ebiten.Image) {

}

func imgFromFile(filename string) *ebiten.Image {
	file, err := ebitenutil.OpenFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}

func newMainCharacter() *mainCharacter {
	walkSprites := make(map[direction][]*ebiten.Image)
	spriteImage := imgFromFile("sprites.png")

	walkDown1 := spriteImage.SubImage(image.Rect(1, 11, 16, 26)).(*ebiten.Image)
	walkDown2 := ebiten.NewImageFromImage(imaging.FlipH(walkDown1))

	walkUp1 := spriteImage.SubImage(image.Rect(18, 11, 33, 26)).(*ebiten.Image)
	walkUp2 := ebiten.NewImageFromImage(imaging.FlipH(walkUp1))

	walkLeft1 := spriteImage.SubImage(image.Rect(35, 11, 50, 26)).(*ebiten.Image)
	walkLeft2 := spriteImage.SubImage(image.Rect(52, 11, 67, 26)).(*ebiten.Image)

	walkRight1 := ebiten.NewImageFromImage(imaging.FlipH(walkLeft1))
	walkRight2 := ebiten.NewImageFromImage(imaging.FlipH(walkLeft2))

	walkSprites[down] = []*ebiten.Image{walkDown1, walkDown2}
	walkSprites[up] = []*ebiten.Image{walkUp1, walkUp2}
	walkSprites[left] = []*ebiten.Image{walkLeft1, walkLeft2}
	walkSprites[right] = []*ebiten.Image{walkRight1, walkRight2}
}
