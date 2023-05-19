package storage2

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T){
	//保存留待恢复的NotifyUser
	saved := notifyUser
	defer func(){
		notifyUser = saved
	}()

	var notifiedUser,notifiedMsg string
	//替换成 模拟发送邮件的函数,只是完成变量的赋值而已
	notifyUser = func(user,msg string) {
		notifiedUser,notifiedMsg = user,msg
	}
	//模拟980MB 使用情况
	const user = "abc@example.com"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg== ""{
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user{
		t.Errorf("wrong user (%s) notified ,want %s",notifiedUser,user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg,wantSubstring){
		t.Errorf("unexpected notification message <<%s>>" + "want substring %q",notifiedMsg,wantSubstring)
	}


 }
