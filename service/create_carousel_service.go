//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:11:04
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:52:26
 */
package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
)

// CreateCarouselService 轮播图创建的服务
type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

// Create 创建商品
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := model.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS

	err := model.DB.Create(&carousel).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousel(carousel),
	}
}
