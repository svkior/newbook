package main
import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"os"
)

const pageSize = 25

type ImageStore interface {
	Save(image *Image) error
	Find(id string) (*Image, error)
	FindAll(offset int) ([]Image, error)
	FindAllByUser(user *User, offset int) ([]Image, error)
}

type FileImageStore struct {
	filename string
	Images   map[string]Image
}

func (store FileImageStore) Save(image *Image) error {
	store.Images[image.ID] = *image

	contents, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(store.filename, contents, 0660)
	if err != nil {
		return err
	}
	return nil
}

func (store FileImageStore) Find(id string) (*Image, error) {
	image, ok := store.Images[id]
	if ok {
		return &image, nil
	}
	return nil, nil
}

func (store FileImageStore) FindAll(offset int) ([]Image, error) {
	images := []Image{}
	idx := 0

	for _, image := range store.Images {
		if idx >= (offset + pageSize) {
			return images, nil
		}
		idx++
		if idx < offset {
			continue
		}
		images = append(images, image)
	}

	return images, nil
}

func (store FileImageStore) FindAllByUser(user *User, offset int) ([]Image, error) {
	images := []Image{}
	idx := 0

	for _, image := range store.Images {
		if idx >= (offset + pageSize) {
			return images, nil
		}
		if image.UserID == user.ID {
			idx++
		} else {
			continue
		}
		if idx < offset {
			continue
		}
		images = append(images, image)
	}

	return images, nil
}

func NewFileImageStore(filename string) (*FileImageStore, error) {
	store := &FileImageStore{
		Images:   map[string]Image{},
		filename: filename,
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return store, nil
		}
		return nil, err
	}

	err = json.Unmarshal(contents, store)
	if err != nil {
		return nil, err
	}

	return store, nil
}

var globalImageStore ImageStore

func init() {
	store, err := NewFileImageStore("./data/images.json")
	if err != nil {
		panic(fmt.Errorf("Error creating images store: %s", err))
	}
	globalImageStore = store
}
