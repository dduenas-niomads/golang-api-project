package repositories

import (
	"errors"
	"go-app/data/requests"
	"go-app/helpers"
	"go-app/models"

	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func NewTagsREpositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete implements TagsRepository
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.Db.Where("id = ?", tagsId).Delete(&tags)
	helpers.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository
func (t *TagsRepositoryImpl) FindAll() []models.Tags {
	var tags []models.Tags
	result := t.Db.Find(&tags)
	helpers.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository
func (t *TagsRepositoryImpl) FindById(tagsId int) (tags models.Tags, err error) {
	var tag models.Tags
	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository
func (t *TagsRepositoryImpl) Save(tags models.Tags) {
	result := t.Db.Create(&tags)
	helpers.ErrorPanic(result.Error)
}

// Update implements TagsRepository
func (t *TagsRepositoryImpl) Update(tags models.Tags) {
	var updateTag = requests.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helpers.ErrorPanic(result.Error)
}
