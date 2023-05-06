package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)
var (
	path = flag.String("path", " ", "subdomain file")
)
func main() {
    // 打开文件
    flag.Parse()
	var subdomain = []string{}
    f, err := os.Open(*path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    // 创建一个新的 Reader
    reader := csv.NewReader(f)

    // 读取所有记录
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    // 遍历记录并打印第一列数据
    for _, record := range records {
        
		 subdomain=append(subdomain,record[5])
		
    }
	
	for _,sub :=range subdomain[1:]{
		fmt.Println(sub)
	}
}