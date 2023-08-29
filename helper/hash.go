package helper

import "golang.org/x/crypto/bcrypt"

func HasPassword(pass string) (string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(pass),14)
	return string(bytes),err
}

func CheckPassword(pass string, hash string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(pass))
	return err==nil
}