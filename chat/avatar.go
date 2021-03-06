package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

//ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するエラー
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

//Avatarはユーザーのプロフィール画像を表す型です
type Avatar interface {
	//AvatarURLは指定されたクライアントのアバターのURLを返す。
	//問題が発生した場合エラーを返す。特にURLを取得できなかった場合にはErrNoAvataURLを返す
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar
func (a TryAvatars) GetAvatarURL(u ChatUser) (string,error){
	for _,avatar := range a{
		if url,err := avatar.GetAvatarURL(u); err == nil{
			return url,nil
		}
	}
	return "",ErrNoAvatarURL
}

type AuthAvatar struct {}
var UseAuthAvatar AuthAvatar
func (AuthAvatar) GetAvatarURL(u ChatUser)(string,error){
	url := u.AvatarURL()
		if url != ""{
			return url,nil
		}
	return "",ErrNoAvatarURL
}

type GravatarAvatar struct{}
var UseGravatar GravatarAvatar
func (GravatarAvatar) GetAvatarURL(u ChatUser) (string,error){
			return "//www.gravatar.com/avatar/"+ u.UniqueID(), nil
}

type FileSystemAvatar struct{}
var UseFileSystemAvatar FileSystemAvatar
func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string,error){
			if files, err := ioutil.ReadDir("avatars"); err == nil{
				for _,file :=range files{
					if file.IsDir() {
						continue
					}
					if match,_ := filepath.Match(u.UniqueID()+"*",file.Name()); match{
						return "/avatars/" + file.Name(),nil
					}
				}
			}
	return "",ErrNoAvatarURL
}
