package cmd1

import (
	"fmt"

	"github.com/spf13/cobra"
	
)
 
var(
	localRootFlag bool
	rootCmd = &cobra.Command{
		Use: "Hi. This is root command ",
		Short: "Test the cobra program",
		Run: func(cmd *cobra.Command, args []string){
			fmt.Print("Hello from the root command")
		},
	}
)
 
func init(){
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag" )
}
 
func Execute(){
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}

	
}

