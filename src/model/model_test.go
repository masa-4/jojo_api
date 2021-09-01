package model_test

import (
	"testing"

	"github.com/masa-4/jojo_api/src/model"
	"github.com/stretchr/testify/assert"
)

func TestGetSubtitleByPart(t *testing.T) {
	subtitle := model.GetSubtitleByPart(1)

	assert.Equal(t, subtitle.SUBTITLE, "Fantom Blood")
}

func TestGetStandNameByCharacter(t *testing.T) {
	charactername := model.GetStandNameByCharacter("空条 承太郎")

	assert.Equal(t, charactername.Name, "スタープラチナ")
}
