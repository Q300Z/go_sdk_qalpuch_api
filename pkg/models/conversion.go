package models

// TaskConfig is the base config shared by all tasks.
type TaskConfig struct {
	Type string `json:"type" validate:"required,oneof=video image audio"`
}

// AudioConversionConfig holds the specific parameters for audio conversion.
type AudioConversionConfig struct {
	TaskConfig `validate:"required"`
	Codec      string `json:"codec,omitempty"   validate:"omitempty,oneof=mp3 aac opus"`
	Bitrate    int    `json:"bitrate,omitempty" validate:"omitempty,gt=0"` // in kbps
}

// NewAudioConfig creates a new AudioConversionConfig with the type pre-filled.

func (c *AudioConversionConfig) WithCodec(codec string) *AudioConversionConfig {
	c.Codec = codec
	return c
}

func (c *AudioConversionConfig) WithBitrate(kbps int) *AudioConversionConfig {
	c.Bitrate = kbps
	return c
}

// --- Video Config ---

// VideoConversionConfig holds the specific parameters for video conversion.
type VideoConversionConfig struct {
	TaskConfig `validate:"required"`
	Codec      string `json:"codec,omitempty"      validate:"omitempty,oneof=h264 vp9 av1"`
	Bitrate    int    `json:"bitrate,omitempty"    validate:"omitempty,gt=0"` // in kbps
	Resolution string `json:"resolution,omitempty" validate:"omitempty,regexp=^[1-9][0-9]*x[1-9][0-9]*$"`
}

// NewVideoConfig creates a new VideoConversionConfig with the type pre-filled.
func NewVideoConfig() *VideoConversionConfig {
	return &VideoConversionConfig{TaskConfig: TaskConfig{Type: "video"}}
}

func (c *VideoConversionConfig) WithCodec(codec string) *VideoConversionConfig {
	c.Codec = codec
	return c
}

func (c *VideoConversionConfig) WithBitrate(kbps int) *VideoConversionConfig {
	c.Bitrate = kbps
	return c
}

func (c *VideoConversionConfig) WithResolution(resolution string) *VideoConversionConfig {
	c.Resolution = resolution
	return c
}

// --- Image Config ---

// ImageConversionConfig holds the specific parameters for image conversion.
type ImageConversionConfig struct {
	TaskConfig `validate:"required"`
	Format     string `json:"format,omitempty"  validate:"omitempty,oneof=jpeg png webp"`
	Quality    int    `json:"quality,omitempty" validate:"omitempty,gte=1,lte=100"` // 1-100
	Width      int    `json:"width,omitempty"   validate:"omitempty,gt=0"`
	Height     int    `json:"height,omitempty"  validate:"omitempty,gt=0"`
}

// NewImageConfig creates a new ImageConversionConfig with the type pre-filled.

func (c *ImageConversionConfig) WithFormat(format string) *ImageConversionConfig {
	c.Format = format
	return c
}

func (c *ImageConversionConfig) WithQuality(quality int) *ImageConversionConfig {
	c.Quality = quality
	return c
}

func (c *ImageConversionConfig) WithWidth(width int) *ImageConversionConfig {
	c.Width = width
	return c
}

func (c *ImageConversionConfig) WithHeight(height int) *ImageConversionConfig {
	c.Height = height
	return c
}

// --- Music Config ---

// MusicConversionConfig holds the specific parameters for music conversion.
type MusicConversionConfig struct {
	TaskConfig `validate:"required"`
	Codec      string `json:"codec,omitempty"   validate:"omitempty,oneof=mp3 aac opus"`
	Bitrate    int    `json:"bitrate,omitempty" validate:"omitempty,gt=0"` // in kbps
}

// NewMusicConfig creates a new MusicConversionConfig with the type pre-filled.

func (c *MusicConversionConfig) WithCodec(codec string) *MusicConversionConfig {
	c.Codec = codec
	return c
}

func (c *MusicConversionConfig) WithBitrate(kbps int) *MusicConversionConfig {
	c.Bitrate = kbps
	return c
}
