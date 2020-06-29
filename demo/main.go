package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

/*
func main() {
	m:=make(map[string]interface{})
	m["haha"]=123
	m["test"]="wod"
	m["go"]=true
	fmt.Println(m)
	v,ok:=m["wc"]
	if ok{
		fmt.Printf("v is %#v\n",v)
	}else{
		fmt.Println("map 不存在这样的key[\"wc\"]")
	}
	rand.Seed(time.Now().UnixNano())
	code:=make([]string,6)
	for i:= 0;i<6;i++{
	   code[i]=strconv.Itoa(rand.Intn(10))
	}
	str:=strings.Join(code,"")
	fmt.Printf("str=========>%#v\n",str)

}
*/

/*
func aa() {
	var debug bool
	result := map[string]string{"uuid": "121566asds"}
	if debug {
		result["code"] = "sssdsad411"
	}
	fmt.Printf("result=====>%#v\n", result)
}
func bb(){
	param:=url.Values{}
	param.Set("phone","13456879565,13588887777,13600087891")
	param.Set("var","varible")
	param.Set("signature","sig")
	strs:=make([]string,0)
	for k:=range param{
		fmt.Printf("k=========>%#v\n",k)
		strs=append(strs,k)
	}
	sort.Strings(strs)
	fmt.Printf("strs=============>%#v\n",strs)
}
**/

type SMSConfig struct {
	Debug      bool
	CodeLength int           `yaml:"codeLength"`
	CodeExpire time.Duration `yaml:"codeExpire"`
	Cooldown   time.Duration
	Limit      string
	Server     struct {
		Cloud     string
		User      string
		Key       string
		Templates []struct {
			Language string
			Template string
		}
	}
}

func init() {
	workDir, _ := os.Getwd()
	fmt.Println(workDir)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s \n", err)
		return
	}

}

func main() {

	var tmp SMSConfig
	if err := viper.Unmarshal(&tmp); err != nil {
		fmt.Println(err)
		return
	}
	for i, v := range tmp.Server.Templates {
		fmt.Printf("template[%d]=====>%#v\n", i, v)
	}
}
