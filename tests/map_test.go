package tests

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"runtime"
	"test1/Utils"
	"testing"
)

var intMap map[int]int
var cnt = 8192

func TestClearMap(t *testing.T) {
	printMemStats()
	initMap()
	runtime.GC()
	printMemStats()
	log.Println(len(intMap))
	for i := 0; i < cnt; i++ {
		delete(intMap, i)
	}
	log.Println(len(intMap))
	runtime.GC()
	printMemStats()
	intMap = nil
	runtime.GC()
	printMemStats()
}
func initMap() {
	intMap = make(map[int]int, cnt)
	for i := 0; i < cnt; i++ {
		intMap[i] = i
	}
}
func printMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("Alloc = %v TotalAlloc = %v Sys = %v NumGC = %v\n", m.Alloc/1024, m.TotalAlloc/1024, m.Sys/1024, m.NumGC)
}

type student struct {
	Name string
	Age  int
}

func TestRange(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		// 可以看到地址一直没有变化,也就是 stu 这个变量只初始化一次,后续都是在前面的同一个地址上赋值,地址一直都是一样.
		// 那么最后 map 里的值也都是一样的,因为循环结束以后 stu 这个地址最后的值是最后一次循环赋的值.
		//fmt.Printf("%p \n", &stu)
		// 为什么这种操作可以,这是新建了一个临时局部变量,会给这个临时变量分配新的地址,值是 stu 的值,那么 map 最后指向的是临时变量的地址对应的值
		// tmp := stu
		m[stu.Name] = &stu
	}
	//map 中存放的是地址, 所以这种更改是有效的
	for _, val := range m {
		val.Age = 999
	}
	fmt.Println("1......")
	Utils.ReceiveStruct(m)
	fmt.Println("1------")
	m2 := make(map[string]student)
	for _, stu := range stus {
		m2[stu.Name] = stu
	}
	//map 中存放的是值, 所以这种更改是无效的
	for _, val := range m2 {
		val.Age = 999
	}
	fmt.Printf("map2 : %v \n", m2)

}

func TestMapAndSlice(t *testing.T) {
	// 测试下 slice, 初始化容量,防止扩容
	arr1 := []int{1, 2, 3, 4, 5}
	slice1 := make([]int, 0, 0)
	map1 := make(map[int]int)

	for _, val := range arr1 {
		slice1 = append(slice1, val)
		map1[val] = val
	}

	fmt.Printf("slice1 : %v \n", slice1)
	fmt.Printf("map1 : %v \n", map1)

	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	slice2 := make([]student, 0, 0)
	map2 := make(map[string]student)

	for _, val := range stus {
		slice2 = append(slice2, val)
		map2[val.Name] = val
	}

	fmt.Printf("slice2 : %v \n", slice2)
	fmt.Printf("map2 : %v \n", map2)
}

var minVal = ^(int(^uint(0) >> 1))
var maxVal = int(^uint(0) >> 1)

func TestInitCapMap(t *testing.T) {
	fmt.Println(maxVal)
	fmt.Println(minVal)
	mm := make(map[int]int, 10)
	for key, val := range mm {
		fmt.Println(key)
		fmt.Println(val)
	}
}

func TestMapAppend(t *testing.T) {
	m1 := make(map[string][]string)
	slice := []string{"1", "2", "3", "4"}
	for _, s := range slice {
		m1["1"] = append(m1["1"], s)
	}
	fmt.Println(m1)
}

type Group struct {
	ParamID    int         `json:"paramId"`
	Operator   int         `json:"operator"`
	ParamValue interface{} `json:"paramValue"`
	Groups     []Group     `json:"groups"`
	Relation   string      `json:"relation"`
}

