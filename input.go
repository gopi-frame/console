package console

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gopi-frame/contract/console"
	"github.com/gopi-frame/exception"
	"github.com/spf13/pflag"
)

type Input struct {
	ctx   context.Context
	args  []string
	flags *pflag.FlagSet
}

func NewInput(ctx context.Context, flags *pflag.FlagSet) *Input {
	return &Input{
		ctx:   ctx,
		args:  flags.Args(),
		flags: flags,
	}
}

func (i *Input) Context() context.Context {
	if i.ctx == nil {
		i.ctx = context.Background()
	}
	return i.ctx
}

func (i *Input) Arg(index int) string {
	return i.args[index]
}

func (i *Input) Args() []string {
	return i.args
}

func (i *Input) GetValue(name string, typ string) (console.Value, error) {
	flag := i.flags.Lookup(name)
	if flag == nil {
		return nil, exception.New(fmt.Sprintf("flag %s not found", name))
	}
	if flag.Value.Type() != typ {
		return nil, exception.New(fmt.Sprintf("flag %s is not a %s", name, typ))
	}
	return flag.Value, nil
}

func (i *Input) GetSliceValue(name string, typ string) (console.SliceValue, error) {
	flag := i.flags.Lookup(name)
	if flag == nil {
		return nil, exception.New(fmt.Sprintf("flag %s not found", name))
	}
	if flag.Value.Type() != typ {
		return nil, exception.New(fmt.Sprintf("flag %s is not a %s", name, typ))
	}
	if value, ok := flag.Value.(console.SliceValue); ok {
		return value, nil
	}
	return nil, errors.New(fmt.Sprintf("flag %s is not a slice", name))
}

func (i *Input) GetString(name string) (string, error) {
	value, err := i.GetValue(name, StringType)
	if err != nil {
		return "", err
	}
	if value, ok := value.(console.Valuer[string]); ok {
		return value.Value(), nil
	}
	return value.String(), nil
}

func (i *Input) GetInt(name string) (int, error) {
	value, err := i.GetValue(name, IntType)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[int]); ok {
		return value.Value(), nil
	}
	return strconv.Atoi(value.String())
}

