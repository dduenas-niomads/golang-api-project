package services

import (
	"go-app/data/requests"
	"go-app/data/responses"
)

type TagsService interface {
	Create(tags requests.CreateTagsRequest)
	Update(tags requests.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) responses.TagsResponse
	FindAll() []responses.TagsResponse
}
