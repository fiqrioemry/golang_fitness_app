package utils

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	drive "google.golang.org/api/drive/v3"
)

func getDriveService() (*drive.Service, error) {
	ctx := context.Background()

	config := &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{drive.DriveFileScope},
	}

	token := &oauth2.Token{
		RefreshToken: os.Getenv("GOOGLE_REFRESH_TOKEN"),
		TokenType:    "Bearer",
	}

	client := config.Client(ctx, token)
	return drive.New(client)
}

func UploadFileToDrive(filePath, fileName string) error {
	service, err := getDriveService()
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	fileMeta := &drive.File{
		Name:     fileName,
		MimeType: "application/sql",
	}
	if folderID := os.Getenv("GOOGLE_FOLDER_ID"); folderID != "" {
		fileMeta.Parents = []string{folderID}
	}

	_, err = service.Files.Create(fileMeta).Media(f).Do()
	if err != nil {
		return fmt.Errorf("upload failed: %w", err)
	}

	fmt.Println("✅ Uploaded", fileName, "to Google Drive")
	return nil
}
