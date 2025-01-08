package src

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"
)

func ResumableCopy(src string, dest string, resumeAt int, chunkSize int, lag int) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to get statistics of destination file: %w", err)
	}
	if chunkSize > int(srcInfo.Size()) {
		return fmt.Errorf("chunk size is bigger than actual source file size")
	}

	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	defer destFile.Close()

	resumeAtInt64 := int64(resumeAt)
	_, err = srcFile.Seek(resumeAtInt64, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to set seek cursor on source file: %w", err)
	}
	_, err = destFile.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to set seek cursor on destination file: %w", err)
	}

	fmt.Printf("source: %v \n", src)
	fmt.Printf("destination: %v \n", dest)
	fmt.Printf("resume at: %v \n", resumeAt)
	fmt.Printf("chunk size: %v \n", chunkSize)
	fmt.Printf("lag: %v \n\n\n", lag)

	buffer := make([]byte, chunkSize)
	for {
		n, err := srcFile.Read(buffer)
		if n > 0 {
			nw, writeErr := destFile.Write(buffer)

			fmt.Printf("copied byte index %v to %v \n", resumeAtInt64, resumeAtInt64+int64(nw)-1)
			// cmd := exec.Command("cat", dest)
			// cmd.Stdout = os.Stdout
			// _ = cmd.Run()
			// fmt.Println()
			resumeAtInt64 += int64(nw)

			if writeErr != nil {
				return fmt.Errorf("failed to write destination file: %w", writeErr)
			}

			time.Sleep(time.Duration(lag) * time.Second)
		}

		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}

			return fmt.Errorf("failed to read source file: %w", err)
		}
	}
}
