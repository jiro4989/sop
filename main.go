package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func CmdCp(srcFile, dstFile string) error {
	if err := BackupFile(dstFile); err != nil {
		return err
	}

	return Copy(srcFile, dstFile)
}

func Copy(srcFile, dstFile string) error {
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

	if err := os.Chown(dstFile, 0, 0); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func BackupFile(srcFile string) error {
	// ファイルの有無判定。存在しなければ終了
	if _, err := os.Stat(srcFile); err != nil {
		return nil
	}
	now := time.Now().Format("2006-01-02_150405")
	dstFile := fmt.Sprintf("%s.%s", srcFile, now)
	return Copy(srcFile, dstFile)
}

func main() {
	args := os.Args
	if err := CmdCp(args[1], args[2]); err != nil {
		panic(err)
	}
}
