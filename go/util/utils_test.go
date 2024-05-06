package util

import "testing"

func Test_RandomAvatarURL(t *testing.T) {
	ans := RandomAvatarURL("genji77@qq.com")
	t.Errorf(ans)
}
