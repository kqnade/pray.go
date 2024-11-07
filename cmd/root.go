/*
Copyright © 2024 Yuzuki Kana contact@k4na.de

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"os"
  "fmt"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pray",
	Short: "Summon shrine for pray",
	Long: `The purpose of the pray is to display the ASCII art of the prayer target in the terminal. This package currently includes the ability to display a Shinto altar as ASCII art, providing a unique and visually interesting way to bring symbolic representation to the terminal.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("  __________________________________________________________")
    fmt.Println("  \\________________________________________________________/")
    fmt.Println("       |      |             |    |             |      | ")
    fmt.Println("     __|      |_____________|____|____________ |      |__ ")
    fmt.Println("    |__|      |________________________________|      |__|")
    fmt.Println("       |      |   _______//  /_||_\\  \\______   |      |")
    fmt.Println("       |      |  /::::://  /|| || ||\\  \\:::::\\ |      |")
    fmt.Println("       |      |//:::://   ~~~~~~~~~~~   \\:::::\\|      |")
    fmt.Println("      /|      | ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ |      |\\")
    fmt.Println("    /::|      | ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ |      |::\\")
    fmt.Println("   ~~~ |      || ~~~~~~~~~~~~~(P)~~~~~~~~~~~~ ||      |~~~")
    fmt.Println("       |      ||      *      / S \\      *    .||      |")
    fmt.Println("       |      || __________/   S   \\_________ ||      |")
    fmt.Println("       |      || :::||+++             +++|::: ||      |")
    fmt.Println("       |      || :::||+++ ___________ +++|::: ||      |")
    fmt.Println("       |      || :::||++  |*********|  ++|::: ||      |")
    fmt.Println("      |~~~~~~~~~|     ==================     |~~~~~~~~~|")
    fmt.Println("  ::::::::::::::::::======================:::::::::::::::::::")
    fmt.Println("  :::::::::::::::::/                      \\::::::::::::::::::")
    fmt.Println("  ::::::::::::::::/                        \\:::::::::::::::::")
    fmt.Println("  :::::::::::::::/                          \\::::::::::::::::")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.pray.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


