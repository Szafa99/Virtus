package cmd

import (
	"errors"
	"log"
	"net/http"
	"os"
	"github.com/spf13/cobra"
)

var  ( 
	path string
	runCmd = &cobra.Command{
	Use:   "run",
	Short: "Serve a html file on the server",
	Long: `Use this command in order to start a http server that servrs a choosen html file .
			Example: virtus run --file index.html`,	
	Run: func (cmd *cobra.Command, args []string){
		if cmd.Flags().Changed("file"){
			if err := run(); err != nil{
				log.Fatal(err)
			}
		}
	} ,
}

)

func run() error {
	// open the file and check for error
	f, err := os.Open(path)
    if err != nil {
        return err
    }
	//at the end close the file
    defer f.Close()
	//check path
    info, err := f.Stat()
    if err != nil {
        return err
    }

	// check if given path is a directory
    if info.IsDir() {
		return errors.New("can not serve content of a directory")
	}

	// if ok, keep the path
    path = info.Name()


			http.HandleFunc("/", readFile)
			
			log.Println("Listening on :8080")
			err = http.ListenAndServe(":8080", nil)
			if err != nil{
				log.Fatal(err)
				return err
			}
		return nil 
}

// serve the file
func readFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, path)
}

// add comand and flags
func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&path,"file","f","","path to html file")
	runCmd.MarkFlagRequired("file")
}
