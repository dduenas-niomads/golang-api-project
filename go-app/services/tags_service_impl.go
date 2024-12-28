package services

import (
	"go-app/data/requests"
	"go-app/data/responses"
	"go-app/helpers"
	"go-app/models"
	"go-app/repositories"

	"github.com/go-playground/validator/v10"
)

type TagsServiceImpl struct {
	TagsRepository repositories.TagsRepository
	Validate       *validator.Validate
}

func NewTagsServiceImpl(tagRepository repositories.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}

// Create implements TagsService
func (t *TagsServiceImpl) Create(tags requests.CreateTagsRequest) {
	err := t.Validate.Struct(tags)
	helpers.ErrorPanic(err)
	tagModel := models.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

// Delete implements TagsService
func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}

// FindAll implements TagsService
func (t *TagsServiceImpl) FindAll() []responses.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []responses.TagsResponse
	for _, value := range result {
		tag := responses.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

// FindById implements TagsService
func (t *TagsServiceImpl) FindById(tagsId int) responses.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helpers.ErrorPanic(err)

	tagResponse := responses.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

// Update implements TagsService
func (t *TagsServiceImpl) Update(tags requests.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helpers.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
