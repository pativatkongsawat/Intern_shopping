package helper

type UserFilter struct {
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	PermissionId string `json:"permission_id"`
}