type RuleObjectDetailRsp struct {
	RuleObjectKey     string `json:"ruleObjectKey"`
	BusinessKey       string `json:"businessKey"`
	RuleObjectName    string `json:"ruleObjectName"`
	LastUseTime       string `json:"lastUseTime"`
	Creator           string `json:"creator"`
	Updater           string `json:"updater"`
	CreateTime        string `json:"createTime"`
	UpdateTime        string `json:"updateTime"`
	IsHaveDraftData   bool   `json:"isHaveDraftData"`
	RuleObjectContent struct {
		Use struct {
			Rules []struct {
				ConditionRoot struct {
					ParamID    int         `json:"paramId"`
					Operator   int         `json:"operator"`
					ParamValue interface{} `json:"paramValue"`
					Relation   string      `json:"relation"`
					Groups     []struct {
						ParamID    int         `json:"paramId"`
						Operator   int         `json:"operator"`
						ParamValue interface{} `json:"paramValue"`
						Groups     []struct {
							ParamID    int         `json:"paramId"`
							Operator   int         `json:"operator"`
							ParamValue interface{} `json:"paramValue"`
							Groups     []string    `json:"groups"`
							Relation   string      `json:"relation"`
						} `json:"groups"`
						Relation string `json:"relation"`
					} `json:"groups"`
				} `json:"conditionRoot"`
				Result []struct {
					ParamID    int         `json:"paramId"`
					ParamKey   string      `json:"paramKey"`
					ParamName  string      `json:"paramName"`
					ParamValue interface{} `json:"paramValue"`
					DataType   string      `json:"dataType"`
				} `json:"result"`
				ElseResults []struct {
					ParamID    int         `json:"paramId"`
					ParamKey   string      `json:"paramKey"`
					ParamName  string      `json:"paramName"`
					ParamValue interface{} `json:"paramValue"`
				} `json:"elseResults"`
				Name string `json:"name"`
			} `json:"rules"`
		} `json:"use"`
		Draft struct {
			Rules []struct {
				ConditionRoot struct {
					ParamID    string `json:"paramId"`
					Operator   int    `json:"operator"`
					ParamValue string `json:"paramValue"`
					Relation   string `json:"relation"`
					Groups     []struct {
						ParamID    int         `json:"paramId"`
						Operator   int         `json:"operator"`
						ParamValue interface{} `json:"paramValue"`
						Groups     []struct {
							ParamID    int         `json:"paramId"`
							Operator   int         `json:"operator"`
							ParamValue interface{} `json:"paramValue"`
							Groups     []string    `json:"groups"`
							Relation   string      `json:"relation"`
						} `json:"groups"`
						Relation string `json:"relation"`
					} `json:"groups"`
				} `json:"conditionRoot"`
				Result []struct {
					ParamID    int         `json:"paramId"`
					ParamKey   string      `json:"paramKey"`
					ParamName  string      `json:"paramName"`
					ParamValue interface{} `json:"paramValue"`
				} `json:"result"`
				ElseResults []struct {
					ParamID    int         `json:"paramId"`
					ParamKey   string      `json:"paramKey"`
					ParamName  string      `json:"paramName"`
					ParamValue interface{} `json:"paramValue"`
				} `json:"elseResults"`
				Name string `json:"name"`
			} `json:"rules"`
		} `json:"draft"`
	} `json:"ruleObjectContent"`
}

