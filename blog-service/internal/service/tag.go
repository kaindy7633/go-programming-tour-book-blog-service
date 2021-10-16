package service

import (
	"github.com/kaindy7633/go-programming-tour-book/blog-service/internal/model"
	"github.com/kaindy7633/go-programming-tour-book/blog-service/pkg/app"
)

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" json:"name" binding:"required,min=3,max=100"`
	CreatedBy string `form:"created_by" json:"created_by" binding:"required,min=3,max=100"`
	State     uint8  `form:"state,default=1" json:"state" binding:"oneof=0 1"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" json:"id" bingding:"required,gte=1"`
	Name       string `form:"name" json:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state" json:"state" binding:"required,oneof=0 1"`
	ModifiedBy string `form:"modified_by" json:"modified_by" binding:"required,min=3,max=100"`
}

type DelteteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountTag(param *CountTagRequest) (int, error) {
	return svc.dao.CountTag(param.Name, param.State)
}

func (svc *Service) GetTagList(param *TagListRequest, pager *app.Pager) ([]*model.Tag, error) {
	return svc.dao.GetTagList(param.Name, param.State, pager.Page, pager.PageSize)
}

func (svc *Service) CreateTag(param *CreateTagRequest) error {
	return svc.dao.CreateTag(param.Name, param.State, param.CreatedBy)
}

func (svc *Service) UpdateTag(param *UpdateTagRequest) error {
	return svc.dao.UpdateTag(param.ID, param.Name, param.State, param.ModifiedBy)
}

func (svc *Service) DeleteTag(param *DelteteTagRequest) error {
	return svc.dao.DeleteTag(param.ID)
}
