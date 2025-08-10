package api

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/ZayneLiu/get_iplayer/pkg/models"
)

// Client represents a BBC iPlayer/Sounds API client
type Client struct {
	baseURL   string
	userAgent string
}

// NewClient creates a new API client
func NewClient() *Client {
	return &Client{
		baseURL:   "https://www.bbc.co.uk",
		userAgent: "get_iplayer/3.36-go",
	}
}

// Search searches for programmes matching the given criteria
func (c *Client) Search(query, progType, channel string) (*models.SearchResult, error) {
	// For now, return mock data to demonstrate functionality
	// In a real implementation, this would make HTTP requests to BBC APIs

	result := &models.SearchResult{
		Query:   query,
		Type:    progType,
		Channel: channel,
		Total:   0,
	}

	// Add some mock programmes for demonstration
	if strings.Contains(strings.ToLower(query), "doctor") || query == ".*" {
		mockProgrammes := []models.Programme{
			{
				Index:       1,
				PID:         "b01rryzz",
				Name:        "Doctor Who: Series 7 Part 2",
				Episode:     "1. The Bells of Saint John",
				Channel:     "BBC One",
				Description: "Clara receives a call from the Doctor as he attempts to crack a sinister organisation called the Great Intelligence.",
				Duration:    2700, // 45 minutes
				Available:   time.Now().Add(-24 * time.Hour),
				Type:        "tv",
				Categories:  []string{"Drama", "Science Fiction"},
				Thumbnail:   "https://ichef.bbci.co.uk/images/ic/1920x1080/p01m1j05.jpg",
				URL:         "https://www.bbc.co.uk/iplayer/episode/b01rryzz",
			},
			{
				Index:       2,
				PID:         "b01rx0lj",
				Name:        "Doctor Who: Series 7 Part 2",
				Episode:     "2. The Rings Of Akhaten",
				Channel:     "BBC One",
				Description: "Clara's first adventure with the Doctor takes them to a festival on a distant planet.",
				Duration:    2700,
				Available:   time.Now().Add(-48 * time.Hour),
				Type:        "tv",
				Categories:  []string{"Drama", "Science Fiction"},
				Thumbnail:   "https://ichef.bbci.co.uk/images/ic/1920x1080/p01m1j06.jpg",
				URL:         "https://www.bbc.co.uk/iplayer/episode/b01rx0lj",
			},
		}

		// Filter by channel if specified
		if channel != "" {
			channelRegex, err := regexp.Compile("(?i)" + channel)
			if err == nil {
				filtered := []models.Programme{}
				for _, prog := range mockProgrammes {
					if channelRegex.MatchString(prog.Channel) {
						filtered = append(filtered, prog)
					}
				}
				mockProgrammes = filtered
			}
		}

		// Filter by type
		if progType != "" && progType != "tv,radio" {
			filtered := []models.Programme{}
			for _, prog := range mockProgrammes {
				if strings.Contains(progType, prog.Type) {
					filtered = append(filtered, prog)
				}
			}
			mockProgrammes = filtered
		}

		result.Programmes = mockProgrammes
		result.Total = len(mockProgrammes)
	}

	return result, nil
}

// GetProgrammeByPID retrieves a programme by its PID
func (c *Client) GetProgrammeByPID(pid string) (*models.Programme, error) {
	// Mock implementation - in reality this would make API calls
	programme := &models.Programme{
		PID:         pid,
		Name:        "Mock Programme",
		Episode:     "Episode from PID " + pid,
		Channel:     "BBC One",
		Description: "This is a mock programme retrieved by PID",
		Duration:    1800,
		Available:   time.Now(),
		Type:        "tv",
		Categories:  []string{"Mock"},
		URL:         fmt.Sprintf("https://www.bbc.co.uk/iplayer/episode/%s", pid),
	}

	return programme, nil
}
