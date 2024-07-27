//声明变量
let personName = "张三"
let personName1:string = "张三"
var personName2 = "李四"

//常量声明(声明之后不能更改)
const tmp = "哈哈"

//声明未赋值 undefined
let num:number
document.write(num + "")

//输出语句
document.write(personName)
document.write(tmp)

//字符串拼接
document.write(`他是${num}`)
document.write("1"+"1")

//number数值类型
//string字符串类型
//boolean 布尔类型
//any任意类型
//array数组
let arr:number[]=[1,2,3,4,5]
//联合类型
let num1:number|string=1

//枚举类型
enum Color{
    red,
    blue,
    green
}
let color:Color=Color.blue
//类型验证
document.write(typeof color)
//类型别名
type NewNumber = number
