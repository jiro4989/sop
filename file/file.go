package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func Copy(srcFile, dstFile string, ids ...int) error {
	b, err := ioutil.ReadFile(srcFile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	dst, err := os.Create(dstFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := dst.Write(b); err != nil {
		fmt.Println(err)
		return err
	}

	if 2 <= len(ids) {
		uid := ids[0]
		gid := ids[1]
		os.Chown(dstFile, uid, gid)
	}

	return nil
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
