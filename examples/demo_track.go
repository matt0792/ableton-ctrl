package examples

import (
	"fmt"
	"time"

	"github.com/matt0792/ableton-ctrl/als"
	"github.com/matt0792/ableton-ctrl/oscclient"
)

// DemoTrack creates a full demo track with drums, bass, and melody
func DemoTrack() {
	// Initialize client
	client := als.NewClient(oscclient.ClientOpts{
		SendAddr:   11000,
		ListenAddr: 11001,
	})

	client.Run()
	defer client.Close()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("🎵 Creating demo track...")

	// Set tempo to 120 BPM
	client.Song.SetTempo(120)
	fmt.Println("✓ Tempo set to 120 BPM")

	// Get current number of tracks
	numTracks := client.Song.GetNumTracks()
	fmt.Printf("✓ Current tracks: %d\n", numTracks)

	// Create 6 MIDI tracks for our demo
	trackNames := []string{"Kick", "Hi-Hat", "Bass", "Lead", "Chords", "Arp"}
	for i, name := range trackNames {
		client.Song.CreateMIDITrack(numTracks + int32(i))
		time.Sleep(50 * time.Millisecond)
		client.Track.SetName(numTracks+int32(i), name)
		fmt.Printf("✓ Created track: %s\n", name)
	}

	time.Sleep(200 * time.Millisecond)

	// Clip length: 4 bars = 16 beats at 4/4
	clipLength := float32(16.0)

	// Track 0: KICK DRUM - Four on the floor
	fmt.Println("\n🥁 Creating kick drum pattern...")
	kickTrack := numTracks
	client.ClipSlot.CreateClip(kickTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	kickNotes := []als.Note{
		{Pitch: 36, StartTime: 0.0, Duration: 0.25, Velocity: 100, Mute: false},  // Beat 1
		{Pitch: 36, StartTime: 4.0, Duration: 0.25, Velocity: 100, Mute: false},  // Beat 2
		{Pitch: 36, StartTime: 8.0, Duration: 0.25, Velocity: 100, Mute: false},  // Beat 3
		{Pitch: 36, StartTime: 12.0, Duration: 0.25, Velocity: 100, Mute: false}, // Beat 4
	}
	client.Clip.AddNotes(kickTrack, 0, kickNotes...)
	client.Clip.SetName(kickTrack, 0, "Four on Floor")
	fmt.Println("✓ Kick pattern created")

	// Track 1: HI-HAT - 16th note groove
	fmt.Println("\n🎩 Creating hi-hat pattern...")
	hihatTrack := numTracks + 1
	client.ClipSlot.CreateClip(hihatTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	hihatNotes := []als.Note{}
	// Create 16th note pattern with accents on off-beats
	for i := 0; i < 64; i++ { // 16 beats * 4 sixteenth notes per beat
		startTime := float32(i) * 0.25
		velocity := int32(70)
		if i%4 == 2 { // Accent every off-beat
			velocity = 90
		}
		if i%8 == 0 { // Stronger accent every 8th note
			velocity = 100
		}
		hihatNotes = append(hihatNotes, als.Note{
			Pitch:     42, // Closed hi-hat
			StartTime: startTime,
			Duration:  0.2,
			Velocity:  velocity,
			Mute:      false,
		})
	}
	client.Clip.AddNotes(hihatTrack, 0, hihatNotes...)
	client.Clip.SetName(hihatTrack, 0, "16th Groove")
	fmt.Println("✓ Hi-hat pattern created")

	// Track 2: BASSLINE - Funky syncopated bass
	fmt.Println("\n🎸 Creating bassline...")
	bassTrack := numTracks + 2
	client.ClipSlot.CreateClip(bassTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	bassNotes := []als.Note{
		// Bar 1 - C, with groove and ghost notes
		{Pitch: 36, StartTime: 0.0, Duration: 0.6, Velocity: 105, Mute: false},   // C (strong root)
		{Pitch: 36, StartTime: 0.75, Duration: 0.15, Velocity: 60, Mute: false},  // ghost note
		{Pitch: 36, StartTime: 1.5, Duration: 0.4, Velocity: 88, Mute: false},    // syncopation
		{Pitch: 38, StartTime: 2.25, Duration: 0.15, Velocity: 65, Mute: false},  // D ghost
		{Pitch: 38, StartTime: 2.5, Duration: 0.4, Velocity: 90, Mute: false},    // D
		{Pitch: 36, StartTime: 3.0, Duration: 0.25, Velocity: 70, Mute: false},   // back to C
		{Pitch: 40, StartTime: 3.5, Duration: 0.4, Velocity: 92, Mute: false},    // E (anticipation)
		// Bar 2 - F/A bass (matches Fmaj7/A chord)
		{Pitch: 33, StartTime: 4.0, Duration: 0.6, Velocity: 102, Mute: false},   // A (root of inversion)
		{Pitch: 33, StartTime: 4.75, Duration: 0.15, Velocity: 58, Mute: false},  // ghost
		{Pitch: 41, StartTime: 5.0, Duration: 0.5, Velocity: 95, Mute: false},    // F up
		{Pitch: 40, StartTime: 5.75, Duration: 0.25, Velocity: 72, Mute: false},  // E passing
		{Pitch: 38, StartTime: 6.25, Duration: 0.4, Velocity: 88, Mute: false},   // D
		{Pitch: 36, StartTime: 7.0, Duration: 0.15, Velocity: 62, Mute: false},   // C ghost
		{Pitch: 38, StartTime: 7.5, Duration: 0.5, Velocity: 85, Mute: false},    // D
		// Bar 3 - A (Am7)
		{Pitch: 33, StartTime: 8.0, Duration: 0.7, Velocity: 100, Mute: false},   // A root
		{Pitch: 33, StartTime: 9.0, Duration: 0.15, Velocity: 60, Mute: false},   // ghost
		{Pitch: 36, StartTime: 9.5, Duration: 0.4, Velocity: 87, Mute: false},    // C
		{Pitch: 40, StartTime: 10.25, Duration: 0.4, Velocity: 90, Mute: false},  // E
		{Pitch: 38, StartTime: 11.0, Duration: 0.25, Velocity: 70, Mute: false},  // D passing
		{Pitch: 36, StartTime: 11.5, Duration: 0.5, Velocity: 92, Mute: false},   // C (setup for G)
		// Bar 4 - G (G7) with walking feel
		{Pitch: 43, StartTime: 12.0, Duration: 0.8, Velocity: 105, Mute: false},  // G (strong)
		{Pitch: 43, StartTime: 13.0, Duration: 0.15, Velocity: 58, Mute: false},  // ghost
		{Pitch: 41, StartTime: 13.5, Duration: 0.6, Velocity: 95, Mute: false},   // F (7th color)
		{Pitch: 38, StartTime: 14.25, Duration: 0.4, Velocity: 88, Mute: false},  // D
		{Pitch: 40, StartTime: 15.0, Duration: 0.5, Velocity: 90, Mute: false},   // E (walk back to C)
		{Pitch: 38, StartTime: 15.5, Duration: 0.25, Velocity: 75, Mute: false},  // D passing
	}
	client.Clip.AddNotes(bassTrack, 0, bassNotes...)
	client.Clip.SetName(bassTrack, 0, "Funky Bass")
	fmt.Println("✓ Bassline created")

	// Track 3: LEAD MELODY - Catchy riff
	fmt.Println("\n🎹 Creating lead melody...")
	leadTrack := numTracks + 3
	client.ClipSlot.CreateClip(leadTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	leadNotes := []als.Note{
		// Bar 1 - Opening phrase with breathing and dynamics
		{Pitch: 60, StartTime: 0.5, Duration: 0.5, Velocity: 75, Mute: false},   // C (soft entry)
		{Pitch: 62, StartTime: 1.25, Duration: 0.5, Velocity: 80, Mute: false},  // D (building)
		{Pitch: 64, StartTime: 2.0, Duration: 0.4, Velocity: 88, Mute: false},   // E
		{Pitch: 65, StartTime: 2.75, Duration: 0.25, Velocity: 92, Mute: false}, // F (grace note)
		{Pitch: 67, StartTime: 3.25, Duration: 0.75, Velocity: 95, Mute: false}, // G (anticipation)
		// Bar 2 - Expressive peak with vibrato effect via varied velocities
		{Pitch: 69, StartTime: 4.0, Duration: 1.75, Velocity: 100, Mute: false}, // A (held peak)
		{Pitch: 67, StartTime: 6.0, Duration: 0.4, Velocity: 85, Mute: false},   // G (descent)
		{Pitch: 65, StartTime: 6.5, Duration: 0.3, Velocity: 78, Mute: false},   // F
		{Pitch: 64, StartTime: 7.0, Duration: 0.8, Velocity: 82, Mute: false},   // E (settling)
		// Bar 3 - Variation with more syncopation
		{Pitch: 62, StartTime: 8.25, Duration: 0.5, Velocity: 70, Mute: false},  // D (breath)
		{Pitch: 64, StartTime: 9.0, Duration: 0.4, Velocity: 80, Mute: false},   // E
		{Pitch: 65, StartTime: 9.5, Duration: 0.3, Velocity: 85, Mute: false},   // F (quick)
		{Pitch: 67, StartTime: 10.0, Duration: 1.2, Velocity: 93, Mute: false},  // G (held)
		{Pitch: 69, StartTime: 11.5, Duration: 0.5, Velocity: 88, Mute: false},  // A (pickup)
		// Bar 4 - Resolution with feeling
		{Pitch: 72, StartTime: 12.0, Duration: 1.5, Velocity: 100, Mute: false}, // C (octave, climax)
		{Pitch: 69, StartTime: 13.75, Duration: 0.5, Velocity: 90, Mute: false}, // A
		{Pitch: 67, StartTime: 14.5, Duration: 0.6, Velocity: 85, Mute: false},  // G
		{Pitch: 65, StartTime: 15.25, Duration: 0.75, Velocity: 75, Mute: false}, // F (fade out)
	}
	client.Clip.AddNotes(leadTrack, 0, leadNotes...)
	client.Clip.SetName(leadTrack, 0, "Lead Riff")
	fmt.Println("✓ Lead melody created")

	// Track 4: CHORDS - Atmospheric stabs
	fmt.Println("\n🎼 Creating chord stabs...")
	chordsTrack := numTracks + 4
	client.ClipSlot.CreateClip(chordsTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	chordNotes := []als.Note{
		// Bar 1 - Cmaj7 with smooth voice leading (staggered entry for warmth)
		{Pitch: 48, StartTime: 0.0, Duration: 3.8, Velocity: 68, Mute: false},  // C
		{Pitch: 52, StartTime: 0.05, Duration: 3.75, Velocity: 62, Mute: false}, // E
		{Pitch: 55, StartTime: 0.1, Duration: 3.7, Velocity: 60, Mute: false},  // G
		{Pitch: 59, StartTime: 0.15, Duration: 3.65, Velocity: 58, Mute: false}, // B (maj7)
		// Subtle chord stab at beat 3
		{Pitch: 52, StartTime: 2.0, Duration: 0.4, Velocity: 50, Mute: false},
		{Pitch: 55, StartTime: 2.0, Duration: 0.4, Velocity: 48, Mute: false},
		// Bar 2 - Fmaj7/A (first inversion for smoother bass movement)
		{Pitch: 53, StartTime: 4.0, Duration: 3.8, Velocity: 70, Mute: false},  // F
		{Pitch: 57, StartTime: 4.05, Duration: 3.75, Velocity: 64, Mute: false}, // A (bass)
		{Pitch: 60, StartTime: 4.1, Duration: 3.7, Velocity: 62, Mute: false},  // C
		{Pitch: 64, StartTime: 4.15, Duration: 3.65, Velocity: 60, Mute: false}, // E (maj7)
		// Add rhythm
		{Pitch: 57, StartTime: 6.5, Duration: 0.4, Velocity: 52, Mute: false},
		{Pitch: 60, StartTime: 6.5, Duration: 0.4, Velocity: 50, Mute: false},
		// Bar 3 - Am7 (adds emotional color, not just repeating C)
		{Pitch: 45, StartTime: 8.0, Duration: 3.8, Velocity: 66, Mute: false},  // A
		{Pitch: 48, StartTime: 8.05, Duration: 3.75, Velocity: 60, Mute: false}, // C
		{Pitch: 52, StartTime: 8.1, Duration: 3.7, Velocity: 58, Mute: false},  // E
		{Pitch: 55, StartTime: 8.15, Duration: 3.65, Velocity: 56, Mute: false}, // G (min7)
		// Anticipation stab
		{Pitch: 48, StartTime: 11.75, Duration: 0.3, Velocity: 54, Mute: false},
		{Pitch: 52, StartTime: 11.75, Duration: 0.3, Velocity: 52, Mute: false},
		// Bar 4 - G7 (dominant tension with extensions)
		{Pitch: 55, StartTime: 12.0, Duration: 4.0, Velocity: 72, Mute: false},  // G
		{Pitch: 59, StartTime: 12.05, Duration: 4.0, Velocity: 66, Mute: false}, // B
		{Pitch: 62, StartTime: 12.1, Duration: 4.0, Velocity: 64, Mute: false},  // D
		{Pitch: 65, StartTime: 12.15, Duration: 4.0, Velocity: 62, Mute: false}, // F (7th)
		// Rhythmic pulse
		{Pitch: 59, StartTime: 14.0, Duration: 0.4, Velocity: 55, Mute: false},
		{Pitch: 62, StartTime: 14.0, Duration: 0.4, Velocity: 53, Mute: false},
	}
	client.Clip.AddNotes(chordsTrack, 0, chordNotes...)
	client.Clip.SetName(chordsTrack, 0, "Chord Stabs")
	fmt.Println("✓ Chord stabs created")

	// Track 5: ARPEGGIO - Dancing texture
	fmt.Println("\n✨ Creating arpeggio pattern...")
	arpTrack := numTracks + 5
	client.ClipSlot.CreateClip(arpTrack, 0, clipLength)
	time.Sleep(50 * time.Millisecond)

	arpNotes := []als.Note{
		// Bar 1 - Cmaj7 arp (16th note triplet feel)
		{Pitch: 60, StartTime: 0.0, Duration: 0.3, Velocity: 65, Mute: false},   // C
		{Pitch: 64, StartTime: 0.33, Duration: 0.3, Velocity: 62, Mute: false},  // E
		{Pitch: 67, StartTime: 0.66, Duration: 0.3, Velocity: 60, Mute: false},  // G
		{Pitch: 71, StartTime: 1.0, Duration: 0.3, Velocity: 68, Mute: false},   // B
		{Pitch: 67, StartTime: 1.33, Duration: 0.3, Velocity: 58, Mute: false},  // G
		{Pitch: 64, StartTime: 1.66, Duration: 0.3, Velocity: 56, Mute: false},  // E
		{Pitch: 60, StartTime: 2.0, Duration: 0.3, Velocity: 63, Mute: false},   // C
		{Pitch: 64, StartTime: 2.33, Duration: 0.3, Velocity: 61, Mute: false},  // E
		{Pitch: 67, StartTime: 2.66, Duration: 0.3, Velocity: 59, Mute: false},  // G
		{Pitch: 71, StartTime: 3.0, Duration: 0.3, Velocity: 66, Mute: false},   // B
		{Pitch: 72, StartTime: 3.33, Duration: 0.3, Velocity: 64, Mute: false},  // C up
		{Pitch: 71, StartTime: 3.66, Duration: 0.3, Velocity: 60, Mute: false},  // B
		// Bar 2 - Fmaj7/A arp
		{Pitch: 57, StartTime: 4.0, Duration: 0.3, Velocity: 66, Mute: false},   // A
		{Pitch: 60, StartTime: 4.33, Duration: 0.3, Velocity: 63, Mute: false},  // C
		{Pitch: 64, StartTime: 4.66, Duration: 0.3, Velocity: 61, Mute: false},  // E
		{Pitch: 65, StartTime: 5.0, Duration: 0.3, Velocity: 69, Mute: false},   // F
		{Pitch: 64, StartTime: 5.33, Duration: 0.3, Velocity: 59, Mute: false},  // E
		{Pitch: 60, StartTime: 5.66, Duration: 0.3, Velocity: 57, Mute: false},  // C
		{Pitch: 57, StartTime: 6.0, Duration: 0.3, Velocity: 64, Mute: false},   // A
		{Pitch: 60, StartTime: 6.33, Duration: 0.3, Velocity: 62, Mute: false},  // C
		{Pitch: 64, StartTime: 6.66, Duration: 0.3, Velocity: 60, Mute: false},  // E
		{Pitch: 69, StartTime: 7.0, Duration: 0.3, Velocity: 67, Mute: false},   // A up
		{Pitch: 65, StartTime: 7.33, Duration: 0.3, Velocity: 63, Mute: false},  // F
		{Pitch: 64, StartTime: 7.66, Duration: 0.3, Velocity: 59, Mute: false},  // E
		// Bar 3 - Am7 arp (more melancholic)
		{Pitch: 57, StartTime: 8.0, Duration: 0.3, Velocity: 64, Mute: false},   // A
		{Pitch: 60, StartTime: 8.33, Duration: 0.3, Velocity: 61, Mute: false},  // C
		{Pitch: 64, StartTime: 8.66, Duration: 0.3, Velocity: 59, Mute: false},  // E
		{Pitch: 67, StartTime: 9.0, Duration: 0.3, Velocity: 67, Mute: false},   // G
		{Pitch: 64, StartTime: 9.33, Duration: 0.3, Velocity: 58, Mute: false},  // E
		{Pitch: 60, StartTime: 9.66, Duration: 0.3, Velocity: 56, Mute: false},  // C
		{Pitch: 57, StartTime: 10.0, Duration: 0.3, Velocity: 62, Mute: false},  // A
		{Pitch: 60, StartTime: 10.33, Duration: 0.3, Velocity: 60, Mute: false}, // C
		{Pitch: 64, StartTime: 10.66, Duration: 0.3, Velocity: 58, Mute: false}, // E
		{Pitch: 67, StartTime: 11.0, Duration: 0.3, Velocity: 65, Mute: false},  // G
		{Pitch: 69, StartTime: 11.33, Duration: 0.3, Velocity: 63, Mute: false}, // A up
		{Pitch: 67, StartTime: 11.66, Duration: 0.3, Velocity: 59, Mute: false}, // G
		// Bar 4 - G7 arp (building tension)
		{Pitch: 62, StartTime: 12.0, Duration: 0.3, Velocity: 67, Mute: false},  // D
		{Pitch: 65, StartTime: 12.33, Duration: 0.3, Velocity: 64, Mute: false}, // F
		{Pitch: 67, StartTime: 12.66, Duration: 0.3, Velocity: 62, Mute: false}, // G
		{Pitch: 71, StartTime: 13.0, Duration: 0.3, Velocity: 70, Mute: false},  // B
		{Pitch: 67, StartTime: 13.33, Duration: 0.3, Velocity: 60, Mute: false}, // G
		{Pitch: 65, StartTime: 13.66, Duration: 0.3, Velocity: 58, Mute: false}, // F
		{Pitch: 62, StartTime: 14.0, Duration: 0.3, Velocity: 65, Mute: false},  // D
		{Pitch: 67, StartTime: 14.33, Duration: 0.3, Velocity: 63, Mute: false}, // G
		{Pitch: 71, StartTime: 14.66, Duration: 0.3, Velocity: 61, Mute: false}, // B
		{Pitch: 74, StartTime: 15.0, Duration: 0.3, Velocity: 72, Mute: false},  // D up (climax)
		{Pitch: 71, StartTime: 15.33, Duration: 0.3, Velocity: 66, Mute: false}, // B
		{Pitch: 67, StartTime: 15.66, Duration: 0.3, Velocity: 62, Mute: false}, // G (resolves to loop)
	}
	client.Clip.AddNotes(arpTrack, 0, arpNotes...)
	client.Clip.SetName(arpTrack, 0, "Dancing Arp")
	fmt.Println("✓ Arpeggio pattern created")

	// Set loop points for all clips
	fmt.Println("\n🔄 Setting loop points...")
	for i := int32(0); i < 6; i++ {
		trackIdx := numTracks + i
		client.Clip.SetLoopStart(trackIdx, 0, 0.0)
		client.Clip.SetLoopEnd(trackIdx, 0, clipLength)
		time.Sleep(20 * time.Millisecond)
	}

	// Fire all clips to start playing
	fmt.Println("\n▶️  Launching all clips...")
	time.Sleep(200 * time.Millisecond)
	for i := int32(0); i < 6; i++ {
		trackIdx := numTracks + i
		client.ClipSlot.Fire(trackIdx, 0)
		time.Sleep(50 * time.Millisecond)
	}

	// Start playback
	client.Song.StartPlaying()
	defer client.Song.StopPlaying()

	fmt.Println("\n✨ Demo track is playing!")
	fmt.Println("Track layout:")
	fmt.Println("  Track 1: Kick - Four on the floor pattern")
	fmt.Println("  Track 2: Hi-Hat - 16th note groove with dynamics")
	fmt.Println("  Track 3: Bass - Groovy bassline with ghost notes and walking feel")
	fmt.Println("  Track 4: Lead - Expressive melody with natural phrasing")
	fmt.Println("  Track 5: Chords - Rich progression (Cmaj7 - Fmaj7 - Am7 - G7) with voice leading")
	fmt.Println("  Track 6: Arp - Dancing arpeggio texture")
	fmt.Println("\n🎉 Enjoy the music!")

	// Keep the program running for a bit so the clips can play
	time.Sleep(10 * time.Second)
}