func (i *Input) GetInt8(name string) (int8, error) {
	value, err := i.GetValue(name, Int8Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[int8]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseInt(value.String(), 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

func (i *Input) GetInt16(name string) (int16, error) {
	value, err := i.GetValue(name, Int16Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[int16]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseInt(value.String(), 10, 16)
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func (i *Input) GetInt32(name string) (int32, error) {
	value, err := i.GetValue(name, Int32Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[int32]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseInt(value.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func (i *Input) GetInt64(name string) (int64, error) {
	value, err := i.GetValue(name, Int64Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[int64]); ok {
		return value.Value(), nil
	}
	return strconv.ParseInt(value.String(), 10, 64)
}

func (i *Input) GetUint(name string) (uint, error) {
	value, err := i.GetValue(name, UintType)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[uint]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseUint(value.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func (i *Input) GetUint8(name string) (uint8, error) {
	value, err := i.GetValue(name, Uint8Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[uint8]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseUint(value.String(), 10, 8)
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

func (i *Input) GetUint16(name string) (uint16, error) {
	value, err := i.GetValue(name, Uint16Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[uint16]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseUint(value.String(), 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func (i *Input) GetUint32(name string) (uint32, error) {
	value, err := i.GetValue(name, Uint32Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[uint32]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseUint(value.String(), 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func (i *Input) GetUint64(name string) (uint64, error) {
	value, err := i.GetValue(name, Uint64Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[uint64]); ok {
		return value.Value(), nil
	}
	return strconv.ParseUint(value.String(), 10, 64)
}

func (i *Input) GetFloat32(name string) (float32, error) {
	value, err := i.GetValue(name, Float32Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[float32]); ok {
		return value.Value(), nil
	}
	v, err := strconv.ParseFloat(value.String(), 32)
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func (i *Input) GetFloat64(name string) (float64, error) {
	value, err := i.GetValue(name, Float64Type)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[float64]); ok {
		return value.Value(), nil
	}
	return strconv.ParseFloat(value.String(), 64)
}

func (i *Input) GetBool(name string) (bool, error) {
	value, err := i.GetValue(name, BoolType)
	if err != nil {
		return false, err
	}
	if value, ok := value.(console.Valuer[bool]); ok {
		return value.Value(), nil
	}
	return strconv.ParseBool(value.String())
}

func (i *Input) GetDuration(name string) (time.Duration, error) {
	value, err := i.GetValue(name, DurationType)
	if err != nil {
		return 0, err
	}
	if value, ok := value.(console.Valuer[time.Duration]); ok {
		return value.Value(), nil
	}
	return time.ParseDuration(value.String())
}

func (i *Input) GetTime(name string, layout string) (time.Time, error) {
	value, err := i.GetValue(name, TimeType)
	if err != nil {
		return time.Time{}, err
	}
	if value, ok := value.(console.Valuer[time.Time]); ok {
		return value.Value(), nil
	}
	return time.Parse(layout, value.String())
}

func (i *Input) GetStringSlice(name string) ([]string, error) {
	value, err := i.GetSliceValue(name, StringSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]string]); ok {
		return value.Value(), nil
	}
	return value.GetSlice(), nil
}

func (i *Input) GetIntSlice(name string) ([]int, error) {
	value, err := i.GetSliceValue(name, IntSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]int]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]int, len(values))
	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func (i *Input) GetInt8Slice(name string) ([]int8, error) {
	value, err := i.GetSliceValue(name, Int8SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]int8]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]int8, len(values))
	for i, v := range values {
		n, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return nil, err
		}
		numbers[i] = int8(n)
	}
	return numbers, nil
}

func (i *Input) GetInt16Slice(name string) ([]int16, error) {
	value, err := i.GetSliceValue(name, Int16SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]int16]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]int16, len(values))
	for i, v := range values {
		n, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return nil, err
		}
		numbers[i] = int16(n)
	}
	return numbers, nil
}

func (i *Input) GetInt32Slice(name string) ([]int32, error) {
	value, err := i.GetSliceValue(name, Int32SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]int32]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]int32, len(values))
	for i, v := range values {
		n, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return nil, err
		}
		numbers[i] = int32(n)
	}
	return numbers, nil
}

func (i *Input) GetInt64Slice(name string) ([]int64, error) {
	value, err := i.GetSliceValue(name, Int64SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]int64]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]int64, len(values))
	for i, v := range values {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func (i *Input) GetUintSlice(name string) ([]uint, error) {
	value, err := i.GetSliceValue(name, UintSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]uint]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]uint, len(values))
	for i, v := range values {
		if n, err := strconv.ParseUint(v, 10, 0); err == nil {
			numbers[i] = uint(n)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetUint8Slice(name string) ([]uint8, error) {
	value, err := i.GetSliceValue(name, Uint8SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]uint8]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]uint8, len(values))
	for i, v := range values {
		if n, err := strconv.ParseUint(v, 10, 8); err == nil {
			numbers[i] = uint8(n)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetUint16Slice(name string) ([]uint16, error) {
	value, err := i.GetSliceValue(name, Uint16SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]uint16]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]uint16, len(values))
	for i, v := range values {
		if n, err := strconv.ParseUint(v, 10, 16); err == nil {
			numbers[i] = uint16(n)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetUint32Slice(name string) ([]uint32, error) {
	value, err := i.GetSliceValue(name, Uint32SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]uint32]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]uint32, len(values))
	for i, v := range values {
		if n, err := strconv.ParseUint(v, 10, 32); err == nil {
			numbers[i] = uint32(n)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetUint64Slice(name string) ([]uint64, error) {
	value, err := i.GetSliceValue(name, Uint64SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]uint64]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]uint64, len(values))
	for i, v := range values {
		if n, err := strconv.ParseUint(v, 10, 64); err == nil {
			numbers[i] = n
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetFloat32Slice(name string) ([]float32, error) {
	value, err := i.GetSliceValue(name, Float32SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]float32]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]float32, len(values))
	for i, v := range values {
		if n, err := strconv.ParseFloat(v, 32); err == nil {
			numbers[i] = float32(n)
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetFloat64Slice(name string) ([]float64, error) {
	value, err := i.GetSliceValue(name, Float64SliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]float64]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	numbers := make([]float64, len(values))
	for i, v := range values {
		if n, err := strconv.ParseFloat(v, 64); err == nil {
			numbers[i] = n
		} else {
			return nil, err
		}
	}
	return numbers, nil
}

func (i *Input) GetBoolSlice(name string) ([]bool, error) {
	value, err := i.GetSliceValue(name, BoolSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]bool]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	bools := make([]bool, len(values))
	for i, v := range values {
		if n, err := strconv.ParseBool(v); err == nil {
			bools[i] = n
		} else {
			return nil, err
		}
	}
	return bools, nil
}

func (i *Input) GetDurationSlice(name string) ([]time.Duration, error) {
	value, err := i.GetSliceValue(name, DurationSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]time.Duration]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	durations := make([]time.Duration, len(values))
	for i, v := range values {
		if n, err := time.ParseDuration(v); err == nil {
			durations[i] = n
		} else {
			return nil, err
		}
	}
	return durations, nil
}

func (i *Input) GetTimeSlice(name string, layout string) ([]time.Time, error) {
	value, err := i.GetSliceValue(name, TimeSliceType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[[]time.Time]); ok {
		return value.Value(), nil
	}
	values := value.GetSlice()
	times := make([]time.Time, len(values))
	for i, v := range values {
		if n, err := time.Parse(layout, v); err == nil {
			times[i] = n
		} else {
			return nil, err
		}
	}
	return times, nil
}

