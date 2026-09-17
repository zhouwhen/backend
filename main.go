package main

import(
	"fmt"
	"golang.org/x/crypto/bcrypt"     //加密   所有有关这个加密的我都是问豆包的 不好意思我不会呀 现在可能懂了一点点
)
	var choice int   //操作


func main(){
	//操作事件
    fmt.Println("--------通行证--------")
	fmt.Println("1，注册")
    fmt.Println("2，登录")
    fmt.Println("0‘退出")
	userInfo :=make(map[string][]byte)   //map存用户信息；key存用户名；value存加密密码
	for{
	//读取
        fmt.Print("请选择操作：")
        fmt.Scan(&choice)
   
    	//不同操作情况
	    switch choice{
	    	case 1:
	    	    register(userInfo)
            //body
	    	case 2:
	    	    login(userInfo)
	        //body
	    	case 0:
		        fmt.Println("再见")
	        //body
	    	    return
	    	default:
		        fmt.Println("不符合要求，请输入数字0/1/2")
	    }   
	}
}
//注册
func register(userInfo map[string][]byte) {
	var name string  //用户名
	var password string  //密码
	fmt.Print("请输入用户名：")
	fmt.Scan(&name)

	_, exist :=userInfo[name]    //看是否存在

	if exist{
		fmt.Println("该账号已存在")
		return
	}

	fmt.Print("请输入密码：")
	fmt.Scan(&password)

	hashpwd, err :=bcrypt.GenerateFromPassword([]byte(password),10)
	if err != nil{
		fmt.Println("密码加密失败")   //加密错误
		return
	}

	userInfo[name]=hashpwd
	fmt.Println("注册成功")
}
//登录
func login(userInfo map[string][]byte) {
	var name string  //用户名
	var password string  //密码
	fmt.Print("请输入用户名：")
	fmt.Scan(&name)

	hashPwd,exist:=userInfo[name]   //看是否存在

	if !exist{
		fmt.Println("账号不存在")
		return
	}
	
	//if exist{
		fmt.Print("请输入密码：")
		fmt.Scan(&password)
		err:=bcrypt.CompareHashAndPassword(hashPwd,[]byte(password)) //验证密码是否正确
		
  
    if err!=nil{
		fmt.Println("密码错误")
	}else{
		fmt.Println("登录成功")
	}
	//}
}