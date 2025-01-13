package day7

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k
`

func TestParse(t *testing.T) {
	res := parseFolder(example)
	assert.Equal(t, 4, len(res.children))
}

func TestTotalSize(t *testing.T) {
	assert.Equal(t, 48381165, parseFolder(example).totalSize())
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 95437, solveA(example))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 1423358, solveA(utils.LoadString("input.txt")))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 24933642, solveB(example))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 545729, solveB(utils.LoadString("input.txt")))
}
