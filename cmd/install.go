/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "installs and builds docker file",
	Long: `1)check if docker exists or not
	2)Pulls docker image from docker hub.
	3)Builds docker image
	4)Creates local registry
	5)pushed docker image to local registry`,
	Run: func(cmd *cobra.Command, args []string) {
		checkDocker()
	},
}

func checkDocker() {
	colorGreen := "\033[32m"
	colorRed := "\033[31m"

	cmd := exec.Command("docker")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
		fmt.Println(string(colorRed), "Docker is not installed. Please install docker and then run this command again.")
		os.Exit(1)
	}
	// fmt.Println(string(output))
	fmt.Println(string(colorGreen), "Docker installation verified.")
	fmt.Println(string(colorGreen), "Pulling docker image")
	dockerPull()
	fmt.Println(string(colorGreen), "Building docker image. Please wait")
	buildDocker()
	fmt.Println(string(colorGreen), "Docker build is complete.")
	fmt.Println(string(colorGreen), "remove any docker container with name registry")
	removeDocker()
	pushDockerToLocalRegistry()
	fmt.Println(string(colorGreen), "Docker pushed to local registry")
	fmt.Println(string(colorGreen), "Now you can run : bbtools greet <Greeting> <Your_Name>")

}

func removeDocker() {
	cmd := exec.Command("docker", "rm", "-f", "registry")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + string(output))
		return
	}
	fmt.Println(string(output))

}

func dockerPull() {
	// colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorWhite := "\033[37m"

	cmd := exec.Command("docker", "pull", "curious98/mini-r")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
		fmt.Println(string(colorRed), "Docker not able to fetch the image from Docker Hub. Check your network connection and try again later!")
		os.Exit(1)
	}
	fmt.Println(string(colorWhite), string(output))
}
func buildDocker() {
	// colorGreen := "\033[32m"
	colorRed := "\033[31m"

	cmd := exec.Command("docker", "run", "curious98/mini-r")
	output, err := cmd.CombinedOutput()
	if err != nil {
		if !(strings.Contains(string(output), "--save")) {
			fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
			fmt.Println(string(colorRed), "Docker Build error. Maybe update your docker and try again.")
			os.Exit(1)
		}
	}
	// fmt.Println(string(output))
}

func pushDockerToLocalRegistry() {

	colorGreen := "\033[32m"
	colorRed := "\033[31m"
	colorWhite := "\033[37m"

	//creating the local registry
	fmt.Println("creating the local registry")
	cmd := exec.Command("docker", "run", "-d", "-p", "5000:5000", "--restart=always", "--name", "registry", "registry:2")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
		fmt.Println(string(colorRed), "Fetching registry failed. Check your network connection and try again.")
		os.Exit(1)
	}
	fmt.Println(string(colorWhite), string(output))

	//creating a image with a new tag mini-r
	fmt.Println("creating a image with a new tag mini-r")
	cmd = exec.Command("docker", "tag", "curious98/mini-r", "localhost:5000/mini-r")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
		fmt.Println(string(colorRed), "Error creating a new tag for the image to be pushed.")
		os.Exit(1)
	}
	fmt.Println(string(colorWhite), string(output))

	//pushing the image to local registry
	fmt.Println(string(colorGreen), "pushing the image to local registry")
	cmd = exec.Command("docker", "push", "localhost:5000/mini-r")
	output, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(colorRed), fmt.Sprint(err)+": "+string(output))
		fmt.Println(string(colorRed), "Error pushing to the local registry.")
		os.Exit(1)
	}
	fmt.Println(string(colorWhite), string(output))

}
func init() {
	rootCmd.AddCommand(installCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
