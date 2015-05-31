package main

import (
	"net/http"
	"time"
	"mime"
	"path/filepath"
	"io"
	"os"
	"mime/multipart"
)

const imageIDLength = 10

type Image struct {
	ID          string
	UserID      string
	Name        string
	Location    string
	Size        int64
	CreatedAt   time.Time
	Description string
}

func NewImage(user *User) *Image {
	return &Image{
		ID:        GenerateID("img", imageIDLength),
		UserID:    user.ID,
		CreatedAt: time.Now(),
	}

}

var mimeExtensions = map[string]string{
	"image/png":  ".png",
	"image/jpeg": ".jpg",
	"image/gif":  ".gif",
}

func (image *Image) CreateFromURL(imageURL string) error {
	// Get the response from the URL
	response, err := http.Get(imageURL)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errImageURLInvalid
	}

	defer response.Body.Close()

	mimeType, _, err := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if err != nil {
		return errInvalidImageType
	}

	// Get An Extension for the file
	ext, valid := mimeExtensions[mimeType]
	if !valid {
		return errInvalidImageType
	}

	image.Name = filepath.Base(imageURL)
	image.Location = image.ID + ext

	savedFile, err := os.Create("./data/images/" + image.Location)

	if err != nil {
		return err
	}
	defer savedFile.Close()

	size, err := io.Copy(savedFile, response.Body)
	if err != nil {
		return err
	}

	image.Size = size

	return globalImageStore.Save(image)
}

func (image *Image) CreateFromFile(file multipart.File, headers *multipart.FileHeader) error {
	image.Name = headers.Filename
	image.Location = image.ID + filepath.Ext(image.Name)

	savedFile, err := os.Create("./data/images/" + image.Location)
	if err != nil {
		return err
	}

	defer savedFile.Close()

	size, err := io.Copy(savedFile, file)
	if err != nil {
		return err
	}

	image.Size = size

	return globalImageStore.Save(image)
}
