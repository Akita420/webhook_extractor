package main
import(
  "os"
  "fmt"
  "regexp"
  "io/ioutil"
)

func main(){
  var f []byte

  fmt.Printf("\nDiscord webhook exctractor made by Akita\n\n")

  if len(os.Args) != 2{//check if arguments have been applied
    fmt.Printf("Usage: %s <filename>\n",os.Args[0])
    os.Exit(1)
  }

  file,err := ioutil.ReadFile(os.Args[1])
  if err!=nil{//if cannot read file exit
    fmt.Printf("Encountered an error while reading the file...\nError: %v\n",err)
    os.Exit(1)
  }

  for i:=0;i<len(file);i++{//strip 0 bytes
    if file[i] == byte(0){
      continue
    }
    f = append(f,file[i])
  }

	re,err := regexp.Compile(`(https:\/\/discordapp.com\/api\/webhooks\/)[0-9]{18}\/.{68}`)
  if err!=nil{//if regex not work exit
    fmt.Printf("Regex Error: %v\n",err)
    os.Exit(1)
  }

  fmt.Printf("===== WEBHOOK =====\n\n%s\n\n===== WEBHOOK =====\n",string(re.Find(f)))//print result

}
