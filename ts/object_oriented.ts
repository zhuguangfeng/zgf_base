//面向对象

//类
class Person{
    static des:string="这是一个person类" //类属性
    static test(){ //类方法

    }
    name: string="默认";
    age: number=0;

    //构造方法
    constructor(name:string,age:number){
        this.name=name
        this.age=age
    }
    say(){
        document.write(this.name)
    }
}



//实例化对象
let a = new Person("张三",10);
a.name="张三";
a.age=10;
a.say();