func Test12111(t *testing.T) {
	var tmp RuleObjectDetailRsp
	str := "{\"version\":\"1.0\",\"timestamp\":1704269762,\"eventId\":\"197013e1-fafc-4f7d-9de8-7b67c7075b5e\",\"componentName\":\"QC_MIXSERVICE\",\"returnValue\":0,\"returnCode\":0,\"returnMessage\":\"ok\",\"interface\":\"mixservice.rule.getRuleObjectDetailByTypeForApi\",\"data\":{\"ruleObjectKey\":\"bms_industry_tree20240103168175\",\"businessKey\":\"bms_industry_tree\",\"ruleObjectName\":\"医疗行业与区域销售对应财管特殊审批人\",\"creator\":\"doreenlong\",\"updater\":\"doreenlong\",\"createTime\":\"2024-01-03 16:10:24\",\"updateTime\":\"2024-01-03 16:16:00\",\"isHaveDraftData\":false,\"ruleObjectContent\":{\"draft\":{},\"use\":{\"rules\":[{\"conditionRoot\":{\"id\":19466,\"ruleId\":2152,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19467,\"ruleId\":2152,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":53582,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19467}],\"nodeId\":19466},\"result\":[{\"paramId\":517,\"paramValue\":[\"sherrihuang\",\"fredxzhang\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19468,\"ruleId\":2153,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19469,\"ruleId\":2153,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":62586,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19469}],\"nodeId\":19468},\"result\":[{\"paramId\":517,\"paramValue\":[\"sherrihuang\",\"fredxzhang\",\"victormeng\",\"amayayang\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19470,\"ruleId\":2154,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19471,\"ruleId\":2154,\"isRoot\":2,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19472,\"ruleId\":2154,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":62587,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19472},{\"id\":19473,\"ruleId\":2154,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":48316,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19473},{\"id\":19474,\"ruleId\":2154,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":60630,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19474}],\"nodeId\":19471}],\"nodeId\":19470},\"result\":[{\"paramId\":517,\"paramValue\":[\"vincentngwu\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19475,\"ruleId\":2155,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19476,\"ruleId\":2155,\"isRoot\":2,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19477,\"ruleId\":2155,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":37319,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19477},{\"id\":19478,\"ruleId\":2155,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":60628,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19478},{\"id\":19479,\"ruleId\":2155,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":62584,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19479},{\"id\":19480,\"ruleId\":2155,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":29682,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19480}],\"nodeId\":19476}],\"nodeId\":19475},\"result\":[{\"paramId\":517,\"paramValue\":[\"victormeng\",\"amayayang\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19481,\"ruleId\":2156,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19482,\"ruleId\":2156,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":1,\"paramValue\":62588,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19482}],\"nodeId\":19481},\"result\":[{\"paramId\":517,\"paramValue\":[\"vincentngwu\",\"victormeng\",\"amayayang\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19483,\"ruleId\":2157,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19484,\"ruleId\":2157,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":13,\"paramValue\":28290,\"createTime\":\"2024-01-03 16:16:00\",\"updateTime\":\"2024-01-03 16:16:00\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19484}],\"nodeId\":19483},\"result\":[{\"paramId\":517,\"paramValue\":[\"olleykong\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19485,\"ruleId\":2158,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19486,\"ruleId\":2158,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":13,\"paramValue\":28291,\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19486}],\"nodeId\":19485},\"result\":[{\"paramId\":517,\"paramValue\":[\"eevayuan\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19487,\"ruleId\":2159,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19488,\"ruleId\":2159,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":13,\"paramValue\":45160,\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19488}],\"nodeId\":19487},\"result\":[{\"paramId\":517,\"paramValue\":[\"sherrihuang\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19489,\"ruleId\":2160,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19490,\"ruleId\":2160,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":13,\"paramValue\":25667,\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19490}],\"nodeId\":19489},\"result\":[{\"paramId\":517,\"paramValue\":[\"jackyzhuang\",\"shirleyxliu\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"},{\"conditionRoot\":{\"id\":19491,\"ruleId\":2161,\"isRoot\":1,\"relation\":\"or\",\"paramId\":0,\"operator\":0,\"paramValue\":\"\",\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":null,\"paramName\":null,\"useType\":null,\"dataType\":null,\"source\":null,\"inputType\":null,\"isValid\":null,\"groups\":[{\"id\":19492,\"ruleId\":2161,\"isRoot\":2,\"relation\":\"\",\"paramId\":510,\"operator\":13,\"paramValue\":28481,\"createTime\":\"2024-01-03 16:16:01\",\"updateTime\":\"2024-01-03 16:16:01\",\"paramKey\":\"unitId\",\"paramName\":\"组织架构\",\"useType\":\"rule\",\"dataType\":\"int\",\"source\":\"common\",\"inputType\":\"unit\",\"isValid\":1,\"groups\":[],\"nodeId\":19492}],\"nodeId\":19491},\"result\":[{\"paramId\":517,\"paramValue\":[\"jackyzhuang\",\"shirleyxliu\",\"daisytu\"],\"paramKey\":\"rtxs\",\"paramName\":\"个人\",\"dataType\":\"string_array\"}],\"name\":\"\"}],\"labels\":[],\"elseResults\":[]}}}}"
	err := jsoniter.UnmarshalFromString(str, &tmp)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v", tmp)
}
