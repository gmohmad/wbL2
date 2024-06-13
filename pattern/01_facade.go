package pattern

import "fmt"

/*
Facade pattern is a structural pattern that provides a simplified interface to a complex subsystem of classes, making
it easier to use.

Pros:
    - Simplifies Usage and Reduces Complexity
	- Decouples Subsystem from Clients
Cons:
    - Introduces an Additional Layer
*/

// Example

// Define 'AudioPlayer' struct 
type AudioPlayer struct{}

// Define 'PlayAudio' method on 'AudioPlayer'
func (a *AudioPlayer) PlayAudio() {
	fmt.Println("playing audio...")
}

// Define 'VideoPlayer' struct 
type VideoPlayer struct{}

// Define 'PlayVideo' method on 'VideoPlayer'
func (v *VideoPlayer) PlayVideo() {
	fmt.Println("playing video...")
}

// Define 'ScreeManager' struct 
type ScreeManager struct{}

// Define 'ShowScreen' method on 'ScreeManager'
func (s *ScreeManager) ShowScreen() {
	fmt.Println("showing screen...")
}

// Define 'MultimediaFacade' struct that encapsulates 'AudioPlayer', 'videoPlayer' and 'ScreenManager'
type MultimediaFacade struct {
	audioPlayer   *AudioPlayer
	videoPlayer   *VideoPlayer
	screenManager *ScreeManager
}

// 'NewMultimediaFacade' returns a new MultimediaFacade
func NewMultimediaFacade() *MultimediaFacade {
	return &MultimediaFacade{
		audioPlayer:   &AudioPlayer{},
		videoPlayer:   &VideoPlayer{},
		screenManager: &ScreeManager{},
	}
}

// 'PlayMovie' method on 'MultimediaFacade' internally calls 'PlayAudio', 'PlayVideo' and 'ShowScreen' methods 
// providing a simplified interface ot 'AudioPlayer', 'videoPlayer' and 'ShowScreen' structs
func (m *MultimediaFacade) PlayMovie() {
	m.audioPlayer.PlayAudio()
	m.videoPlayer.PlayVideo()
	m.screenManager.ShowScreen()
}

// Example usage
func main() {

	// Create a MultimediaFacade
	multimediaSystem := NewMultimediaFacade()

	fmt.Println("Playing a movie...")

	// Using the Facade to play a movie
	multimediaSystem.PlayMovie()
}
