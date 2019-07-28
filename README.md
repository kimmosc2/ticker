Ticker  
=============================
 
## Ticker是什么?  
过去,当我们在命令行下做一些耗费时间的工作时(如npm install等),等待命令完成的这段时
间是十分无趣的,当我们切换至其他任务时可能许久之后才记起这个"不知道什么时候会完成"的
任务,ticker配合其他命令使用,可以以各种方式提醒您任务完成.(作者:没什么用)  

## 示例
1.配合ping命令使用不加参数的ticker
```
ping example.com & ticker
```
在ping命令完成后,不加参数的tiker会输出响铃符(\a),如果你的电脑开着声音,会听
到"叮"一声(windows版本不同声音也不同)  
![](https://github.com/kimmosc2/ticker/raw/master/assets/image/p2.png)  
  
  
2.配合ping命令使用加参数的ticker
```
ping example.com & ticker -b
```
在ping命令完成后,加了-b参数的ticker会弹出一个消息框,默认内容为`ticker complete.`
![](https://github.com/kimmosc2/ticker/raw/master/assets/image/p3.png)  
    
您也可以使用`-m message`参数来自定义消息内容.
![](https://github.com/kimmosc2/ticker/raw/master/assets/image/p4.png)    
  
  
## 如何安装? 
您可以使用以下两种方式来使用Ticker: 
#### 下载并安装二进制包 
我们提供了编译好的二进制包,您下载并解压后将其放入C:\Windows(只要是Path环境变量所
在目录都可以),运行cmd输入ticker就可以使用了.  
下载地址: [百度网盘](https://pan.baidu.com/s/1J9O73CnvfFhDHaF83F1k2A)  
提取码:`k5ad`
#### 源代码编译:
```
cd ticker/app & go build ticker.go
```

## 平台
Ticker的Message box底层调用windows提供的API,所以只能在Windows平台下使用
## 提交Bug  
发现Bug请至issue板块提出,贡献代码请单独拉PR
## 版权和许可信息  
[MIT license](https://github.com/kimmosc2/ticker/blob/master/LICENSE)
## 贡献者
[BuTn](https://github.com/kimmosc2)