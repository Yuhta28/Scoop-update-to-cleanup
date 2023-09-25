package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/schollz/progressbar/v3"
)

func main() {
	// コマンドを作成
	cmd := exec.Command("cmd.exe", "/C", "scoop update && scoop update *")

	// 標準出力をキャプチャ
	stdoutPipe, _ := cmd.StdoutPipe()

	// コマンドを実行
	err := cmd.Start()
	if err != nil {
		fmt.Println("コマンドの実行中にエラーが発生しました:", err)
		return
	}

	// プログレスバーの設定
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

	// 標準出力からデータを読み取り、プログレスバーに反映
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

	// コマンドの終了を待機
	err = cmd.Wait()
	if err != nil {
		fmt.Println("コマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("Scoopパッケージのアップデートが完了しました。")

	// クリーンアップとキャッシュ削除コマンドを作成
	cleanupCmd := exec.Command("cmd.exe", "/C", "scoop cleanup * && scoop cache rm *")
	
	// 標準出力をキャプチャ
	cleanupCmd.Stdout = os.Stdout
	cleanupCmd.Stderr = os.Stderr

	// クリーンアップとキャッシュ削除コマンドを実行
	err = cleanupCmd.Run()
	if err != nil {
		fmt.Println("クリーンアップコマンドの実行中にエラーが発生しました:", err)
		return
	}

	fmt.Println("古いパッケージのクリーンアップとキャッシュの削除が完了しました。")

	// プログラムが終了しないように待機
	fmt.Println("Enterキーを押してプログラムを終了してください...")
	var input string
	fmt.Scanln(&input)
}
