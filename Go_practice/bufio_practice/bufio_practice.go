package bufio_practice

import(
	"fmt"
	"bufio"
	"log"
	"strings"
	"os"
	
)

func Readers(){
	fmt.Print("Enter your name: ")

	r := bufio.NewReader(os.Stdin)

	name, err := r.ReadString('\n')

	if err != nil{
		log.Fatal(err)
	}
	fmt.Printf("Hello %s!) \n", strings.TrimSpace(name))
}
func Writers(){
	data := []string{"an old falcon","misty mountains","A wise man", "a rainy morning"}

	f, err := os.Create("words.txt")
	if err != nil{
		log.Fatal(err)
	}

	defer f.Close()
	wr := bufio.NewWriter(f)

	for _, line := range data {
		wr.WriteString(line+"\n")
	}
	wr.Flush()
	fmt.Println("data written")
}	

func ReadFileLinesScanner(){
	f, err := os.Open("words1.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}