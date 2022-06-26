/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/andriiluk/workouts/internal/musclesvc"
	"github.com/spf13/cobra"
)

const (
	EnvLogLevel = "WORKOUTSVC_LOG_LEVEL"
)

// runCmd represents the run command
var (
	logLevels = map[string]log.Level{
		"warn":  log.WarnLevel,
		"info":  log.InfoLevel,
		"debug": log.DebugLevel,
	}

	addr   string
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run Workout HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			RunWorkoutSvc()
		},
	}
)

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&addr, "addr", "a", ":8080", "-a localhost:8080")

	lvl, ok := logLevels[strings.ToLower(os.Getenv(EnvLogLevel))]
	if !ok {
		lvl = log.InfoLevel
	}

	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp:  true,
		PrettyPrint:       true,
		DisableHTMLEscape: true,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d", f.File, f.Line)
		},
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(lvl)
	log.SetReportCaller(true)

	log.WithField("log_level", lvl).Info()
}

func RunWorkoutSvc() {
	log.WithField("addr", addr).Info("Run workout service")

	muscleStore := musclesvc.NewInMemStore()
	muscleSvc := musclesvc.NewService(muscleStore)
	muscleSvc = musclesvc.WithLoggingMidleware(muscleSvc)

	muscleHandler := musclesvc.MakeHTTPHandler(muscleSvc)

	log.Fatal(http.ListenAndServe(addr, muscleHandler))
}
