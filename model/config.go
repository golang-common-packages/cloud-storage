package model

// Service struct provide all services
type Service struct {
	Database Database `json:"database,omitempty"`
}

// Database provide info for database connection.
type Database struct {
	OneDrive   OneDrive   `json:"onedrive,omitempty"`
	Dropbox    Dropbox    `json:"dropbox,omitempty"`
	Drive      Drive      `json:"drive,omitempty"`
	Sharepoint Sharepoint `json:"sharepoint,omitempty"`
}

// OneDrive provide a connection information for onedrive
type OneDrive struct {
	URL          string `json:"url,omitempty"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

// Dropbox provide a connection information for dropbox
type Dropbox struct {
	Token string `json:"token,omitempty"`
}

// Drive provide a connection information for drive
type Drive struct {
	APIKey string `json:"apiKey,omitempty"`
}

// Sharepoint provide a connection information for sharepoint
type Sharepoint struct {
	SiteURL  string `json:"siteURL,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
