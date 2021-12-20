/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// greetCmd represents the greet command
var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Provide <greeting> and <name>",
	Long: `Provide <greeting> and <name>:
	Example : bbtools greet hello John
	Output  : [1] "hello John"   `,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(args[0])
		callDocker(args)
	},
}

func callDocker(args []string) {
	cmd := exec.Command("docker", "run", "curious98/mini-r", "Rscript", "hello.R", args[0], args[1])
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))
	// out, err := exec.Command("bash", "list.sh").Output()

	// // if there is an error with our execution
	// // handle it here
	// if err != nil {
	// 	fmt.Printf("%s", err)
	// }
	// // as the out variable defined above is of type []byte we need to convert
	// // this to a string or else we will see garbage printed out in our console
	// // this is how we convert it to a string
	// fmt.Println("Command Successfully Executed")
	// output := string(out[:])
	// fmt.Println(output)
}
func init() {
	rootCmd.AddCommand(greetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// greetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// greetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
