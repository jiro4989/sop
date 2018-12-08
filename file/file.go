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

func Copy(srcFile, dstFile string, ids ...int) error {
	return cp(srcFile, dstFile, -1, -1)
}

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

	// rootユーザは 0
	// UID/GIDにマイナス値は使わない
	if 0 <= uid && 0 <= gid {
		if err := dst.Chown(uid, gid); err != nil {
			log.Println(err)
			// chownはできなくてもしかたないのでreturnしない
		}
	}

	if 1 <= len(m) {
		if err := dst.Chmod(m[0]); err != nil {
			log.Println(err)
			// chmodはできなくてもしかたないのでreturnしない
		}
	}

	return nil
}

func CopyByName(srcFile, dstFile, owner, group, mode string) error {
	if owner == "" || group == "" {
		uid := -1
		gid := -1

		var m uint64
		var err error
		m, err = strconv.ParseUint(mode, 10, 32)
		if err != nil {
			log.Println(err)
			return err
		}

		return cp(srcFile, dstFile, uid, gid, os.FileMode(m))
	}

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

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		log.Println(err)
		return err
	}

	gid, err := strconv.Atoi(g.Gid)
	if err != nil {
		log.Println(err)
		return err
	}

	var m uint64
	m, err = strconv.ParseUint(mode, 10, 32)
	if err != nil {
		log.Println(err)
		return err
	}

	return cp(srcFile, dstFile, uid, gid, os.FileMode(m))
}

func Backup(srcFile string) error {
	// ファイルの有無判定。存在しなければ終了
	if _, err := os.Stat(srcFile); err != nil {
		return nil
	}
	// TODO
	now := time.Now().Format("2006-01-02_150405")
	dstFile := fmt.Sprintf("%s.%s", srcFile, now)
	return Copy(srcFile, dstFile)
}
