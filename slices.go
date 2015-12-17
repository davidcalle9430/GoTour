package main

import(
  "fmt"
  "strings"
)

func main() {
  s := []int {2, 3, 5, 7, 11, 13}
  fmt.Println("s ==", s)

  game := [][]string{
    []string{"_","_","_"},
    []string{"_","_","_"},
    []string{"_","_","_"},
  }

  game[0][0] = "X"
  game[2][2] = "0"
  game[2][0] = "X"
  game[1][0] = "0"
  game[0][2] = "X"

  printBoard(game)

  k := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", k)
	fmt.Println("s[1:4] ==", k[1:4])

	// missing low index implies 0
	fmt.Println("s[:3] ==", k[:3])

	// missing high index implies len(s)
	fmt.Println("s[4:] ==", k[4:])
}

func printBoard(s [][]string){
  for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
