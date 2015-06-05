package main

import "fmt"

func main() {
    fmt.Println(" 整数常量 212 215u 0xFeel 078 032UU \n")
    fmt.Println(" 整数常量 85  0213  0x4b  30  30u  30l  30ul \n ")

    fmt.Println(" 浮点文本常量   3.14159  314159E-5L  510E  210f  .e55 ")
    
    fmt.Println(" 转义序列 ") 
    
    //   \\   \字符
    //   \'   '字符  ?无效转义
    //   \"   "字符
    //   \?   ?字符
    //   \a   警报或钟
    //   \b   退格
    //   \f   换页
    //   \n   换行符
    //   \r   回车
    //   \t   水平制表
    //   \v   垂直制表
    //   \ooo       一到三位数字的八进制数
    //   \xhh...    一个或多个数字的十六进制

    fmt.Println("  \\\\    is   \\   ")
    fmt.Println("  \\'     is   '    ")
    fmt.Println("  \\\"    is   \"   ")
    // fmt.Println("  \\\?    is   \?   ")
    fmt.Println("  \\\a    is   \a   ")
    fmt.Println("  \\\b    is   \b   ")
    fmt.Println("  \\\f    is   \f   ")
    fmt.Println("  \\\n    is   \n   ")
    fmt.Println("  \\\r    is   \r   ")
    fmt.Println("  \\\t    is   \t   ")
    fmt.Println("  \\\v    is   \v   ")
    fmt.Println("  \\\377         is       \377       ")
    fmt.Println("  \\\xffff     is       \xffff   ")

    // \t效果
    fmt.Println("\t\tHello\tWorld")

    fmt.Println(" hello, dear ")
    fmt.Println(" hello, \t dear ")
    fmt.Println(" hello,  ", "d","ear")

    // const
    const LENGTH int = 10
    const WIDTH int = 5
    var area int

    area = LENGTH * WIDTH
    fmt.Println("LENGTH=", LENGTH, "  WIDTH=", WIDTH)
    fmt.Println("area=", area)
    fmt.Printf(" area value %d ",  area)
    fmt.Printf(" area type %T  ", area)


}
