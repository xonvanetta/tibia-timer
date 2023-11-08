package main

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

var (
	audioContext = audio.NewContext(32000)
)

type Alert struct {
	when time.Duration

	mute bool

	audioPlayer *audio.Player
}

func (a *Alert) SetVolume(vol float64) {
	a.audioPlayer.SetVolume(vol)
}

func (a *Alert) SetMute(m bool) {
	a.mute = m
}

func (a *Alert) Update(duration time.Duration) {
	if a.when == 0 {
		return
	}
	
	if a.mute {
		return
	}

	if duration == a.when && !a.audioPlayer.IsPlaying() {
		a.audioPlayer.Play()
		a.audioPlayer.Rewind()
	}
}

func NewAlert(when time.Duration, file string) (*Alert, error) {
	audioPlayer, err := loadAudioPlayer(file)
	if err != nil {
		return nil, err
	}

	return &Alert{
		when:        when,
		audioPlayer: audioPlayer,
	}, nil
}

func loadAudioPlayer(file string) (*audio.Player, error) {
	audioFile, err := fs.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	stream, err := mp3.DecodeWithoutResampling(audioFile)
	if err != nil {
		return nil, fmt.Errorf("failed to decode audio file: %w", err)
	}

	audioPlayer, err := audioContext.NewPlayer(stream)
	if err != nil {
		return nil, fmt.Errorf("failed to create audio player: %w", err)
	}
	return audioPlayer, nil
}
