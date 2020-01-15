# Cloud Storage Service

```go
import "github.com/golang-common-packages/cloud-storage"
```

```go
cloudStorageConfig := &cloudStorage.Config {
		URL:          "",
		Username:     "",
		Password:     "",
		AccessToken:  "",
		RefreshToken: "",
}
```

```go
googleDrive := cloudStorage.NewFilestore(cloudStorage.DRIVE, cloudStorageConfig),
dropBox := cloudStorage.NewFilestore(cloudStorage.DROPBOX, cloudStorageConfig),
oneDrive := cloudStorage.NewFilestore(cloudStorage.ONEDRIVE, cloudStorageConfig),
sharePoint := cloudStorage.NewFilestore(cloudStorage.SHAREPOINT, cloudStorageConfig),
```

```go
fileInfo := &cloudStorage.FileModel {
			ParentID:      "",
			SourcesID:     "",
			DestinationID: "",
			Source:        "",
			Sources:       nil,
			Destination:   "",
			Destinations:  nil,
			Name:          "",
			MimeType:      "",
			Path:          "",
			Content:       nil,
			Query:         "",
}
```

```go
googleDrive.List(fileInfo)
googleDrive.Upload(fileInfo)
googleDrive.Download(fileInfo)
googleDrive.Delete(fileInfo)
...
```

# Note
[Check this template for more information and know how to use this service](https://github.com/golang-microservices/template)