package web

type CategoryUpdateRequest struct {
	Id   int
	Name string `validate:"required,min=1,max=255" json:"name"`
}
