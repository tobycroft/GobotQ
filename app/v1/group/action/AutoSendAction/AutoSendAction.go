package AutoSendAction

import (
	"errors"
)

func App_autosend_verify(sep, count int) error {
	if sep < 0 {
		return errors.New("时间间隔需要大于0分钟")
	}
	if count < 0 {
		return errors.New("剩余重复次数应该大于0")
	}
}
