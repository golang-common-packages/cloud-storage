package cloudStorage

import "context"

// Filestore interface for API storage
type Filestore interface {
	Search(fileModel *FileModel) (interface{}, error)
	Metadata(fileModel *FileModel) (interface{}, error)
	List(fileModel *FileModel) (interface{}, error)
	Upload(fileModel *FileModel) (interface{}, error)
	Download(fileModel *FileModel) (interface{}, error)
	Delete(fileModel *FileModel) error
	Move(fileModel *FileModel) (interface{}, []error)
	CreateFolder(fileModel *FileModel) (interface{}, error)
}

var (
	ctx = context.Background()
)

/*
	@DRIVE: Google Drive service
	@DROPBOX: Dropbox service
	@ONEDRIVE: OneDrive service
	@SHAREPOINT: Sharepoint service
*/
const (
	DRIVE = iota
	DROPBOX
	ONEDRIVE
	SHAREPOINT
)

// NewFilestore function for Factory Pattern
func NewFilestore(storageType int, config *Config) Filestore {
	switch storageType {
	case DRIVE:
		return NewDrive()
	case DROPBOX:
		return NewDropbox(config.AccessToken)
	case ONEDRIVE:
		return NewOneDrive(config.URL, config.AccessToken, config.RefreshToken)
	case SHAREPOINT:
		return NewSharepoint(config.URL, config.AccessToken, config.RefreshToken)
	}

	return nil
}
