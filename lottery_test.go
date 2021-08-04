package sportlottery

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T)  {
	client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	client.SetMods("pro")
	err,fff := client.Dlt("21069")
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//err,fff:= client.Pls("21090ww")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//err,fff:= client.Plw(map[string]string{"drawNo":"21047"})
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//err,fff:= client.Qxc("21047")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//err,fff:=client.ZuCaiIssue("94")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//	err,fff:=client.ZuCai("21001","90")
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	fmt.Println(fff)

	//client := NewLottery("7f56bbec939e","1f41f1a672cc","430003")
	//err,fff:=client.JcMatchList("2")
	//if err != nil {
	//	fmt.Println(err)
	//}

	fmt.Println(fff)
}
