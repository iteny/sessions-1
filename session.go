package sessions

import (
	"net/http"
	"strconv"
)

//session对象
type sessionObject struct {
	ID   string
	req  *http.Request
	resp http.ResponseWriter
}

type Value struct {
	Key   string
	Value string
	Error error
}

//设置一个键值，如果键名存在则覆盖
func (obj *sessionObject) Set(key string, value interface{}) error {
	return redisClient.HSet(obj.ID, key, value).Err()
}

//读取参数值
func (obj *sessionObject) Get(key string) *Value {
	var result Value
	result.Key = key
	value, err := redisClient.HGet(obj.ID, key).Result()
	if err != nil {
		result.Error = err
		return &result
	}
	result.Value = value
	return &result
}

//将值转为string类型
func (v *Value) String() (string, error) {
	if v.Error != nil {
		return "", v.Error
	}
	return v.Value, nil
}

//将值转为int类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Int(def ...int) (int, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.Atoi(v.Value)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return value, nil
}

//将参数值转为int32类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Int32(def ...int32) (int32, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseInt(v.Value, 10, 32)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return int32(value), nil
}

//将参数值转为int64类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Int64(def ...int64) (int64, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseInt(v.Value, 10, 64)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return value, nil
}

//将参数值转为uint32类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Uint32(def ...uint32) (uint32, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseUint(v.Value, 10, 32)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return uint32(value), nil
}

//将参数值转为uint64类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Uint64(def ...uint64) (uint64, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseUint(v.Value, 10, 64)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return value, nil
}

//将参数值转为float32类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Float32(def ...float32) (float32, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseFloat(v.Value, 32)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return float32(value), nil
}

//将参数值转为float64类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Float64(def ...float64) (float64, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return 0, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseFloat(v.Value, 64)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return 0, err
		}
	}
	return value, nil
}

//将参数值转为bool类型
//如果传入了def参数值，在转换出错时返回def，并且第二个出参永远为nil
func (v *Value) Bool(def ...bool) (bool, error) {
	defLen := len(def)
	if v.Error != nil {
		if defLen == 0 {
			return false, v.Error
		} else {
			return def[0], nil
		}
	}
	value, err := strconv.ParseBool(v.Value)
	if err != nil {
		if defLen > 0 {
			return def[0], nil
		} else {
			return false, err
		}
	}
	return value, nil
}

//删除一个键值，如果键名不存在则忽略，不会报错
func (this *sessionObject) Delete(key string) error {
	return redisClient.HDel(this.ID, key).Err()
}

//清除所有redis中的session数据，但不删除cookie中的sessionID
func (this *sessionObject) ClearData() error {
	return redisClient.Del(this.ID).Err()
}
