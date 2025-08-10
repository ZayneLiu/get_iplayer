package models

import "time"

// Programme represents a BBC iPlayer/Sounds programme
type Programme struct {
	Index       int       `json:"index"`
	PID         string    `json:"pid"`
	Name        string    `json:"name"`
	Episode     string    `json:"episode"`
	Channel     string    `json:"channel"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"` // in seconds
	Available   time.Time `json:"available"`
	Type        string    `json:"type"` // "tv" or "radio"
	Categories  []string  `json:"categories"`
	Thumbnail   string    `json:"thumbnail"`
	URL         string    `json:"url"`
}

// SearchResult represents search results from BBC iPlayer/Sounds
type SearchResult struct {
	Programmes []Programme `json:"programmes"`
	Total      int         `json:"total"`
	Query      string      `json:"query"`
	Type       string      `json:"type"`
	Channel    string      `json:"channel"`
}

// DownloadOptions represents options for downloading programmes
type DownloadOptions struct {
	TVQuality    []string `json:"tv_quality"`
	RadioQuality []string `json:"radio_quality"`
	Subtitles    bool     `json:"subtitles"`
	OutputDir    string   `json:"output_dir"`
	OutputFormat string   `json:"output_format"`
}

// Config represents application configuration
type Config struct {
	CacheDir    string          `json:"cache_dir"`
	OutputDir   string          `json:"output_dir"`
	DefaultType string          `json:"default_type"`
	Download    DownloadOptions `json:"download"`
}
