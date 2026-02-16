package audio

// Capturer abstracts an audio input source.
// The default implementation uses PortAudio, but other backends
// (e.g., Android Oboe, iOS AVAudioEngine, WebRTC) can implement this.
type Capturer interface {
	// Start begins audio capture.
	Start() error
	// ReadFrame reads one frame of PCM audio. Blocks until a frame is available.
	ReadFrame() ([]int16, error)
	// Stop stops audio capture.
	Stop() error
	// Close releases all audio resources.
	Close() error
}

// Player abstracts an audio output sink.
type Player interface {
	// Start begins audio playback.
	Start() error
	// WriteFrame writes one frame of PCM audio to the output. Blocks until written.
	WriteFrame(frame []int16) error
	// Stop stops audio playback.
	Stop() error
}

// AudioEncoder abstracts audio encoding (e.g., PCM → Opus).
type AudioEncoder interface {
	// Encode encodes a PCM frame to compressed audio bytes.
	Encode(pcm []int16) ([]byte, error)
}

// AudioDecoder abstracts audio decoding (e.g., Opus → PCM).
type AudioDecoder interface {
	// Decode decodes compressed audio bytes to a PCM frame.
	Decode(data []byte) ([]int16, error)
	// DecodePLC performs Packet Loss Concealment.
	DecodePLC() ([]int16, error)
}

// DecoderFactory creates new AudioDecoder instances (one per remote speaker).
type DecoderFactory interface {
	// NewDecoder creates a new AudioDecoder.
	NewDecoder() (AudioDecoder, error)
}

// VoiceDetector abstracts Voice Activity Detection.
type VoiceDetector interface {
	// Process analyzes a PCM frame and returns true if voice is detected.
	Process(pcm []int16) bool
	// IsActive returns the current voice activity state without processing.
	IsActive() bool
	// PreBufferedFrames returns chronologically-ordered pre-buffered frames.
	PreBufferedFrames() [][]int16
	// SetThreshold updates the detection threshold.
	SetThreshold(threshold float64)
}

// DeviceLister abstracts the ability to enumerate audio devices.
type DeviceLister interface {
	// ListInputDevices returns available audio input devices.
	ListInputDevices() ([]DeviceEntry, error)
	// ListOutputDevices returns available audio output devices.
	ListOutputDevices() ([]DeviceEntry, error)
}

// Compile-time interface checks for the PortAudio/Opus implementations.
var (
	_ Capturer      = (*CaptureDevice)(nil)
	_ Player        = (*PlaybackDevice)(nil)
	_ AudioEncoder  = (*Encoder)(nil)
	_ AudioDecoder  = (*Decoder)(nil)
	_ VoiceDetector = (*VAD)(nil)
)
