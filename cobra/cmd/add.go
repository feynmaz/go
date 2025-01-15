/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add adds two numbers",
	Long:  `add adds two numbers of type int or float`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInt(args []string) {
	var sum int

	for _, value := range args {
		number, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + number
	}
	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, value := range args {
		number, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + number
	}
	fmt.Printf("Sum of floating numbers %s is %f\n", args, sum)
}
