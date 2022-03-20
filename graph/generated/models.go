// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"time"

	"github.com/99designs/gqlgen/graphql"

	"github.com/stegoer/server/ent"
	"github.com/stegoer/server/ent/image"
)

type Auth struct {
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type CreateUserPayload struct {
	User *ent.User `json:"user"`
	Auth *Auth     `json:"auth"`
}

type DecodeImageInput struct {
	LsbUsed int            `json:"lsbUsed"`
	Channel image.Channel  `json:"channel"`
	File    graphql.Upload `json:"file"`
}

type DecodeImagePayload struct {
	Message string `json:"message"`
}

type EncodeImageInput struct {
	Message string         `json:"message"`
	LsbUsed int            `json:"lsbUsed"`
	Channel image.Channel  `json:"channel"`
	File    graphql.Upload `json:"file"`
}

type EncodeImagePayload struct {
	Image *ent.Image `json:"image"`
	File  *FileType  `json:"file"`
}

type FileType struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ImagesConnection struct {
	TotalCount int              `json:"totalCount"`
	PageInfo   *ent.PageInfo    `json:"pageInfo"`
	Edges      []*ent.ImageEdge `json:"edges"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginPayload struct {
	User *ent.User `json:"user"`
	Auth *Auth     `json:"auth"`
}

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type OverviewPayload struct {
	User *ent.User `json:"user"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type RefreshTokenPayload struct {
	User *ent.User `json:"user"`
	Auth *Auth     `json:"auth"`
}

type UpdateUser struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type UpdateUserPayload struct {
	User *ent.User `json:"user"`
}
