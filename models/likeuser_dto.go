package models

/*
LikeUserDTO key d t o

swagger:model LikeUserDTO
*/
type LikeUserDTO struct {
	Path string
	GlobalKey string
	Name string
	Avatar string
	Follow int32
	Followed int32
}