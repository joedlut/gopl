package word

import (
	"math/rand"
	"testing"
	"time"
)
//随机测试 随机生成回文字符串
func randomPalindrome(rng  *rand.Rand)string{
	n := rng.Intn(25)
	runes := make([]rune,n)
	for i:=0;i<(n+1)/2;i++{
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T){
	rng := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	t.Logf("random seed %d",rng)
	for i:=0;i<1000;i++{
		p := randomPalindrome(rng)
		//fmt.Println(p)
		if !IsPalindrome(p){
			t.Errorf("IsPalindrome(%q)=false",p)
		}
	}
}