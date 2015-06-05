package main

import "fmt"

func main() {
  fmt.Println(" \n\n\t  算术运算符   ")
  fmt.Println("  +     两个数相加  ")
  fmt.Println("  -     第一个数 减 第二个数  ")
  fmt.Println("  *     两个数相乘  ")
  fmt.Println("  /     分子 除以 分母  ")
  fmt.Println("  %     模运算和整数除法后的余数  ")
  fmt.Println("  ++    运算符递增, 整数值加一   ")
  fmt.Println("  --    运算符递减, 整数值减一   ")


  fmt.Println(" \n\n\t  关系运算符   ")
  fmt.Println("  ==     检查两个操作数值是否相等, 相等的话条件为 true    ")
  fmt.Println("  !=     检查两个操作数是否不等,  不等的话条件为 true     ")
  fmt.Println("  >      检查左边的操作数是否大于右边的操作数, 左 大于 右 为 true  ")
  fmt.Println("  <      检查左边的操作数是够小于右边的操作数, 左 小于 右 为 true  ")
  fmt.Println("  >=     检查左边的操作数是否大于等于右边的操作数, 左 大于或等于 右 为 true   ")
  fmt.Println("  <=     检查左边的操作数是否小于等于右边的操作数, 左 小于或等于 由 为 true   ")


  fmt.Println(" \n\n\t  逻辑运算符   ")
  fmt.Println("  &&      逻辑与运算符, 两个操作数都非0, 则为 true ")
  fmt.Println("  ||      逻辑或运算符, 两个操作数任一一个是非0, 则为 true   ")
  fmt.Println("  !       逻辑非运算符, 反转操作数的逻辑状态   ")


  fmt.Println(" \n\n\t  位运算符     ")
  fmt.Println(" 位运算符真值表  ")
  fmt.Println(" |   p   |   q   |   p & q   |   p | q   |   p ^ q   | ")
  fmt.Println(" |   0   |   0   |     0     |     0     |     0     | ")
  fmt.Println(" |   0   |   1   |     0     |     1     |     1     | ")
  fmt.Println(" |   1   |   1   |     1     |     1     |     0     | ")
  fmt.Println(" |   1   |   0   |     0     |     1     |     1     | ")

  fmt.Println(" A = 60, b = 13 则 A= 0011 1100, B = 0000 1101 ")
  fmt.Println("  A&B = 0000 1100  ")
  fmt.Println("  A|B = 0011 1101  ")
  fmt.Println("  A^B = 0011 0001  ")
  fmt.Println("  ~A  = 1100 0011  ")

  fmt.Println("  &   二进制与 操作副本的结果. (A & B) = 12 , 也就是 0000 1100 ")
  fmt.Println("  |   二进制或 操作副本的结果. (A | B) = 61 , 也就是 0011 1101 ")
  fmt.Println("  ^   二进制异或 操作副本的结果. (A ^ B) = 49, 也就是 0011 0001 ")
  fmt.Println("  <<  二进制左移位运算符, 左边的操作数的值向左 左移位数由右操作数指定位数.  A << 2 是240, 也就是 1111 0000 ")
  fmt.Println("  >>  二进制右移位运算符, 左边的操作数的值向右 右移位数由右操作数指定位数.  A >> 2 = 15, 也就是 0000 1111 ")


  fmt.Println(" \n\n\t  赋值运算符  ")
  fmt.Println("   ==    简单的赋值操作符, 分配值从右侧的操作数到左侧的操作数  C = A + B, 分配 A + B的值到C")
  fmt.Println("   +=    相加并赋值运算符, 左侧操作数与右侧操作数之和分配结果给左侧操作数 C += A, 相当于 C = C + A  ")
  fmt.Println("   -=    减法赋值运算符, 从左侧操作数减去右侧操作数的值分配给左操作数, C = -= A 相当于 C = C - A  ")
  fmt.Println("   *=    乘法运算操作符, 左侧操作数与右侧操作数相乘的值分配给左侧操作数, C *= A 相当于 C = C * A  ")
  fmt.Println("   /=    除法赋值运算符, 左侧操作数与右侧操作数相除的值分配给左侧操作数. C /= A 相当于 C = C / A   ")
  fmt.Println("   %=    模赋值运算符, 左侧操作数与右侧操作数模运算的值分配给左操作数, C %= A 相当于 C = C % A ")
  fmt.Println("   <<=   左移位并赋值操作符, C << 2 相当于 C = C << 2  ")
  fmt.Println("   >>=   右移位并赋值操作符, C >> 2 相当于 C = C >> 2  ")
  fmt.Println("   &=    按位与赋值运算符, C &= 2 相当于 C = C & 2   ")
  fmt.Println("   ^=    按位异或赋值运算符, C ^= 2 相当于 C = C ^ 2   ")
  fmt.Println("   |=    按位或赋值运算符, C |= 2 相当于  C = C | 2    ")


  fmt.Println(" \n\n\t  其它运算符 ")
  fmt.Println("  &   返回一个变量地址, &a 将得到变量a 的实际地址 ")
  fmt.Println("  *   指针的变量, *a 将指向变量a ")
   

}