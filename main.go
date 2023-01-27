package main

import "fmt"

/*1. package name and file name should be same
2. There cannot be multiple files with same name (obviously!!!)
3. "go mod init example.com" : The go mod init command initializes and writes a new go.mod file in the current directory, in effect creating a new module rooted at the current directory. The go.mod file must not already exist. We use this to resolve/maintain our dependencies when our project becomes large.
4. "go mod tidy" : go mod tidy ensures that the go.mod file matches the source code in the module. It adds any missing module requirements necessary to build the current module’s packages and dependencies, and it removes requirements on modules that don’t provide any relevant packages. It also adds any missing entries to go.sum and removes unnecessary entries.
5. The red lines after a variable is defined, this is a way of go compiler to tell you that you have not used the variable any where, thus informing you that you are wasting the memory
*/
func main() {
	/*fmt.Print("Hello World")    //No new line
	fmt.Println("Hello World2") //Prints then new line

	name := "Aaryan" // variable name is declared with string value as Aaryan
	fmt.Printf("hello %v", name)
	// %d -> int
	//  %f -> float
	//  %s -> string
	// %w -> error (It is a data type in Go)
	// %v -> any data (You can use it when you are confused)

	name1 := "aaryan"
	name2 := "tarun"
	name3 := "subhesh"
	name4 := "nabin"

	fmt.Printf("\nhello %v %v %v %v\n", name1, name2, name3, name4)
	fmt.Printf("\nhello %[1]v %[1]v %[1]v %[1]v\n", name1) //%[1]v means just take the first argument

	//Defining Variables in Go
	// 1. Old school way

	var name5 string = "tarun"
	fmt.Print(name5)
	// 2. New and interesting way
	name6 := "nabin"
	fmt.Print(name6)
	*/
	greetingMessage := Greeting("tarun")
	fmt.Println(greetingMessage)
	v1 := GreetingNamedOutputParam("rahul") // := is used for intialisation
	v1 = GreetingNamedOutputParam("rohit")  // after initialization we can use = normally for updating
	fmt.Println(v1)

}

func Print(input string) {
	fmt.Println(input)
}

func Greeting(personName string) string {
	return "hello " + personName
}

func GreetingNamedOutputParam(personName string) (greetingMessage string) {
	if personName == "tarun" {
		greetingMessage = " hello sir"
		return
	}
	greetingMessage = "hello " + personName
	return

}

// GreetingWithVarNumOfParam("tarun","rahul", "kanit", "asdsfa")
// func GreetingWithVarNumOfParam(personName ...string) (greetingMessage []string){

// }
