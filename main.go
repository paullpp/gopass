/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	"golang.design/x/clipboard"

	"math/rand"
	"strconv"

	"github.com/spf13/cobra"
)

func main() {
	var numbers, special, all bool

	var rootCmd = &cobra.Command{
		Use:   "gopass",
		Short: "Generate passwords",
		Long: `Usage: ./gopass <length_of_passwords> 
		
	Flags: 
	 -n: use numbers
	 -s: use special characters
	 -A: use letters, numbers, and special characters`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := clipboard.Init()
			if err != nil {
				panic(err)
			}
			passLen, err := strconv.Atoi(args[0])
			fmt.Println("=== Generating Password of length ", passLen)
			if err != nil {
				fmt.Println("=== Error: ", err)
			} else {
				if all {
					special = true
					numbers = true
				}
				pass := genPassword(special, numbers, passLen)
				clipboard.Write(clipboard.FmtText, []byte(pass))
				fmt.Println(pass)
			}
		},
	}

	rootCmd.Flags().BoolVarP(&numbers, "numbers", "n", false, "Include numbers in the password")
	rootCmd.Flags().BoolVarP(&special, "special", "s", false, "Include special characters in the password")
	rootCmd.Flags().BoolVarP(&all, "all", "A", false, "Include special characters and numbers in the password")

	rootCmd.Execute()
}

func genPassword(special, numbers bool, passLen int) string {
	var allChars []string
	var specialChars = []string{"[", "]", "(", ")", ",", ".", "!", "@", "#", "$", "%", "^", "&", "*", "-", "_", "+", "=", "?", "<", ">"}
	var pass string

	if numbers {
		for i := '0'; i <= '9'; i++ {
			allChars = append(allChars, string(i))
		}
	}
	if special {
		allChars = append(allChars, specialChars...)
	}
	for i := 'a'; i <= 'z'; i++ {
		allChars = append(allChars, string(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		allChars = append(allChars, string(i))
	}

	for i := 0; i < passLen; i++ {
		var idx = rand.Intn(len(allChars))
		pass += string(allChars[idx])
	}

	return pass
}
