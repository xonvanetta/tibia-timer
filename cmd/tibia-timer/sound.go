package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sound interface {
	Component
	SetVolume(vol float64)
}

type sound struct {
	volumeChanged bool
	mute          bool
	volume        float64

	sounds []Sound
}

func NewSoundController(volume float64) *sound {
	return &sound{
		volume:        volume,
		volumeChanged: true,
	}
}

func (s *sound) Add(sound Sound) {
	s.sounds = append(s.sounds, sound)
}

func (s *sound) Update(duration time.Duration) {
	if s.volumeChanged {
		for _, sound := range s.sounds {
			sound.SetVolume(s.volume)
		}
		s.volumeChanged = false
	}

	if isPressed(ebiten.KeyUp) && s.volume <= 1 {
		s.volume += .01
		s.volumeChanged = true
	}
	if isPressed(ebiten.KeyDown) && s.volume > 0.01 {
		s.volume -= .01
		s.volumeChanged = true
	}

	if isPressed(ebiten.KeyM) {
		s.mute = !s.mute
	}

	if s.mute {
		return
	}

	for _, sound := range s.sounds {
		sound.Update(duration)
	}
}

func (s *sound) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Volume: %s", s.FormatVolume()), 5, 30)
}

func (s *sound) FormatVolume() string {
	if s.mute {
		return "mute"
	}
	return fmt.Sprintf("%.0f%%", s.volume*100)
}
