package main

import (
	"errors"
)

type ValidationError error

var (
	errNoUsername          = ValidationError(errors.New("Нужно указать имя файла"))
	errNoEmail             = ValidationError(errors.New("Нужно указать адрес электронной почты"))
	errNoPassword          = ValidationError(errors.New("Нужно указать пароль"))
	errPasswordTooShort    = ValidationError(errors.New("Пароль слишком короткий"))
	errUsernameExists      = ValidationError(errors.New("Такое имя пользователя уже есть"))
	errEmailExists         = ValidationError(errors.New("Такой адрес электронной почты уже зарегестрирован"))
	errCredintalsIncorrent = ValidationError(errors.New("Имя пользователя или пароль неправильный"))
	errPasswordIncorrect   = ValidationError(errors.New("Пароли не совпадают"))
	errInvalidIP		   = ValidationError(errors.New("Некорректный IP адрес"))
	errInvalidMask         = ValidationError(errors.New("Некорректная IP маска"))
	errInvalidGw           = ValidationError(errors.New("Некорректный Шлюз"))
	errInvalidMAC          = ValidationError(errors.New("Некорректный MAC адрес"))
	errInvalidArtnetInputs = ValidationError(errors.New("Некорректное значение числа входов ArtNet"))
)

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}
