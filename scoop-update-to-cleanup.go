//go:generate go-winres make --product-version=git-tag
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {
	// make commands of update
	cmd := exec.Command("cmd.exe", "/C", "scoop update && scoop update *")

	// capture stdout and stderr
	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	// start command
	err := cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, "コマンドの実行中にエラーが発生しました:", err)
		return
	}

	// configure progress bar
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetDescription("Scoopアップデート中"),
		progressbar.OptionSpinnerType(14),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	// read data from stdout and stderr, and update progress bar
	go readPipe(stdoutPipe, bar)
	go readPipe(stderrPipe, bar)

	// wait for command to finish
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, "コマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("Scoopパッケージのアップデートが完了しました。")

	// make commands of cleanup and cache remove
	cleanupCmd := exec.Command("cmd.exe", "/C", "scoop cleanup * && scoop cache rm *")

	// capture stdout and stderr
	cleanupCmd.Stdout = os.Stdout
	cleanupCmd.Stderr = os.Stderr

	// run commands of cleanup and cache remove
	err = cleanupCmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "クリーンアップコマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("古いパッケージのクリーンアップとキャッシュの削除が完了しました。")

	// wait for enter key
	fmt.Println("Enterキーを押してプログラムを終了してください...")
	var input string
	fmt.Scanln(&input)
}

func readPipe(pipe io.Reader, bar *progressbar.ProgressBar) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		output := scanner.Text()
		if strings.Contains(output, "Updating") {
			bar.Add(1)
		}
		fmt.Println(output)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "パイプの読み取り中にエラーが発生しました:", err)
	}
}
