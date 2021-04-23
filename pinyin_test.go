package pinyin

import (
	"testing"
)

func TestConvert(t *testing.T) {
	str,str2, err := New("hi,我是中国人").Split("").Mode(InitialsInCapitals).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str,str2)
	}

	str,str2, err = New("hi, 我是中国人").Split(" ").Mode(WithoutTone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str,str2)
	}

	str,str2, err = New("我是hahah中国人").Split("-").Mode(Tone).Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str,str2)
	}

	str,str2, err = New("我是h中国人a").Convert()
	if err != nil {
		t.Error(err)
	} else {
		t.Log(str,str2)
	}
}
