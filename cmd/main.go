package main

import (
	"fmt"
	"os"
	"time"

	"github.com/simplesvet/dbexport"
	"github.com/spf13/cobra"
)

var objName string
var timer int

func ExportDbObjectsCmd(objectType string, objectName string, silent bool) int {
	if !silent {
		fmt.Println("Exporting", objectType)
	}

	var dbObjects []dbexport.DbObject

	if objectType == dbexport.PROCEDURES {
		dbObjects = dbexport.GetProcedures(objectName)
	} else if objectType == dbexport.FUNCTIONS {
		dbObjects = dbexport.GetFunctions(objectName)
	} else if objectType == dbexport.TRIGGERS {
		dbObjects = dbexport.GetTriggers(objectName)
	} else if objectType == dbexport.EVENTS {
		dbObjects = dbexport.GetEvents(objectName)
	} else if objectType == dbexport.VIEWS {
		dbObjects = dbexport.GetViews(objectName)
	} else if objectType == dbexport.TABLES {
		dbObjects = dbexport.GetTables(objectName)
	} else {
		dbObjects = dbexport.GetAll()
	}

	if len(dbObjects) == 0 {
		fmt.Println("nenhum objeto encontrado, revise o arquivo de conexão com o banco")
		return 0
	}

	savedFiles := dbexport.SaveDbObjects(dbObjects)

	if len(savedFiles) == 0 {
		fmt.Println("nenhum resultado encontrado, revise o arquivo de conexão com o banco")
		return 0
	}

	if !silent {
		for _, file := range savedFiles {
			fmt.Println("File saved in", file)
		}
	}

	return len(savedFiles)
}

func Observe(interval int) {
	fmt.Println("listening database")
	fmt.Println("press ctrl+c to stop")
	for {
		duration := time.Duration(interval) * time.Second
		time.Sleep(duration)
		fmt.Println("scaning database...")
		savedObjects := ExportDbObjectsCmd("", "", true)
		fmt.Printf("saved %d objects\n", savedObjects)
	}
}

var rootCmd = &cobra.Command{
	Use:   "dbexport all | dbexport [object_type] [object_name] | dbexport [object_type] --name object_name",
	Short: "DBExport is a fast tool to sync databases objects with the file system",
	Long:  `DBExport is a fast tool to sync databases objects with the file system`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		firstArg := args[0]

		if len(args) == 2 {
			objName = args[1]
		}

		if firstArg == "help" {
			cmd.Help()
		} else if firstArg == "all" {
			ExportDbObjectsCmd("", "", false)
		} else {
			ExportDbObjectsCmd(firstArg, objName, false)
		}
	},
}

var observeCmd = &cobra.Command{
	Use:   "observe --interval [intervalo] | i [intervalo]",
	Short: "Observe listen get changes in your database",
	Long:  `Observe listen get changes in your database`,
	Args:  cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		firstArg := ""
		if len(args) > 0 {
			firstArg = args[0]
		}

		if firstArg == "help" {
			cmd.Help()
		} else {
			Observe(timer)
		}
	},
}

func main() {
	dbexport.GetConfig()
	rootCmd.Flags().StringVarP(&objName, "name", "n", "", "nome do objeto no banco")
	observeCmd.Flags().IntVarP(&timer, "interval", "i", 5, "tempo de espera para baixar atualizações")

	rootCmd.AddCommand(observeCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
