package main

import (
	"embed"
	"fmt"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Controller or System
type Controller interface {
	Update(duration time.Duration)
	Draw(screen *ebiten.Image)
}

type Component interface {
	Update(duration time.Duration)
}

type Screen struct {
	time      time.Time
	startTime time.Time

	controllers []Controller

	config *Config

	audioPlayer *audio.Player
	font        font.Face
}

var (
	//go:embed files
	fs embed.FS
)

func (s *Screen) Update() error {
	if time.Now().After(s.time) {
		s.time = s.time.Add(s.config.Interval)
	}

	duration := time.Now().Sub(s.time).Round(time.Second).Abs()
	for _, controller := range s.controllers {
		controller.Update(duration)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		s.time = time.Now()
		s.startTime = s.time
	}

	if isPressed(ebiten.KeyLeft) {
		s.time = s.time.Add(-time.Second)
	}
	if isPressed(ebiten.KeyRight) {
		s.time = s.time.Add(time.Second)
	}
	return nil
}

func isPressed(key ebiten.Key) bool {
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		return ebiten.IsKeyPressed(key)
	}
	return inpututil.IsKeyJustPressed(key)
}

func (s *Screen) Draw(screen *ebiten.Image) {
	text.Draw(screen, fmt.Sprintf("Time left: %s", FormatTime(s.time)), s.font, 5, 25, color.White)
	ebitenutil.DebugPrintAt(screen, "Backspace: Reset time", 5, 45)
	ebitenutil.DebugPrintAt(screen, "<-: Remove one second", 5, 60)
	ebitenutil.DebugPrintAt(screen, "->: Add one second", 5, 75)
	ebitenutil.DebugPrintAt(screen, "^: Increase volume", 5, 90)
	ebitenutil.DebugPrintAt(screen, "v: Lower volume", 5, 105)
	ebitenutil.DebugPrintAt(screen, "M: Mute", 5, 120)
	ebitenutil.DebugPrintAt(screen, "Shift: Increase action", 5, 135)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Total Time: %s", FormatTime(s.startTime)), 5, 150)

	for _, controller := range s.controllers {
		controller.Draw(screen)
	}

}

func FormatTime(t time.Time) string {
	seconds := time.Now().Sub(t).Round(time.Second).Abs()

	minute := seconds / time.Minute
	seconds -= minute * time.Minute
	seconds = seconds / time.Second

	return fmt.Sprintf("%02d:%02d", minute, seconds)
}

func (s *Screen) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	config := &Config{
		FirstAlarm:  MustParseDuration(FirstAlarm),
		SecondAlarm: MustParseDuration(SecondAlarm),
		Interval:    MustParseDuration(Interval),
		Volume:      0.1,
	}

	err := config.Validate()
	if err != nil {
		log.Fatalf("failed to validate config: %s", err)
	}

	font, err := loadTextFont()
	if err != nil {
		log.Fatalf("failed to load font: %s", err)
	}

	firstAlert, err := NewAlert(config.FirstAlarm, "files/quack.mp3")
	if err != nil {
		log.Fatalf("failed to load alert: %s", err)
	}

	secondAlert, err := NewAlert(config.SecondAlarm, "files/goat.mp3")
	if err != nil {
		log.Fatalf("failed to load alert: %s", err)
	}

	sounds := NewSoundController(config.Volume)
	sounds.Add(firstAlert)
	sounds.Add(secondAlert)

	s := &Screen{
		config:    config,
		time:      time.Now().Add(config.Interval),
		startTime: time.Now(),

		controllers: []Controller{sounds},
		font:        font,
	}

	ebiten.SetWindowSize(200, 180)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetWindowFloating(true)

	if err := ebiten.RunGame(s); err != nil {
		log.Fatal(err)
	}

}

func loadTextFont() (font.Face, error) {
	ttfFile, err := fs.ReadFile("files/tahoma.ttf")
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	tt, err := opentype.Parse(ttfFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	tahoma, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    26,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new font face: %w", err)
	}
	// Adjust the line height.
	//mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)

	return tahoma, nil
}
