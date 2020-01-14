package cloudStorage

import (
	"io"
	"log"
	"mime"
	"os"

	"google.golang.org/api/drive/v3"
)

// DriveServices manage all drive action
type DriveServices struct {
	driveService *drive.Service
}

// NewDrive function return a new mongo client based on singleton pattern
func NewDrive() Filestore {
	currentSession := &DriveServices{nil}

	driveService, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve GOOGLE_APPLICATION_CREDENTIALS %v", err)
	}

	currentSession.driveService = driveService
	log.Println("Connected to Google Drive")

	return currentSession
}

// Search ...
func (dr *DriveServices) Search(fileModel *FileModel) (interface{}, error) {
	return nil, nil
}

// Metadata ...
func (dr *DriveServices) Metadata(fileModel *FileModel) (interface{}, error) {
	return nil, nil
}

// List function return all files
func (dr *DriveServices) List(fileModel *FileModel) (interface{}, error) {
	files, err := dr.driveService.Files.List().Do()
	return files, err
}

// Upload function upload file to drive
func (dr *DriveServices) Upload(fileModel *FileModel) (interface{}, error) {
	f := &drive.File{
		MimeType: fileModel.MimeType,
		Name:     fileModel.Name,
		// Parents:  []string{parentID},
	}

	result, err := dr.driveService.Files.Create(f).Media(fileModel.Content).Do()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Download function will return a file base on fileID
func (dr *DriveServices) Download(fileModel *FileModel) (interface{}, error) {
	res, err := dr.driveService.Files.Get(fileModel.SourcesID).Download()
	if err != nil {
		return nil, err
	}

	// Get file extension
	fileExtension, err := mime.ExtensionsByType(res.Header.Get("Content-Type"))
	if err != nil {
		log.Println("Could not get file extension: " + err.Error())
	}

	// Create empty file with extension
	outFile, err := os.Create("uname" + fileExtension[0])
	if err != nil {
		return nil, err
	}
	defer outFile.Close()

	// Copy content to file that is created
	_, err = io.Copy(outFile, res.Body)
	if err != nil {
		log.Println("Could not copy content to file: " + err.Error())
	}

	return "uname" + fileExtension[0], nil
}

// Delete function will delete a file base on fileID
func (dr *DriveServices) Delete(fileModel *FileModel) error {
	err := dr.driveService.Files.Delete(fileModel.SourcesID).Do()
	return err
}

// Move function will move a file base on 'Sources' and 'Destination'
func (dr *DriveServices) Move(fileModel *FileModel) (interface{}, []error) {
	return nil, nil
}

// CreateFolder function will create a folder base on 'Destination'
func (dr *DriveServices) CreateFolder(fileModel *FileModel) (interface{}, error) {
	return nil, nil
}
