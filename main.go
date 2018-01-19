// Copyright 2017 NDP Systèmes. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	hexyacmd "github.com/hexya-erp/hexya/cmd"
	"github.com/hexya-erp/hexya/hexya/server"
	"github.com/hexya-erp/hexya/hexya/tools/generate"
	"github.com/hexya-erp/hexya/hexya/tools/logging"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/halybang/haxey/config"
)

var log *logging.Logger

func init() {
	log = logging.GetLogger("init")
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringP("config", "c", "", "Alternate configuration file to read. Defaults to $HOME/.hexya/")
	RootCmd.PersistentFlags().StringSliceP("modules", "m", []string{"github.com/hexya-erp/hexya-base/web"}, "List of module paths to load. Defaults to ['github.com/hexya-erp/hexya-base/web']")

	RootCmd.PersistentFlags().StringP("log-level", "L", "info", "Log level. Should be one of 'debug', 'info', 'warn', 'error' or 'crit'")
	RootCmd.PersistentFlags().String("log-file", "", "File to which the log will be written")
	RootCmd.PersistentFlags().BoolP("log-stdout", "o", false, "Enable stdout logging. Use for development or debugging.")
	RootCmd.PersistentFlags().Bool("debug", false, "Enable server debug mode for development")

	RootCmd.PersistentFlags().String("data-dir", "", "Path to the directory where Hexya should store its data")
	RootCmd.PersistentFlags().String("root-dir", "", "Path to the directory where Hexya root dir")

	RootCmd.PersistentFlags().String("db-driver", "postgres", "Database driver to use")
	RootCmd.PersistentFlags().String("db-host", "/var/run/postgresql",
		"The database host to connect to. Values that start with / are for unix domain sockets directory")
	RootCmd.PersistentFlags().String("db-port", "5432", "Database port. Value is ignored if db-host is not set")
	RootCmd.PersistentFlags().String("db-user", "", "Database user. Defaults to current user")
	RootCmd.PersistentFlags().String("db-password", "", "Database password. Leave empty when connecting through socket")
	RootCmd.PersistentFlags().String("db-name", "hexya", "Database name")

	viper.BindPFlag("ConfigFileName", RootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("Modules", RootCmd.PersistentFlags().Lookup("modules"))

	viper.BindPFlag("LogLevel", RootCmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("LogFile", RootCmd.PersistentFlags().Lookup("log-file"))
	viper.BindPFlag("LogStdout", RootCmd.PersistentFlags().Lookup("log-stdout"))
	viper.BindPFlag("Debug", RootCmd.PersistentFlags().Lookup("debug"))

	viper.BindPFlag("DataDir", RootCmd.PersistentFlags().Lookup("data-dir"))
	viper.BindPFlag("RootDir", RootCmd.PersistentFlags().Lookup("root-dir"))

	viper.BindPFlag("DB.Driver", RootCmd.PersistentFlags().Lookup("db-driver"))
	viper.BindPFlag("DB.Host", RootCmd.PersistentFlags().Lookup("db-host"))
	viper.BindPFlag("DB.Port", RootCmd.PersistentFlags().Lookup("db-port"))
	viper.BindPFlag("DB.User", RootCmd.PersistentFlags().Lookup("db-user"))
	viper.BindPFlag("DB.Password", RootCmd.PersistentFlags().Lookup("db-password"))
	viper.BindPFlag("DB.Name", RootCmd.PersistentFlags().Lookup("db-name"))

	serverCmd.PersistentFlags().StringP("interface", "i", "", "Interface on which the server should listen. Empty string is all interfaces")
	serverCmd.PersistentFlags().StringP("port", "p", "8080", "Port on which the server should listen.")
	serverCmd.PersistentFlags().StringSliceP("languages", "l", []string{}, "Comma separated list of language codes to load (ex: fr,de,es).")
	viper.BindPFlag("Server.Interface", serverCmd.PersistentFlags().Lookup("interface"))
	viper.BindPFlag("Server.Port", serverCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("Server.Languages", serverCmd.PersistentFlags().Lookup("languages"))

	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(updateDBCmd)
	RootCmd.AddCommand(serverCmd)
	RootCmd.AddCommand(linkCmd)
}

// RootCmd is the base 'haxey' command of the commander
var RootCmd = &cobra.Command{
	Use:   "haxey",
	Short: "Haxey is an open source app wrapper for Hexya ERP",
	Long: `Haxey is an open source app wrapper for Hexya ERP written in Go.
It is designed for high demand business data processing while being easily customizable`,
	Run: func(cmd *cobra.Command, args []string) {
		hexyacmd.StartServer(viper.AllSettings())
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version haxey",
	Long:  `Print the version of the haxey application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("haxey version 0.0.1")
	},
}

var updateDBCmd = &cobra.Command{
	Use:   "updatedb",
	Short: "Update the database schema",
	Long:  `Synchronize the database schema with the models definitions.`,
	Run: func(cmd *cobra.Command, args []string) {
		hexyacmd.UpdateDB(viper.AllSettings())
	},
}

var serverCmd = &cobra.Command{
	Use:   "server [projectDir]",
	Short: "Start the Hexya server",
	Long: `Start the Hexya server of the project in 'projectDir'.
If projectDir is omitted, defaults to the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		hexyacmd.StartServer(viper.AllSettings())
	},
}

var symlinkDirs = []string{"static", "templates", "data", "resources", "i18n"}

func cleanSymLinks() {
	rootDir := viper.GetString("RootDir")
	if rootDir == "" {
		return
	}
	for _, dir := range symlinkDirs {
		dirPath := filepath.Join(rootDir, "hexya", "server", dir)
		os.RemoveAll(dirPath)
		os.MkdirAll(dirPath, 0775)
		// => ${ROOT}/hexya/server/{static,data,i18n,resources,templates}
	}
}

func createSymLinks() {
	rootDir := viper.GetString("RootDir")
	if rootDir == "" {
		return
	}
	for _, mod := range server.Modules {
		for _, dir := range symlinkDirs {
			srcPath := filepath.Join(mod.Dir, dir)
			dstPath := filepath.Join(rootDir, "hexya", "server", dir, mod.Name)
			if _, err := os.Stat(srcPath); err == nil {
				//os.RemoveAll(dstPath)
				fmt.Println("CreateSymLinks", srcPath, " <= ", dstPath)
				os.Symlink(srcPath, dstPath)
			}
		}
	}
}

var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Update symlinks to root-dir.",
	Long:  `Update symlinks to root-dir.`,
	Run: func(cmd *cobra.Command, args []string) {
		cleanSymLinks()
		createSymLinks()
	},
}

func main() {

	// Maximize goroutines
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig() {
	cfgFile := viper.GetString("ConfigFileName")

	if runtime.GOOS != "windows" {
		viper.AddConfigPath("/etc/gut")
	}

	osUser, err := user.Current()
	if err != nil {
		log.Panic("Unable to retrieve current user", "error", err)
	}
	defaultHexyaDir := filepath.Join(osUser.HomeDir, ".gut")
	viper.SetDefault("DataDir", defaultHexyaDir)

	rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err == nil {
		viper.SetDefault("RootDir", rootDir)
	}

	viper.AddConfigPath(defaultHexyaDir)
	viper.AddConfigPath(".")

	viper.SetConfigName("hexya")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config: ", err)
	}
	rootDir = viper.GetString("RootDir")
	if _, err := os.Stat(rootDir); err != nil {
		os.MkdirAll(rootDir, 0755)
	}
	// Comment out below line for orgin code
	generate.HexyaDir = rootDir
	log.Info(fmt.Sprintf("RootDir %s", generate.HexyaDir))
}
