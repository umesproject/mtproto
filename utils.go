// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of mtproto package.
// See https://github.com/umesproject/mtproto/blob/master/LICENSE for details

package mtproto

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/umesproject/mtproto/internal/encoding/tl"
	"github.com/umesproject/mtproto/internal/mtproto/objects"
	"github.com/xelaj/go-dry"
)

type any = interface{}
type null = struct{}

// это неофициальная информация, но есть подозрение, что список датацентров АБСОЛЮТНО идентичный для всех
// приложений. Несмотря на это, любой клиент ОБЯЗАН явно указывать список датацентров, ради надежности.
// данный список лишь эксперементальный и не является частью протокола.
func defaultDCList() map[int]string {
	return map[int]string{
		1: "149.154.175.58:443",
		2: "149.154.167.50:443",
		3: "149.154.175.100:443",
		4: "149.154.167.91:443",
		5: "91.108.56.151:443",
	}
}

func MessageRequireToAck(msg tl.Object) bool {
	switch msg.(type) {
	case /**objects.Ping,*/ *objects.MsgsAck:
		return false
	default:
		return true
	}
}

func CloseOnCancel(ctx context.Context, c io.Closer) {
	go func() {
		<-ctx.Done()
		c.Close()
	}()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func PrepareAppStorageForExamples() (appStoragePath string) {
	appStorage, err := GetAppStorage("mtproto-example", NamespaceUser)
	dry.PanicIfErr(err)

	if !dry.FileExists(appStorage) {
		if !dry.PathIsWritable(appStorage) {
			fmt.Printf("cant create app local storage at %v\n", appStorage)
			os.Exit(1)
		}
		err := os.MkdirAll(appStorage, 0755)
		dry.PanicIfErr(err)
	}

	return appStorage
}

type Namespace uint8

const (
	NamespaceUnknown Namespace = iota
	NamespaceGlobal
	NamespaceUser
	NamespaceDirectory
)

func GetAppStorage(appName string, namespace Namespace) (string, error) {
	switch namespace {
	case NamespaceGlobal:
		return filepath.Join("var", "lib", appName), nil
	case NamespaceUser:
		p, err := GetAppStorage(appName, NamespaceGlobal)
		if err != nil {
			return "", err
		}
		u, _ := user.Current()
		userPath, err := GetUserNamespaceDir(u.Username)
		if err != nil {
			return "", err
		}
		return filepath.Join(userPath, p), nil
	default:
		return "", errors.New("Incompatible feature for this namespace")
	}
}

func GetUserNamespaceDir(username string) (string, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return "", errors.Wrapf(err, "looking up '%v'", username)
	}

	return filepath.Join(u.HomeDir, ".local"), nil
}
