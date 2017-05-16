package main

import "fmt"

//The use of defer key word is that it will defer the execution of the hello function and start executing
// rest of the function in the sequence and when the main function is about to exit;It will call the
//defer function but remember it will be rite before the main function is exiting.


//The use of it is in when you want to close the file and you can defer it

/*

THIS IS A SAMPLE WHERE WE CAN USE IT IN OUR CODE.

So the use of it is when i am opening a file.I can close it immediately and can defer it as i might have
to use it...so now when main is exiting it will take care of it.

func main(){
src,err := os.open("src.txt")

if(err != nil){
    os.panic("error")
 }
 defer src.close()

 dest,err := os.open("dest.txt")

if(err != nil){
    os.panic("error")
 }
 defer dest.close()

 bs := make([]byte,5)
 src.read()
 dest.write()
 }
*/

func hello(){
	fmt.Println("hello")
}

func world(){
	fmt.Println(" world")
}

func main(){
	defer hello()
	world()

}