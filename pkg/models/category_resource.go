// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"time"
)

const CategoryResourceTableName = "category_resource"

type CategoryResource struct {
	CategoryId string
	ResourceId string
	Status     string
	CreateTime time.Time
	StatusTime time.Time
}

var CategoryResourceColumns = GetColumnsFromStruct(&CategoryResource{})

func NewCategoryResource(categoryId, resourceId, status string) *CategoryResource {
	return &CategoryResource{
		CategoryId: categoryId,
		ResourceId: resourceId,
		Status:     status,
		CreateTime: time.Now(),
		StatusTime: time.Now(),
	}
}
