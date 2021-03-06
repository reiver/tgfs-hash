package main

import (
	"github.com/reiver/tgfs-hash/cl"
	"github.com/reiver/tgfs-hash/hash"
	"github.com/reiver/tgfs-hash/src"

	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Printf("\x1b[1mThe Great File System\x1b[0m (\x1b[1mTGFS\x1b[0m) — \x1b[1m%s\x1b[0m\n\n", cl.CommandName)

	filename, err := src.FileName()
	if nil != err {
		switch err {
		case src.ErrFileNameNotFound:
			cl.Usage()
			os.Exit(1)
		default:
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
	}
	fmt.Printf("File Name: %q\n\n", filename)

	digest := hashFunc(filename)

	var hexadecimal string
	{
		hexadecimal = fmt.Sprintf("%x", digest)
	}
	fmt.Printf("Digest (hexadecimal): %s\n\n", hexadecimal)

	var base64url string
	{
		base64url = base64.StdEncoding.EncodeToString(digest)
		base64url = strings.Replace(base64url, "+", "-", -1) // 62nd character
		base64url = strings.Replace(base64url, "/", "_", -1) // 63rd character
		base64url = strings.Replace(base64url, "=", "", -1) // remove padding
	}
	fmt.Printf("Digest (base64url): %s\n\n", base64url)

	rootPath := fmt.Sprintf(".tgfs/digest/sha-3-512/%s/%s/%s", base64url[0:1], base64url[1:2], base64url[2:])
	contentPath := fmt.Sprintf("%s/content", rootPath)
	headPath := fmt.Sprintf("%s/head", rootPath)
	fmt.Printf("Root Path: %s\n\n", rootPath)
	fmt.Printf("Content Path: %s\n\n", contentPath)
	fmt.Printf("Head Path: %s\n\n", headPath)

	if cl.Put {
		fmt.Printf("Will try to put file %q into content-addressable storage.\n\n", filename)

		if err := os.MkdirAll(rootPath, 0755); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Could not created directory %q: %s\n", rootPath, err)
			os.Exit(1)
		}

		if err := os.Rename(filename, contentPath); nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Problem moving file %q to content-addressable storage, at %q: %s\n", filename, contentPath, err)
			os.Exit(1)
		}
		fmt.Printf("File %q moved to %q\n\n", filename, contentPath)
	}
}


func hashFunc(filename string) []byte {
	var file *os.File
	{
		var err error

		file, err = os.Open(filename)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: Could not open file %q: %s\n", filename, err)
			os.Exit(1)
		}
		if nil == file {
			fmt.Fprintf(os.Stderr, "ERROR: Could not open file %q: return file is nil\n", filename)
			os.Exit(1)
		}
	}
	defer file.Close()

	func(){
		fileinfo, err := file.Stat()
		if nil != err {
			fmt.Fprintf(os.Stderr, "WARNING: Could not get file info about file %q: %s\n\n", filename, err)
			return
		}
		if nil == fileinfo {
			fmt.Fprintf(os.Stderr, "WARNING: Could not get file info about file %q: returned file info is nil\n\n", filename)
			return
		}

		fmt.Printf("File %q is %d bytes long.\n\n", filename, fileinfo.Size())
	}()

	digest, n, err := hash.Func(file)
	if nil != err {
		fmt.Fprintf(os.Stderr, "ERROR: Could not generate hash function digest for file %s: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%d bytes hashed using %s.\n\n", n, hash.Name)

	return digest
}
