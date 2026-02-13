package utils

import (
	"context"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
)

type FileSaver struct {
	envs configs.Environments
}

func NewFileSaver(envs configs.Environments) *FileSaver {
	if err := os.MkdirAll(envs.FILES_PATH.String(), os.ModePerm); err != nil {
		panic("failed to create files dest")
	}

	return &FileSaver{
		envs: envs,
	}
}

func (f *FileSaver) Save(ctx context.Context, destDirName, fileExtension string, file io.Reader) (string, error) {
	nowString := strconv.Itoa(int(time.Now().Unix()))
	randSuffix := strconv.Itoa(rand.Intn(10000))
	fileName := destDirName + "-" + nowString + "-" + randSuffix + "." + fileExtension

	fullDir := f.envs.FILES_PATH.String() + "/" + destDirName
	if err := os.MkdirAll(fullDir, os.ModePerm); err != nil {
		return "", apperrors.E(apperrors.ErrServerInternalError, err)
	}

	destPath := fullDir + "/" + fileName

	destFile, err := os.Create(destPath)
	if err != nil {
		return "", apperrors.E(apperrors.ErrServerInternalError, err)
	}
	defer destFile.Close()

	if _, err := io.Copy(destFile, file); err != nil {
		return "", apperrors.E(apperrors.ErrServerInternalError, err)
	}

	return destDirName + "/" + fileName, nil
}
