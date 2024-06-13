package code_gen

import "testing"

func TestGenStruct(t *testing.T) {
	//此处传入表名以及结构体名
	//GenStruct("ms_project", "Project")
	GenProtoMessage("ms_project", "ProjectMessage")
	//可以在命令行中看到生成的结构体
}
