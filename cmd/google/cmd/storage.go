/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	_ "github.com/TakeruTakeru/poc-go-micro-service/configs"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gstorage"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	"github.com/spf13/cobra"
)

var (
	Action  string
	Path    string
	Output  string
	Verbose bool
)

const (
	// READ_COMMAND = "read"
	UPLOAD_COMMAND   = "upload"
	DOWNLOAD_COMMAND = "download"
	// DELETE_COMMAND = "delete"
)

// storageCmd represents the storage command
var storageCmd = &cobra.Command{
	Use:   "storage",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			panic(fmt.Errorf("Needs at least one args."))
		}
		switch Action {
		case UPLOAD_COMMAND:
			upload(args)
		case DOWNLOAD_COMMAND:
			download(args)
		default:
			panic(fmt.Errorf("Invalid action type: %s", Action))
		}
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.Flags().StringVarP(&Action, "action", "a", UPLOAD_COMMAND, "Action type. Upload, upload, or etc.")
	storageCmd.Flags().StringVarP(&Path, "path", "p", "", "Target file path.")
	storageCmd.Flags().StringVarP(&Output, "output", "o", "", "Output file path.")
	storageCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Show detail.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func upload(args []string) {
	file, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Errorf("Failed to Read file. %v", err)
	}
	ctx := context.Background()
	conn := gstorage.NewGoogleStorageConnector(ctx, "GOOGLE_CLOUD_KEYFILE_JSON", "sodium-chalice-256606")
	client, err := conn.NewClient()
	if err != nil {
		fmt.Printf("Failed create google api client: %s\n", err)
		os.Exit(1)
	}
	fm, _ := models.NewFile(filepath.Base(Path), 0, file, filepath.Dir(Path), time.Now(), time.Now(), "", "")
	client.Upload(fm)

	if Verbose {
		pwd, _ := os.Getwd()
		fmt.Printf("pwd: %s\npath: %s\nData:\n %s\n", pwd, Path, string(file))
	}
}

func download(args []string) {
	fp, err := os.OpenFile(Output, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Failed open")
		os.Exit(1)
	}
	ctx := context.Background()
	conn := gstorage.NewGoogleStorageConnector(ctx, "GOOGLE_CLOUD_KEYFILE_JSON", "sodium-chalice-256606")
	client, err := conn.NewClient()
	if err != nil {
		fmt.Printf("Failed create google api client: %s\n", err)
		os.Exit(1)
	}
	fm, err := client.Download(args[0])
	if err != nil {
		fmt.Printf("Failed to download data. %v", err)
		os.Exit(1)
	}
	_, err = fmt.Fprintf(fp, string(fm.GetData()))
	if err != nil {
		fmt.Printf("Failed to read data. `%s` :%v", string(fm.GetData()), err)
		os.Exit(1)
	}
}
