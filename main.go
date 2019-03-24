/**
 * 功能说明：“模块”初始化文件
 *
 * @desc 模块名称
 * ---------------------------------------------------------------------
 * @author		wuxiaoyong <vip120@126.com>
 * @date		2019-03-24
 * @copyright	公司名称
 * ---------------------------------------------------------------------
 */

package main

import (
	"github.com/Unknwon/goconfig"
	"github.com/lunny/log"
	"fmt"
)

func main(){

	//加截配置文件到内存中
	cfg,err:=goconfig.LoadConfigFile("conf/config.ini")
	if err!=nil{
		log.Fatalf("无法加载配置文件,错误原因：%s",err)
	}

	//对文件的基本操作
	/**a
	section为空中，默认区goconfig.DEFAULT_SECTION,只能读写没有分section块的配置项
	*/
	value,err:=cfg.GetValue(goconfig.DEFAULT_SECTION,"search")
	fmt.Println(value)
	//读取demo块中的key1的值
	key1,err:=cfg.GetValue("Demo","key1")
	fmt.Println(key1)

	//读取parent.child块中的key1的值
	age,err:=cfg.GetValue("parent.child","age")
	fmt.Println(age)

	//写入值

	istrue:=cfg.SetValue("Demo","myname","wuxiaoyong")
	if istrue{
		fmt.Printf("写入%s的值是%s\n","myname","wuxiaoyong")
	}
	myname,err:=cfg.GetValue("Demo","myname")
	if err!=nil{
		fmt.Printf("读取%s的值失败,失败原因%s","myname",err)
	}
	fmt.Printf("读取%s的值是%s\n","myname",myname)

	//读取一个块内的所有配置项,返回一个map
	Demo,err:=cfg.GetSection("Demo")
	//读取map中的一个值
	fmt.Println(Demo["myname"])


	// 获取冒号为分隔符的键值
	value, err = cfg.GetValue("super", "key_super2")
	if err != nil {
		log.Fatalf("无法获取键值（%s）：%s", "key_super2", err)
	}
	log.Printf("%s: %s","key_super2", value)


	//读取分区的注释信息
	comment:=cfg.GetSectionComments("super")
	log.Printf("super的注释信息：%s\n",comment)
	//读取某个键的注释信息
	keycomment:=cfg.GetKeyComments("super","name")
	log.Printf("name的注释信息：%s\n",keycomment)


	//设置分区的注释信息 返回false 时，说明注释被重写（注意）
	v:=cfg.SetSectionComments("super","; 这是super新的注释信息")
	log.Printf("super的新的注释信息：%v",v)

	//设置某个键的注释信息
	keyv:=cfg.SetKeyComments("super","name","; 我的新注释")
	log.Printf("name的新的注释信息：%v",keyv)


	//删除某个键值
	ok:=cfg.DeleteKey("Demo","key1")
	log.Printf("是否删除成功:%v",ok)

	//保存键值
	err=goconfig.SaveConfigFile(cfg,"conf/config.ini")
	if err!=nil{
		log.Printf("写入失败，%s",err)
	}




}

