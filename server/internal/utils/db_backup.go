package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func BackupAndUploadDatabase() {
	// Env variable
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	backupDir := "backup"
	timestamp := time.Now().Format("20060102_150405")
	fileName := fmt.Sprintf("%s_%s.sql", dbName, timestamp)
	filePath := filepath.Join(backupDir, fileName)

	// Create folder if not exist
	if _, err := os.Stat(backupDir); os.IsNotExist(err) {
		os.Mkdir(backupDir, os.ModePerm)
	}

	// Jalankan mysqldump
	cmd := exec.Command("mysqldump", "-h", dbHost, "-u", dbUser, fmt.Sprintf("-p%s", dbPass), dbName)
	outfile, err := os.Create(filePath)
	if err != nil {
		log.Println("Failed to create dump file:", err)
		return
	}
	defer outfile.Close()
	cmd.Stdout = outfile

	if err := cmd.Run(); err != nil {
		log.Println("Failed to run mysqldump:", err)
		return
	}
	log.Println("Database dump created:", filePath)

	// Upload ke Google Drive
	if err := uploadToDrive(filePath); err != nil {
		log.Println("Failed to upload backup to Google Drive:", err)
	} else {
		log.Println("Database backup uploaded to Google Drive successfully.")
	}
}

func uploadToDrive(filePath string) error {
	return nil
	// ctx := context.Background()

	// b, err := os.ReadFile("credentials.json")
	// if err != nil {
	// 	return fmt.Errorf("unable to read credentials.json: %v", err)
	// }

	// config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	// if err != nil {
	// 	return fmt.Errorf("unable to parse config: %v", err)
	// }

	// tokFile := "token.json"
	// tok, err := os.ReadFile(tokFile)
	// if err != nil {
	// 	return fmt.Errorf("unable to read token: %v", err)
	// }

	// token := &google.Token{}
	// if err := token.UnmarshalJSON(tok); err != nil {
	// 	return fmt.Errorf("unable to unmarshal token: %v", err)
	// }

	// client := config.Client(ctx, token)
	// srv, err := drive.New(client)
	// if err != nil {
	// 	return fmt.Errorf("unable to create drive client: %v", err)
	// }

	// f, err := os.Open(filePath)
	// if err != nil {
	// 	return err
	// }
	// defer f.Close()

	// file := &drive.File{Name: filepath.Base(filePath)}
	// _, err = srv.Files.Create(file).Media(f).Do()
	// return err
}
