package main

import (
	"runtime"
	"net/http"
	"time"
	"mime"
	"path/filepath"
	"io"
	"os"
	"mime/multipart"
	"github.com/disintegration/imaging"
	"image"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

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

	err = image.CreateResizedImages()
	if err != nil {
		return err
	}

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

	err = image.CreateResizedImages()
	if err != nil {
		return err
	}

	return globalImageStore.Save(image)
}

func (image *Image) StaticRoute() string {
	return "/im/" + image.Location
}

func (image *Image) ShowRoute() string {
	return "/image/" + image.ID
}

func (image *Image) UpdateResizedImages() error {
	ok := image.checkThumbnail()
	if !ok {
		image.CreateResizedImages()
	}
	return nil
}

func (image *Image) CreateResizedImages() error {
	srcImage, err := imaging.Open("./data/images/" + image.Location)
	if err != nil {
		return err
	}

	errorChan := make(chan error)

	go image.resizePreview(errorChan, srcImage)
	go image.resizeThumbnail(errorChan, srcImage)

	//var err error
	//err = nil
	for i := 0; i < 2; i++ {
		e := <-errorChan
		if e != nil {
			err = e
		}
	}
	return err
}

var widthThumbnail = 400

func (image *Image) checkThumbnail() bool {
	thumbFilename := "./data/images/thumbnail/" + image.Location
	if _, err := os.Stat(thumbFilename); os.IsNotExist(err){
		return false
	}
	return true
}

func (image *Image) resizeThumbnail(errorChan chan error, srcImage image.Image) {
	dstImage := imaging.Thumbnail(srcImage, widthThumbnail, widthThumbnail, imaging.Lanczos)
	destination := "./data/images/thumbnail/" + image.Location
	errorChan <- imaging.Save(dstImage, destination)
}

var widthPreview = 800

func (image *Image) resizePreview(errorChan chan error, srcImage image.Image){
	size := srcImage.Bounds().Size()
	ratio := float64(size.Y) / float64(size.X)
	targetHeight := int(float64(widthPreview) * ratio)
	dstImage := imaging.Resize(srcImage, widthPreview, targetHeight, imaging.Lanczos)

	destination := "./data/images/preview/" + image.Location

	errorChan <- imaging.Save(dstImage, destination)
}

func (image *Image) StaticThumbnailRoute() string {
	return "/im/thumbnail/"+ image.Location
}

func (image *Image) StaticPreviewRoute() string {
	return "/im/preview/"+ image.Location
}