一个从数据库结构到model和proto message的生成器

依据sql指令describe生成的结果
![img.png](img.png)

使用方法:
调用GenStruct(table_name, struct_name),传入相应的参数。  
code_gen_test.go文件中有例子  
  
完善想法：使得生成的结果能写入.go文件中,或者传入.go文件名将结果写入进去(待实现)  
![img_1.png](img_1.png)