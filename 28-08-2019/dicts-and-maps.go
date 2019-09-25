package main
import ("fmt"; "strconv"; "strings"; "os"; "bufio"; "io")

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
    nrofEntries, _ := strconv.ParseInt(readLine(reader), 10, 64)
    m := make(map[string]string)
    for i:=0; i< int(nrofEntries); i++{
        sli := strings.Fields(readLine(reader))
        key := sli[0]
        value := sli[1]
        m[key] = value
    }


    a := readLine(reader)
    for{
        if len(a) == 0 {
            break
        }
        if value, exists := m[a]; exists{
            fmt.Printf("%s=%s\n",a,value)
        } else{
            fmt.Println("Not found")
        }
        a = readLine(reader)
    }    
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

