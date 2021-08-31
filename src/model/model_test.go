package model_test

import (
	"testing"

	"github.com/masa-4/jojo_api/src/model"
	"github.com/stretchr/testify/assert"
)

func TestGetSubtitleByPart(t *testing.T) {
	subtitle := model.GetSubtitleByPart(1)

	assert.Equal(t, subtitle, "Fantom Blood")
}

// func TestGetCharacterNameByStand(t *testing.T) {
// 	charactername := model.GetCharacterNameByStand("test character")

// 	assert.Equal(t, charactername)
// }
