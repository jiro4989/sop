// +build windows

package file

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"time"
)

// uid, gidはWindows環境では使用しない
func cp(srcFile, dstFile string, uid, gid int, m ...os.FileMode) error {
	b, err := ioutil.ReadFile(srcFile)
	if err != nil {
		log.Println(err)
		return err
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		log.Println(err)
		return err
	}
	defer dst.Close()

	if _, err := dst.Write(b); err != nil {
		log.Println(err)
		return err
	}

	if 1 <= len(m) {
		if err := dst.Chmod(m[0]); err != nil {
			log.Println(err)
			// chmodはできなくてもしかたないのでreturnしない
		}
	}

	return nil
}

func Copy(srcFile, dstFile, owner, group, mode string) error {
	var (
		uid int
		gid int
	)

	if owner == "" || group == "" {
		uid = -1
		gid = -1
	} else {
		u, err := user.Lookup(owner)
		if err != nil {
			log.Println(err)
			return err
		}

		g, err := user.LookupGroup(group)
		if err != nil {
			log.Println(err)
			return err
		}

		uid, err = strconv.Atoi(u.Uid)
		if err != nil {
			log.Println(err)
			return err
		}

		gid, err = strconv.Atoi(g.Gid)
		if err != nil {
			log.Println(err)
			return err
		}

	}

	if mode == "" {
		return cp(srcFile, dstFile, uid, gid)
	}

	m, err := strconv.ParseUint(mode, 8, 32)
	if err != nil {
		log.Println(err)
		return err
	}

	return cp(srcFile, dstFile, uid, gid, os.FileMode(m))
}

func Backup(srcFile string) error {
	var (
		fi  os.FileInfo
		err error
	)
	// ファイルの有無判定。存在しなければ終了
	fi, err = os.Stat(srcFile)
	if err != nil {
		return nil
	}

	var (
		uid = -1
		gid = -1
		m   = fi.Mode()
	)

	now := time.Now().Format("2006-01-02_150405")
	dstFile := fmt.Sprintf("%s.%s", srcFile, now)
	return cp(srcFile, dstFile, int(uid), int(gid), m)
}
