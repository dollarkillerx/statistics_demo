package test

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

func TestRp(t *testing.T) {
	r := math.Round(0.012345*100) / 100
	fmt.Println(r)
}

func TestLogs(t *testing.T) {
	create, err := os.Create("test.log")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		time.Sleep(time.Second)
		TailFile("test.log")
	}()

	for i := 0; i < 100; i++ {
		create.WriteString(fmt.Sprintf("line %d\n", i))
		time.Sleep(time.Second)
	}
}

// TailFile 持续读取日志文件中的新内容
func TailFile(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	// 移动到文件末尾
	file.Seek(0, io.SeekEnd)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// 如果到达文件末尾，则等待新内容
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			return err
		}
		fmt.Println(strings.TrimSpace(line))
	}
}
