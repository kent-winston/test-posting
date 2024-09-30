package tools

import "myapp/model"

func UserToUserData(input model.User) *model.UserData {
	return &model.UserData{
		ID:        input.ID,
		Name:      input.Name,
		Email:     input.Email,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
		DeletedAt: input.DeletedAt,
	}
}
