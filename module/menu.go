package module

import (
	"./color"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Menu() {

	fmt.Println("/------------- run ascii-art menu -------------\\")
	progressBar(35, 2*60*60*60*60)
	for true {
		result := input(textMenu("start") + "\nselect param")
		if result == "1" {
			menuReverse()
		} else if result == "2" {
			menuOutput(false)
		} else if result == "3" {
			menuPrint()
		} else if result == "e" {
			break
		}
		fmt.Print("result : \n")

	}
	return
}

func menuReverse() {
	for true {
		temp := input(textMenu("reverse") + "\nselect param")
		if temp == "1" {
			for true {
				files := indexingFiles("output")
				temp2 := input("select param")
				if "e" == temp2 {
					break
				} else if v, found := files[temp2]; found {
					for true {
						if len(v) > 24 {
							fmt.Printf("╓────────────────────────────────────╖\n║    file : %s ║\n", v[:11]+"..."+v[len(v)-10:])
						} else {
							fmt.Printf("╓────────────────────────────────────╖\n║    file : %s %s║\n", v, Print(" ", 24-len(v)))
						}
						a := input(textMenu("menu-file") + "\nselect param")
						if a == "1" {
							cat(v)
							fmt.Print("result : ")
							Start([]string{"--reverse=" + v})
							break
						} else if a == "2" {
							b := input("are you sure to delete the file (y/n)")
							if b == "y" {
								_ = os.Remove("output/" + v)
								break
							}
						} else if a == "e" {
							break
						}
					}
				}
			}
		} else if temp == "2" {
			menuOutput(true)
		} else if temp == "e" {
			break
		}
	}
	return
}

func menuOutput(reverse bool) {
	filetype := "standard"
	for true {
		fmt.Print("╓────────────────────────────────────╖\n" +
			"║   select your Parameter -- Output  ║\n" +
			"╟───┬────────────────────────────────╢\n")
		fmt.Printf("║ 1 │  file-type : "+filetype+"%s║\n", Print(" ", 18-len(filetype)))
		fmt.Print("╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Create New file               ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ e │  Exit                          ║\n" +
			"╙───┴────────────────────────────────╜\n")
		temp := input("\nselect param")
		if temp == "1" {
			a := input(textMenu("type-file") + "\nselect param")
			if a == "1" {
				filetype = "standard"
			} else if a == "2" {
				filetype = "shadow"
			} else if a == "3" {
				filetype = "thinkertoy"
			}
		} else if temp == "2" {
			file := input("name file")
			if !strings.HasSuffix(file, ".txt") {
				file += ".txt"
			}
			create := []string{input("enter a string"), filetype, "--output=" + file}
			Start(create)
			cat(file)
			if reverse {
				reverses := []string{"--reverse=" + file}
				fmt.Print("result : ")
				Start(reverses)
				break
			}
		} else if temp == "e" {
			break
		}
	}
}

func menuPrint() {
	colors := "white"
	filetype := "standard"
	align := "left"
	for true {
		fmt.Print("╓────────────────────────────────────╖\n" +
			"║    select your Parameter -- print  ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  color : " + color.MapColor()[colors] + "example" + "\u001b[0m" + "               ║\n" +
			"╟───┼────────────────────────────────╢\n")
		fmt.Printf("║ 2 │  file-type : "+filetype+"%s║\n", Print(" ", 18-len(filetype)))
		fmt.Print("╟───┼────────────────────────────────╢\n")
		fmt.Printf("║ 3 │  align : "+align+"%s║\n", Print(" ", 22-len(align)))
		fmt.Print("╟───┼────────────────────────────────╢\n" +
			"║ 4 │  Print sentence                ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ e │  Exit                          ║\n" +
			"╙───┴────────────────────────────────╜\n")
		temp := input("\nselect param")
		if temp == "1" {
			c := input(textMenu("color") + "\nselect param")
			if c == "1" {
				colors = "white"
			} else if c == "2" {
				colors = "red"
			} else if c == "3" {
				colors = "blue"
			} else if c == "4" {
				colors = "purple"
			} else if c == "5" {
				colors = "green"
			} else if c == "6" {
				colors = "yellow"
			} else if c == "7" {
				colors = "cyan"
			}
		} else if temp == "2" {
			c := input(textMenu("type-file") + "\nselect param")
			if c == "1" {
				filetype = "standard"
			} else if c == "2" {
				filetype = "shadow"
			} else if c == "3" {
				filetype = "thinkertoy"
			}
		} else if temp == "3" {
			c := input(textMenu("align") + "\nselect param")
			if c == "1" {
				align = "left"
			} else if c == "2" {
				align = "right"
			} else if c == "3" {
				align = "center"
			} else if c == "4" {
				align = "justify"
			}
		} else if temp == "4" {
			a := []string{input("sentence to print"), filetype, "--color=" + colors, "--align=" + align}
			fmt.Println("result : ")
			Start(a)
		} else if temp == "e" {
			break
		}
	}
}

func textMenu(num string) string {
	switch num {
	case "start":
		return "╓────────────────────────────────────╖\n" +
			"║        select your Parameter       ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Reverse                       ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Output                        ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 3 │  Print                         ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ e │  Exit                          ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "reverse":
		return "╓────────────────────────────────────╖\n" +
			"║  select your Parameter -- Reverse  ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Existing file                 ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Create file                   ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ e │  Exit                          ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "menu-file":
		return "╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Use file                      ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Delete file                   ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ e │  Exit                          ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "color":
		return "╓────────────────────────────────────╖\n" +
			"║       select your text-color       ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  " + "\u001b[38m" + "White : example " + "\u001b[0m" + "              ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  " + "\u001b[31m" + "Red : example " + "\u001b[0m" + "                ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 3 │  " + "\u001b[34m" + "Blue : example " + "\u001b[0m" + "               ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 4 │  " + "\u001b[35m" + "Purple : example " + "\u001b[0m" + "             ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 5 │  " + "\u001b[32m" + "Green : example " + "\u001b[0m" + "              ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 6 │  " + "\u001b[33m" + "Yellow : example " + "\u001b[0m" + "             ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 7 │  " + "\u001b[36m" + "Cyan : example " + "\u001b[0m" + "               ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "type-file":
		return "╓────────────────────────────────────╖\n" +
			"║        select your text-type       ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Standard                      ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Shadow                        ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 3 │  Thinkertoy                    ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "type-file-reverse":
		return "╓────────────────────────────────────╖\n" +
			"║        select your text-type       ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Auto detect                   ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Standard                      ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 3 │  Shadow                        ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 4 │  Thinkertoy                    ║\n" +
			"╙───┴────────────────────────────────╜\n"
	case "align":
		return "╓────────────────────────────────────╖\n" +
			"║       select your text-align       ║\n" +
			"╟───┬────────────────────────────────╢\n" +
			"║ 1 │  Left                          ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 2 │  Right                         ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 3 │  Center                        ║\n" +
			"╟───┼────────────────────────────────╢\n" +
			"║ 4 │  Justify                        ║\n" +
			"╙───┴────────────────────────────────╜\n"

	}
	return ""
}

func indexingFiles(directory string) map[string]string {
	files := make(map[string]string)
	file, err := ioutil.ReadDir(directory + "/")
	if err != nil {
		log.Fatal(err)
	}
	num := 1
	for _, f := range file {
		if strings.HasSuffix(f.Name(), ".txt") {
			files[strconv.Itoa(num)] = f.Name()
			num++
		}

	}
	files[strconv.Itoa(num)] = "exit"
	fmt.Print("╓────────────────────────────────────╖\n" +
		"║       Existing file -- reverse     ║\n" +
		"╟────┬───────────────────────────────╢\n")
	for i := 1; i <= len(files); i++ {
		if len(files[strconv.Itoa(i)]) < 32 && files[strconv.Itoa(i)] != "exit" {
			fmt.Printf("║ %d%s│ %s %s\n╟────┼───────────────────────────────╢\n", i, Print(" ", 3-len(strconv.Itoa(i))), files[strconv.Itoa(i)], Print(" ", 29-len(files[strconv.Itoa(i)]))+"║")
		} else if files[strconv.Itoa(i)] == "exit" {
			fmt.Printf("║ e  │ Exit                          ║\n╙────┴───────────────────────────────╜\n")
		} else if len(files[strconv.Itoa(i)]) > 32 {
			file := files[strconv.Itoa(i)]
			fmt.Printf("║ %d%s│ %s %s\n╟────┼───────────────────────────────╢\n", i, Print(" ", 3-len(strconv.Itoa(i))), file[:14]+"..."+file[len(file)-12:], "║")
		}
	}
	return files
}

func input(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text, " : ")
	scanner.Scan()
	return scanner.Text()
}

func progressBar(lens int, times time.Duration) {
	fmt.Print("Starting : [", Print(" ", lens), "]\r")
	for i := 0; i < lens; i++ {
		fmt.Print("Starting : [", Print("-", i), ">", Print(" ", lens-i-1), "]\r")
		time.Sleep(times)
	}
	fmt.Print("\n")
}

func Print(carac string, lens int) (result string) {
	for i := 0; i < lens; i++ {
		result += carac
	}
	return
}

func cat(file string) {
	temp, err := ioutil.ReadFile("output/" + file)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
	}
	fmt.Print(string(temp))
}
