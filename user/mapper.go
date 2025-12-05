package user

// Map User ↔ UserResponse
func ToUserResponse(u User) UserResponse {
	return UserResponse(u)
}

func ToUserResponseList(list []User) []UserResponse {
	res := make([]UserResponse, len(list))
	for i, u := range list {
		res[i] = UserResponse(u)
	}
	return res
}
