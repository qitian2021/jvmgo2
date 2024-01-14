# java命令
* Java虚拟机规范没有明确规定，是有虚拟机实现自行决定的。java命令有如下四种形式：

* java [-options] class [args]

* java [-options] -jar jarfile [args]

* javaw [-options] class [args]

* javaw [-options] -jar jarfile [args]

* 选项分为两类：标准选项和非标准选项  https://docs.oracle.com/javase/8/docs/

  ## 可以向java命令传递三组参数：选项、主类名（或者JAR文件名）和main（）方法参数



# 搜索class文件

### Java虚拟机实现根据类路径（class path）来搜索类。由搜索先后顺序，类路径分为3个部分：

#### 启动类路径（bootstrap classpath）---jre\lib 目录

#### 扩展类路径（extension classpath）---jre\lib\ext 目录

#### 用户类路径（user classpath）默认当前目录 也就是 “.”

##### 可以使用 java 命令 -classpath(简写 -cp) 可以覆盖CLASSPATH 环境变量设置来修改用户类路径

-classpath/-cp 选项既可以指定目录，也可以指定JAR文件或者ZIP文件：

java -cp path\to\classes ...

java -cp path\to\lib1.jar ....

java -cp path\to\lib2.zip ...



注意：Go函数或方法允许返回多个值，按照惯例，可以使用最后一个返回值作为错误信息

#### CompositeEntry

compositeEntry由更小的Entry组成，正好可以表示成[]Entry,在Go语言中，数组属于比较低层的数据结构很少直接使用。大部分情况下，使用更便利的slice类型。

#### ch01代码测试结果

![ch01测试代码结果](D:\go\workspace\src\jvmgo\imgdata\ch01测试代码结果.png)

第二章:项目目录结构为：

![ch02项目结构](D:\go\workspace\src\jvmgo\imgdata\ch02项目结构.png)

测试代码 ：CMD直接进入你的ch02目录下，执行go install 

然后你的bin 目录下会生成 ch02.exe ,CMD 进入你的bin目录下 使用命令： 

```powershell
.\ch02.exe -Xjre "D:\java\jdk1.8.0_202\jre" java.lang.Object 
```

测试结果：

![ch02测试结果](D:\go\workspace\src\jvmgo\imgdata\ch02测试结果.png)

### 第三章： 解析class文件

Java虚拟机规范中所指的class文件，并非特指位于磁盘中的.class文件,而是泛指任何格式符合规范的class数据

#### 魔数

特定格式的文件必须以某几个固定字节开头，这几个字节主要起标识作用，叫做魔数（magic number）

例如：PDF文件以4字节"%PDF" 开头，ZIP文件以2字节"PK"(0x50、0x4B)开头

#### class文件的魔数是："0xCAFEBABE"

Java虚拟机规范规定：如果加载的class文件不符合要求的格式，Java虚拟机实现就抛出java.lang.ClassFormatError异常

#### 版本号

魔数之后就是class文件的次版本号和主版本号，都是u2类型

JavaSE 8支持版本号为45.0~52.0的class文件。

如果版本号不在支持的范围内，Java虚拟机实现就抛出java.lang.UnsupportClassVersionError异常

#### 类访问标志

16位的bitmask,指出class文件定义的是类还是接口，访问级别是public还是private.

#### 类和超类索引

类访问标志之后是两个u2类型的常量池索引，分别给出类名和超类名

#### 接口索引表

类和超类索引后面是接口索引表，表中存放的也是常量池索引，给出该类实现的所有接口的名字

#### 字段和方法表

分别存储字段和方法信息。字段和方法的基本结构大致相同，差别仅在于属性表

#### 常量池

常量池占据class文件很大的一部分数据，里面存放着各式各样的常量信息，包括数字和字符串常量、类和接口名称、字段和方法名

