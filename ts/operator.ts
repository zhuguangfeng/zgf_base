let num :number=10
//运算符
// + — * / % ++ --
//++在前先自增在使用 ++在后先使用在自增
document.write(num++ + "")
document.write(++num + "")
//==只比较值  ===比较类型和值

//三目运算符
num = num>100 ? 1 : 2

//while循环
let i:number=0;
while (i<5){
    document.write(i + "")
    i++
}

//do while循环
do {
    document.write(i + "")
}while(i<0)

//for循环
for (let i =0;i<10;i++){
}
//直接获取元素的值
let names:string[]=["1","2","3"]
for (let item of names){
    document.write(item)
}
//获取下标
for (let index in names){
    document.write(index)
}