func (i *Input) GetStringMap(name string) (map[string]string, error) {
	value, err := i.GetValue(name, StringMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]string]); ok {
		return value.Value(), nil
	}
	values := make(map[string]string)
	pairs := strings.Split(value.String(), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			values[kv[0]] = kv[1]
		} else {
			return nil, fmt.Errorf("invalid string map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetIntMap(name string) (map[string]int, error) {
	value, err := i.GetValue(name, IntMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]int]); ok {
		return value.Value(), nil
	}
	values := make(map[string]int)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseInt(kv[1], 0, 32); err == nil {
				values[kv[0]] = int(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid int map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetInt8Map(name string) (map[string]int8, error) {
	value, err := i.GetValue(name, Int8MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]int8]); ok {
		return value.Value(), nil
	}
	values := make(map[string]int8)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseInt(kv[1], 0, 8); err == nil {
				values[kv[0]] = int8(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid int8 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetInt16Map(name string) (map[string]int16, error) {
	value, err := i.GetValue(name, Int16MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]int16]); ok {
		return value.Value(), nil
	}
	values := make(map[string]int16)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseInt(kv[1], 0, 16); err == nil {
				values[kv[0]] = int16(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid int16 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetInt32Map(name string) (map[string]int32, error) {
	value, err := i.GetValue(name, Int32MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]int32]); ok {
		return value.Value(), nil
	}
	values := make(map[string]int32)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseInt(kv[1], 0, 32); err == nil {
				values[kv[0]] = int32(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid int32 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetInt64Map(name string) (map[string]int64, error) {
	value, err := i.GetValue(name, Int64MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]int64]); ok {
		return value.Value(), nil
	}
	values := make(map[string]int64)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseInt(kv[1], 0, 64); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid int64 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetUintMap(name string) (map[string]uint, error) {
	value, err := i.GetValue(name, UintMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]uint]); ok {
		return value.Value(), nil
	}
	values := make(map[string]uint)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseUint(kv[1], 0, 32); err == nil {
				values[kv[0]] = uint(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid uint map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetUint8Map(name string) (map[string]uint8, error) {
	value, err := i.GetValue(name, Uint8MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]uint8]); ok {
		return value.Value(), nil
	}
	values := make(map[string]uint8)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseUint(kv[1], 0, 8); err == nil {
				values[kv[0]] = uint8(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid uint8 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetUint16Map(name string) (map[string]uint16, error) {
	value, err := i.GetValue(name, Uint16MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]uint16]); ok {
		return value.Value(), nil
	}
	values := make(map[string]uint16)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseUint(kv[1], 0, 16); err == nil {
				values[kv[0]] = uint16(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid uint16 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetUint32Map(name string) (map[string]uint32, error) {
	value, err := i.GetValue(name, Uint32MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]uint32]); ok {
		return value.Value(), nil
	}
	values := make(map[string]uint32)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseUint(kv[1], 0, 32); err == nil {
				values[kv[0]] = uint32(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid uint32 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetUint64Map(name string) (map[string]uint64, error) {
	value, err := i.GetValue(name, Uint64MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]uint64]); ok {
		return value.Value(), nil
	}
	values := make(map[string]uint64)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseUint(kv[1], 0, 64); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid uint64 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetFloat32Map(name string) (map[string]float32, error) {
	value, err := i.GetValue(name, Float32MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]float32]); ok {
		return value.Value(), nil
	}
	values := make(map[string]float32)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseFloat(kv[1], 32); err == nil {
				values[kv[0]] = float32(n)
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid float32 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetFloat64Map(name string) (map[string]float64, error) {
	value, err := i.GetValue(name, Float64MapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]float64]); ok {
		return value.Value(), nil
	}
	values := make(map[string]float64)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseFloat(kv[1], 64); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid float64 map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetBoolMap(name string) (map[string]bool, error) {
	value, err := i.GetValue(name, BoolMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]bool]); ok {
		return value.Value(), nil
	}
	values := make(map[string]bool)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := strconv.ParseBool(kv[1]); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid bool map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetDurationMap(name string) (map[string]time.Duration, error) {
	value, err := i.GetValue(name, DurationMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]time.Duration]); ok {
		return value.Value(), nil
	}
	values := make(map[string]time.Duration)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := time.ParseDuration(kv[1]); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid duration map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) GetTimeMap(name string, layout string) (map[string]time.Time, error) {
	value, err := i.GetValue(name, TimeMapType)
	if err != nil {
		return nil, err
	}
	if value, ok := value.(console.Valuer[map[string]time.Time]); ok {
		return value.Value(), nil
	}
	values := make(map[string]time.Time)
	pairs := strings.Split(strings.Trim(value.String(), "[]"), ",")
	for _, pair := range pairs {
		kv := strings.Split(pair, "=")
		if len(kv) == 2 {
			if n, err := time.Parse(layout, kv[1]); err == nil {
				values[kv[0]] = n
			} else {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid time map value: %s", pair)
		}
	}
	return values, nil
}

func (i *Input) VisitAll(fn func(name string, typ string)) {
	i.flags.VisitAll(func(flag *pflag.Flag) {
		fn(flag.Name, flag.Value.Type())
	})
}

func (i *Input) Unmarshal(v any) error {
	return NewDecoder(i).Decode(v)
}
