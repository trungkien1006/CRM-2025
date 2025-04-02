package helpers

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
)

func StoreFile(image *multipart.FileHeader) error {
	// Lưu file vào hệ thống
	file, err := image.Open()
	if err != nil {
		return errors.New("loi mo file")
	}
	defer file.Close()

	dst, err := os.Create("public/images/" + image.Filename)
	if err != nil {
		return errors.New("loi luu file")
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return errors.New("loi ghi file")
	}

	return nil
}