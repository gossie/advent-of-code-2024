package day9

import (
	"bufio"
	"math"
	"os"
)

type blockType int

const (
	file blockType = iota
	free
)

type block struct {
	bt blockType
	id int
}

type filesystem struct {
	blocks               []block
	fileSizes, freeSizes map[int]int
}

func readFile(filename string) filesystem {
	theFile, err := os.Open(filename)
	if err != nil {
		panic("failed opening file")
	}
	defer theFile.Close()

	fs := filesystem{
		blocks:    make([]block, 0),
		fileSizes: make(map[int]int),
		freeSizes: make(map[int]int),
	}
	fileIdCounter := 0
	freeIdCounter := 0

	scanner := bufio.NewScanner(theFile)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	for i, c := range scanner.Text() {
		size := int(c - '0')

		var bt blockType
		var id int
		switch {
		case i%2 == 0:
			bt = file
			id = fileIdCounter
			fs.fileSizes[id] = size
			fileIdCounter++
		default:
			bt = free
			id = freeIdCounter
			fs.freeSizes[id] = size
			freeIdCounter++
		}

		for range size {
			fs.blocks = append(fs.blocks, block{bt: bt, id: id})
		}
	}

	return fs
}

func rearrangeBlockwise(fs filesystem) filesystem {
	i, j := 0, len(fs.blocks)-1
	for i < j {
		if fs.blocks[i].bt == file {
			i++
			continue
		}

		if fs.blocks[j].bt == free {
			j--
			continue
		}

		fs.blocks[i], fs.blocks[j] = fs.blocks[j], fs.blocks[i]
	}
	return fs
}

func Part1(filename string) int {
	checksum := 0
	fs := rearrangeBlockwise(readFile(filename))
	for i := 0; i < len(fs.blocks) && fs.blocks[i].bt == file; i++ {
		checksum += i * fs.blocks[i].id
	}
	return checksum
}

func findLeftMostFittingFreeBlock(fs filesystem, neededSize int) int {
	i, found := 0, false

	for !found && i < len(fs.blocks) {
		if fs.blocks[i].bt == file {
			i += fs.fileSizes[fs.blocks[i].id]
			continue
		}

		if fs.freeSizes[fs.blocks[i].id] < neededSize {
			i += int(math.Max(float64(fs.freeSizes[fs.blocks[i].id]), 1))
			continue
		}

		return i
	}
	return -1
}

func rearrangeFilewise(fs filesystem) filesystem {
	visitedFiles := make(map[int]bool)
	for i := len(fs.blocks) - 1; i >= 0; i-- {
		blockId := fs.blocks[i].id
		if fs.blocks[i].bt == file && !visitedFiles[blockId] {
			fittingIndex := findLeftMostFittingFreeBlock(fs, fs.fileSizes[blockId])
			if fittingIndex != -1 && fittingIndex < i {
				fs.freeSizes[fs.blocks[fittingIndex].id] -= fs.fileSizes[blockId]
				for j := 0; j < fs.fileSizes[blockId]; j++ {
					fs.blocks[fittingIndex+j], fs.blocks[i-j] = fs.blocks[i-j], fs.blocks[fittingIndex+j]
				}
			}
			visitedFiles[blockId] = true
		}
	}
	return fs
}

func Part2(filename string) int {
	checksum := 0
	fs := rearrangeFilewise(readFile(filename))
	for i := 0; i < len(fs.blocks); i++ {
		if fs.blocks[i].bt == file {
			checksum += i * fs.blocks[i].id
		}
	}
	return checksum
}
