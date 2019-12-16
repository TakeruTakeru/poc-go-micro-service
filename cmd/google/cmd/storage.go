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
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/gstorage"
	"github.com/TakeruTakeru/poc-go-micro-service/internal/app/fileservice/models"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var (
	Path    string
	Verbose bool
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
		f, err := os.Stat(Path)
		if err != nil {
			fmt.Println("Invalid Path")
			os.Exit(1)
		}
		fp, err := os.Open(Path)
		if err != nil {
			fmt.Println("Failed open")
			os.Exit(1)
		}
		defer fp.Close()
		buf := make([]byte, 1024)
		for {
			n, err := fp.Read(buf)
			if n == 0 {
				break
			}
			if err != nil {
				fmt.Printf("Failed read: %s\n", err)
				os.Exit(1)
			}
		}
		ctx := context.Background()
		conn := gstorage.NewGoogleStorageConnector(ctx, "GOOGLE_CLOUD_KEYFILE_JSON", "sodium-chalice-256606")
		client, err := conn.NewClient()
		if err != nil {
			fmt.Printf("Failed create google api client: %s\n", err)
		}
		fm, _ := models.NewFile(f.Name(), 0, "test-dir021900/"+f.Name(), time.Now(), time.Now(), "", "")
		fm.Data = buf
		client.Upload(fm)

		if Verbose {
			pwd, _ := os.Getwd()
			fmt.Printf("pwd: %s\npath: %s\nData:\n %s\n", pwd, Path, string(buf))
		}
	},
}

func init() {
	rootCmd.AddCommand(storageCmd)
	storageCmd.Flags().StringVarP(&Path, "path", "p", "", "Target file path.")
	storageCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Show detail.")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
