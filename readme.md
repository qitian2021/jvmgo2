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