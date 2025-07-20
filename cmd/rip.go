package cmd

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var batch string
var single string

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}

func downloadFromURL(fileURL string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Range", "bytes=0-")
	req.Header.Set("Referer", "https://cardiff.cloud.panopto.eu/")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	os.MkdirAll("lectures", os.ModePerm)
	filename := filepath.Join("lectures", fmt.Sprintf("lecture_%s.mp4", time.Now().Format("20060102_150405")))
	fmt.Println("Downloading to:", filename)

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("✅ Download complete: %s\n\n", filename)
	return nil
}

// ripCmd represents the rip command
var ripCmd = &cobra.Command{
	Use:   "rip",
	Short: "The rip command downloads full-length Panopto recordings based on a recording URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		if single != "" {
			if !isValidURL(single) {
				return fmt.Errorf("invalid URL: %s", single)
			}
			return downloadFromURL(single)
		}

		if batch != "" {
			file, err := os.Open(batch)
			if err != nil {
				return fmt.Errorf("cannot read file: %s", batch)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				url := strings.TrimSpace(scanner.Text())
				if url == "" || !isValidURL(url) {
					continue
				}
				if err := downloadFromURL(url); err != nil {
					fmt.Printf("❌ Error downloading %s: %v\n", url, err)
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
		}

		return nil
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if (single == "" && batch == "") || (single != "" && batch != "") {
			cmd.Help()
			return fmt.Errorf("you must specify exactly one of --single or --batch")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ripCmd)

	ripCmd.Flags().StringVarP(&batch, "batch", "b", "", "Path to a file containing URLs to download")
	ripCmd.Flags().StringVarP(&single, "single", "s", "", "A single Panopto video URL to download")
}
