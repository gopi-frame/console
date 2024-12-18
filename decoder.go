package console

import (
	"fmt"

	"github.com/go-viper/mapstructure/v2"
	"github.com/gopi-frame/contract/console"
)

type Decoder struct {
	input              console.Input
	TimeLayout         string
	ArgsTagPrefix      string
	IgnoreUnknownFlags bool
	Config             *mapstructure.DecoderConfig
}

func NewDecoder(input console.Input) *Decoder {
	return &Decoder{
		input:         input,
		TimeLayout:    "2006-01-02 15:04:05",
		ArgsTagPrefix: "args_",
		Config: &mapstructure.DecoderConfig{
			WeaklyTypedInput: true,
			TagName:          "flag",
			DecodeHook: mapstructure.ComposeDecodeHookFunc(
				mapstructure.TextUnmarshallerHookFunc(),
				mapstructure.StringToTimeDurationHookFunc(),
				mapstructure.StringToBasicTypeHookFunc(),
				mapstructure.StringToSliceHookFunc(","),
			),
		},
	}
}

func (d *Decoder) Decode(v any) error {
	input := d.input
	var values = make(map[string]any)
	input.VisitAll(func(name, typ string) {
		switch typ {
		case StringType:
			values[name], _ = input.GetString(name)
		case StringSliceType:
			values[name], _ = input.GetStringSlice(name)
		case StringMapType:
			values[name], _ = input.GetStringMap(name)
		case BoolType:
			values[name], _ = input.GetBool(name)
		case BoolSliceType:
			values[name], _ = input.GetBoolSlice(name)
		case BoolMapType:
			values[name], _ = input.GetBoolMap(name)
		case DurationType:
			values[name], _ = input.GetDuration(name)
		case DurationSliceType:
			values[name], _ = input.GetDurationSlice(name)
		case DurationMapType:
			values[name], _ = input.GetDurationMap(name)
		case TimeType:
			values[name], _ = input.GetTime(name, d.TimeLayout)
		case TimeSliceType:
			values[name], _ = input.GetTimeSlice(name, d.TimeLayout)
		case TimeMapType:
			values[name], _ = input.GetTimeMap(name, d.TimeLayout)
		case Float32Type:
			values[name], _ = input.GetFloat32(name)
		case Float32SliceType:
			values[name], _ = input.GetFloat32Slice(name)
		case Float32MapType:
			values[name], _ = input.GetFloat32Map(name)
		case Float64Type:
			values[name], _ = input.GetFloat64(name)
		case Float64SliceType:
			values[name], _ = input.GetFloat64Slice(name)
		case Float64MapType:
			values[name], _ = input.GetFloat64Map(name)
		case IntType:
			values[name], _ = input.GetInt(name)
		case IntSliceType:
			values[name], _ = input.GetIntSlice(name)
		case IntMapType:
			values[name], _ = input.GetIntMap(name)
		case Int8Type:
			values[name], _ = input.GetInt8(name)
		case Int8SliceType:
			values[name], _ = input.GetInt8Slice(name)
		case Int8MapType:
			values[name], _ = input.GetInt8Map(name)
		case Int16Type:
			values[name], _ = input.GetInt16(name)
		case Int16SliceType:
			values[name], _ = input.GetInt16Slice(name)
		case Int16MapType:
			values[name], _ = input.GetInt16Map(name)
		case Int32Type:
			values[name], _ = input.GetInt32(name)
		case Int32SliceType:
			values[name], _ = input.GetInt32Slice(name)
		case Int32MapType:
			values[name], _ = input.GetInt32Map(name)
		case Int64Type:
			values[name], _ = input.GetInt64(name)
		case Int64SliceType:
			values[name], _ = input.GetInt64Slice(name)
		case Int64MapType:
			values[name], _ = input.GetInt64Map(name)
		case UintType:
			values[name], _ = input.GetUint(name)
		case UintSliceType:
			values[name], _ = input.GetUintSlice(name)
		case UintMapType:
			values[name], _ = input.GetUintMap(name)
		case Uint8Type:
			values[name], _ = input.GetUint8(name)
		case Uint8SliceType:
			values[name], _ = input.GetUint8Slice(name)
		case Uint8MapType:
			values[name], _ = input.GetUint8Map(name)
		case Uint16Type:
			values[name], _ = input.GetUint16(name)
		case Uint16SliceType:
			values[name], _ = input.GetUint16Slice(name)
		case Uint16MapType:
			values[name], _ = input.GetUint16Map(name)
		case Uint32Type:
			values[name], _ = input.GetUint32(name)
		case Uint32SliceType:
			values[name], _ = input.GetUint32Slice(name)
		case Uint32MapType:
			values[name], _ = input.GetUint32Map(name)
		case Uint64Type:
			values[name], _ = input.GetUint64(name)
		case Uint64SliceType:
			values[name], _ = input.GetUint64Slice(name)
		case Uint64MapType:
			values[name], _ = input.GetUint64Map(name)
		default:
			values[name], _ = input.GetString(name)
		}
	})
	for index, value := range input.Args() {
		values[fmt.Sprintf("%s%d", d.ArgsTagPrefix, index)] = value
	}
	d.Config.Result = v
	decoder, err := mapstructure.NewDecoder(d.Config)
	if err != nil {
		return err
	}
	return decoder.Decode(values)
}
