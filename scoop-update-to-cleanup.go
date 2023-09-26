package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {
	// make commands of update
	cmd := exec.Command("cmd.exe", "/C", "scoop update && scoop update *")

	// capture stdout
	stdoutPipe, _ := cmd.StdoutPipe()

	// catch error
	err := cmd.Start()
	if err != nil {
		fmt.Println("コマンドの実行中にエラーが発生しました:", err)
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

	// read data from stdout and update progress bar
	buffer := make([]byte, 1024)
	for {
		n, err := stdoutPipe.Read(buffer)
		if err != nil {
			break
		}
		output := string(buffer[:n])
		// Scoopの出力に進捗情報が含まれているか確認
		if strings.Contains(output, "Updating") {
			bar.Add(1)
		}
		fmt.Print(output)
	}

	// catch error
	err = cmd.Wait()
	if err != nil {
		fmt.Println("コマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("Scoopパッケージのアップデートが完了しました。")

	// make commands of cleanup and cache remove
	cleanupCmd := exec.Command("cmd.exe", "/C", "scoop cleanup * && scoop cache rm *")
	
	// capture stdout
	cleanupCmd.Stdout = os.Stdout
	cleanupCmd.Stderr = os.Stderr

	// run commands of cleanup and cache remove
	err = cleanupCmd.Run()
	if err != nil {
		fmt.Println("クリーンアップコマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("古いパッケージのクリーンアップとキャッシュの削除が完了しました。")

	// wait for enter key
	fmt.Println("Enterキーを押してプログラムを終了してください...")
	var input string
	fmt.Scanln(&input)
}
