package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/qor/media"
	"github.com/qor/media/media_library"
	"github.com/qor/sorting"
)

// Image is a model.
type Image struct {
	gorm.Model
	sorting.Sorting
	Name        string       `gorm:"size:255;"`
	Description string       `gorm:"type:text"`
	File        ImageStorage `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
	Tags        []Tag        `gorm:"many2many:image_tags;"`
}

// MarshalJSON returns the JSON representation.
func (t *Image) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		URL         string `json:"url"`
		Preview     string `json:"preview"`
		Tags        []Tag  `json:"tags"`
	}{
		Name:        t.Name,
		Description: t.Description,
		Tags:        t.Tags,
		URL:         t.File.URL("big"),
		Preview:     t.File.URL("small"),
	})
}

// ImageStorage is image storage.
type ImageStorage struct {
	media_library.MediaLibraryStorage
}

// GetSizes return the sizes.
func (ImageStorage) GetSizes() map[string]*media.Size {
	c := NewConfig()

	return map[string]*media.Size{
		"small": {
			Width:   c.ImageSmallWidth * 2,
			Height:  c.ImageSmallHeight * 2,
			Padding: true,
		},
		"small@ld": {
			Width:  c.ImageSmallWidth,
			Height: c.ImageSmallHeight,
		},
		"middle": {
			Width:  c.ImageMiddleWidth * 2,
			Height: c.ImageMiddleHeight * 2,
		},
		"middle@ld": {
			Width:  c.ImageMiddleWidth,
			Height: c.ImageMiddleHeight,
		},
		"big": {
			Width:  c.ImageBigWidth * 2,
			Height: c.ImageBigHeight * 2,
		},
		"big@ld": {
			Width:  c.ImageBigWidth,
			Height: c.ImageBigHeight,
		},
	}
}
