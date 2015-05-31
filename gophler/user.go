package main

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordLength = 8
	hashCost       = 10
	userIDLength   = 20
)

func NewUser(username, email, password string) (User, error) {
	user := User{
		Email:    email,
		Username: username,
	}
	if username == "" {
		return user, errNoUsername
	}

	if email == "" {
		return user, errNoEmail
	}

	if password == "" {
		return user, errNoPassword
	}

	if len(password) < passwordLength {
		return user, errPasswordTooShort
	}

	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errUsernameExists
	}

	existingUser, err = globalUserStore.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if existingUser != nil {
		return user, errEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), hashCost)

	user.HashedPassword = string(hashedPassword)
	user.ID = GenerateID("usr", userIDLength)

	return user, nil
}

func FindUser(username, password string) (*User, error) {
	out := &User{
		Username: username,
	}

	existingUser, err := globalUserStore.FindByUsername(username)
	if err != nil {
		return out, err
	}

	if existingUser == nil {
		return out, errCredintalsIncorrent
	}

	if bcrypt.CompareHashAndPassword([]byte(existingUser.HashedPassword), []byte(password)) != nil {
		return out, errCredintalsIncorrent
	}

	return existingUser, nil
}

func UpdateUser(user *User, email, currentPassword, newPassword string) (User, error) {
	out := *user
	out.Email = email

	existingUser, err := globalUserStore.FindByEmail(email)
	if err != nil {
		return out, err
	}
	if existingUser != nil && existingUser.ID != user.ID {
		return out, errEmailExists
	}

	user.Email = email
	if currentPassword == "" {
		return out, nil
	}

	if bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(currentPassword)) != nil {
		return out, errPasswordIncorrect
	}

	if newPassword == "" {
		return out, errNoPassword
	}

	if len(newPassword) < passwordLength {
		return out, errPasswordTooShort
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), hashCost)
	user.HashedPassword = string(hashedPassword)
	return out, err
}
