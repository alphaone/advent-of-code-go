package day7

import (
	"fmt"
	"strings"
)

type folder struct {
	name     string
	isDir    bool
	size     int
	parent   *folder
	children []*folder
}

func parseFolder(input string) *folder {
	var currentFolder *folder
	var disk *folder

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		switch {
		case line == "$ cd /":
			// initial
			disk = &folder{name: "/", size: 0, children: []*folder{}, isDir: true}
			currentFolder = disk
		case line == "$ cd ..":
			// one level up
			currentFolder = currentFolder.parent
		case strings.HasPrefix(line, "$ cd "):
			var dst string
			fmt.Sscanf(line, "$ cd %s", &dst)

			var dstFolder *folder
			for _, child := range currentFolder.children {
				if child.isDir && child.name == dst {
					dstFolder = child
				}
			}
			currentFolder = dstFolder
		case line == "$ ls":
			// do nothing
		case strings.HasPrefix(line, "dir "):
			// new dir
			newFolder := &folder{size: 0, children: []*folder{}, isDir: true}
			fmt.Sscanf(line, "dir %s", &newFolder.name)

			newFolder.parent = currentFolder
			currentFolder.children = append(currentFolder.children, newFolder)
		default:
			newFile := &folder{children: []*folder{}, isDir: false}
			fmt.Sscanf(line, "%d %s", &newFile.size, &newFile.name)

			newFile.parent = currentFolder
			currentFolder.children = append(currentFolder.children, newFile)
		}
	}

	return disk
}

func (f *folder) totalSize() int {
	if len(f.children) == 0 {
		return f.size
	}

	total := f.size
	for _, child := range f.children {
		total += child.totalSize()
	}

	return total
}

func (f *folder) allDirs() []*folder {
	res := make([]*folder, 0)
	for _, child := range f.children {
		if child.isDir {
			res = append(res, child)
			res = append(res, child.allDirs()...)
		}
	}
	return res
}

func solveA(input string) int {
	disk := parseFolder(input)

	smallFolderSize := 0
	for _, child := range disk.allDirs() {
		if child.totalSize() < 100000 {
			smallFolderSize += child.totalSize()
		}
	}
	return smallFolderSize
}

func solveB(input string) int {
	disk := parseFolder(input)

	diskSpace := 40_000_000
	toFree := disk.totalSize() - diskSpace

	minFolderSizeToDelete := disk.totalSize()
	for _, child := range disk.allDirs() {
		s := child.totalSize()
		if s > toFree && s < minFolderSizeToDelete {
			minFolderSizeToDelete = s
		}
	}
	return minFolderSizeToDelete
}
