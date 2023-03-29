// Code generated by mfd-generator v0.4.0; DO NOT EDIT.

//nolint:all
//lint:file-ignore U1000 ignore unused code, it's generated
package db

import (
	"time"
)

var Columns = struct {
	Category struct {
		ID, Title, OrderNumber, StatusID string
	}
	News struct {
		ID, Title, Preview, Content, CategoryID, TagIDs, CreatedAt, PublishedAt, StatusID string

		Category string
	}
	Tag struct {
		ID, Title, StatusID string
	}
}{
	Category: struct {
		ID, Title, OrderNumber, StatusID string
	}{
		ID:          "categoryId",
		Title:       "title",
		OrderNumber: "orderNumber",
		StatusID:    "statusId",
	},
	News: struct {
		ID, Title, Preview, Content, CategoryID, TagIDs, CreatedAt, PublishedAt, StatusID string

		Category string
	}{
		ID:          "newsId",
		Title:       "title",
		Preview:     "preview",
		Content:     "content",
		CategoryID:  "categoryId",
		TagIDs:      "tagIds",
		CreatedAt:   "createdAt",
		PublishedAt: "publishedAt",
		StatusID:    "statusId",

		Category: "Category",
	},
	Tag: struct {
		ID, Title, StatusID string
	}{
		ID:       "tagId",
		Title:    "title",
		StatusID: "statusId",
	},
}

var Tables = struct {
	Category struct {
		Name, Alias string
	}
	News struct {
		Name, Alias string
	}
	Tag struct {
		Name, Alias string
	}
}{
	Category: struct {
		Name, Alias string
	}{
		Name:  "categories",
		Alias: "t",
	},
	News: struct {
		Name, Alias string
	}{
		Name:  "news",
		Alias: "t",
	},
	Tag: struct {
		Name, Alias string
	}{
		Name:  "tags",
		Alias: "t",
	},
}

type Category struct {
	tableName struct{} `pg:"categories,alias:t,discard_unknown_columns"`

	ID          int    `pg:"categoryId,pk"`
	Title       string `pg:"title,use_zero"`
	OrderNumber int    `pg:"orderNumber,use_zero"`
	StatusID    int    `pg:"statusId,use_zero"`
}

type News struct {
	tableName struct{} `pg:"news,alias:t,discard_unknown_columns"`

	ID          int        `pg:"newsId,pk"`
	Title       string     `pg:"title,use_zero"`
	Preview     *string    `pg:"preview"`
	Content     *string    `pg:"content"`
	CategoryID  int        `pg:"categoryId,use_zero"`
	TagIDs      []int      `pg:"tagIds,array"`
	CreatedAt   time.Time  `pg:"createdAt,use_zero"`
	PublishedAt *time.Time `pg:"publishedAt"`
	StatusID    int        `pg:"statusId,use_zero"`

	Category *Category `pg:"fk:categoryId,rel:has-one"`
}

type Tag struct {
	tableName struct{} `pg:"tags,alias:t,discard_unknown_columns"`

	ID       int    `pg:"tagId,pk"`
	Title    string `pg:"title,use_zero"`
	StatusID int    `pg:"statusId,use_zero"`
}