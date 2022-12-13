package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"marshallformula.codes/utils"
)

var cdReg *regexp.Regexp
var lsReg *regexp.Regexp
var fileReg *regexp.Regexp
var dirReg *regexp.Regexp

var dirRefs = make(map[string]*directory, 0)

type file struct {
	name string
	size int
}

type directory struct {
	files     []*file
	dirs      []*directory
	totalSize int
}

func newDir() *directory {
	files := make([]*file, 0)
	dirs := make([]*directory, 0)
	return &directory{files, dirs, 0}
}

func (dir *directory) calculate() {
	sum := 0
	for _, f := range dir.files {
		sum += f.size
	}

	for _, d := range dir.dirs {
		d.calculate()
		sum += d.totalSize
	}

	dir.totalSize = sum
}

func main() {

	is, err := utils.InputScanner("input.txt")
	defer is.Close()

	if err != nil {
		log.Fatalln(err)
	}

	cdReg = regexp.MustCompile("\\$ cd (.+)")
	lsReg = regexp.MustCompile("\\$ ls")
	fileReg = regexp.MustCompile("(\\d+) (.+)")
	dirReg = regexp.MustCompile("dir (.+)")

	// root
	context := newDir()
	dirRefs["/"] = context
	pathStack := []string{"/"}

	cdRoot := true
	listing := false

	is.Scan(func(s string) {
		if cdRoot {
			cdRoot = false
			return
		}

		if lsReg.MatchString(s) {
			listing = true
			return
		}

		if fileReg.MatchString(s) {
			if !listing {
				log.Fatalln("somehow matched a file descriptor when we are not listing", s)
			}

			match := fileReg.FindStringSubmatch(s)
			size, err := strconv.Atoi(match[1])

			if nil != err {
				log.Fatalln("Failed to convert file size", match)
			}

			context.files = append(context.files, &file{match[2], size})
			return
		}

		if dirReg.MatchString(s) {
			if !listing {
				log.Fatalln("somehow matched a dir descriptor when we are not listing", s)
			}

			match := dirReg.FindStringSubmatch(s)

			dir := newDir()

			fullpath := strings.Join(pathStack, "/") + "/" + match[1]

			dirRefs[fullpath] = dir
			context.dirs = append(context.dirs, dir)
			return
		}

		if cdReg.MatchString(s) {
			listing = false

			match := cdReg.FindStringSubmatch(s)
			destDir := match[1]

			switch destDir {

			case "..":
				// pop last dir
				lastIdx := len(pathStack) - 1
				pathStack = pathStack[:lastIdx]

				fullPath := strings.Join(pathStack, "/")
				parentDir, ok := dirRefs[fullPath]

				if !ok {
					log.Fatalln("Couldn't access parent dir", pathStack, pathStack)
				}

				context = parentDir

			default:

				fullPath := strings.Join(pathStack, "/")
				knownDir, ok := dirRefs[fullPath+"/"+destDir]

				if !ok {
					log.Fatalln("Couldn't access known dir", pathStack, pathStack)
				}

				context = knownDir
				pathStack = append(pathStack, destDir)
			}

			return
		}

		log.Fatalln("Should not have made it here!")

	})

	for _, dir := range dirRefs {
		dir.calculate()
		// if dir.totalSize <= 100000 {
		// 	sum += dir.totalSize
		// }
	}

	root := dirRefs["/"]
	used := root.totalSize
	free := 70000000 - used
	needed := 30000000 - free

	candidate := root.totalSize

	for _, dir := range dirRefs {
		if dir.totalSize >= needed && dir.totalSize < candidate {
			candidate = dir.totalSize
		}
	}

	fmt.Println(candidate)

}
