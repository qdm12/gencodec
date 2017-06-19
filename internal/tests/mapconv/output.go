// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package mapconv

import (
	"encoding/json"
)

var _ = (*Xo)(nil)

func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
		Func        map[replacedString]replacedInt
	}
	var enc X
	if x.Map != nil {
		enc.Map = make(map[replacedString]replacedInt, len(x.Map))
		for k, v := range x.Map {
			enc.Map[replacedString(k)] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedMap2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[replacedString(k)] = replacedInt(v)
		}
	}
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	tmp := x.Func()
	if tmp != nil {
		enc.Func = make(map[replacedString]replacedInt, len(tmp))
		for k, v := range tmp {
			enc.Func[replacedString(k)] = replacedInt(v)
		}
	}
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Map != nil {
		x.Map = make(map[string]int, len(dec.Map))
		for k, v := range dec.Map {
			x.Map[string(k)] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedMap, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[string(k)] = int(v)
		}
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}

func (x X) MarshalYAML() (interface{}, error) {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
		Func        map[replacedString]replacedInt
	}
	var enc X
	if x.Map != nil {
		enc.Map = make(map[replacedString]replacedInt, len(x.Map))
		for k, v := range x.Map {
			enc.Map[replacedString(k)] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedMap2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[replacedString(k)] = replacedInt(v)
		}
	}
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	tmp := x.Func()
	if tmp != nil {
		enc.Func = make(map[replacedString]replacedInt, len(tmp))
		for k, v := range tmp {
			enc.Func[replacedString(k)] = replacedInt(v)
		}
	}
	return &enc, nil
}

func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Map != nil {
		x.Map = make(map[string]int, len(dec.Map))
		for k, v := range dec.Map {
			x.Map[string(k)] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedMap, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[string(k)] = int(v)
		}
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}

func (x X) MarshalTOML() (interface{}, error) {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
		Func        map[replacedString]replacedInt
	}
	var enc X
	if x.Map != nil {
		enc.Map = make(map[replacedString]replacedInt, len(x.Map))
		for k, v := range x.Map {
			enc.Map[replacedString(k)] = replacedInt(v)
		}
	}
	if x.Named != nil {
		enc.Named = make(namedMap2, len(x.Named))
		for k, v := range x.Named {
			enc.Named[replacedString(k)] = replacedInt(v)
		}
	}
	enc.NoConv = x.NoConv
	enc.NoConvNamed = x.NoConvNamed
	tmp := x.Func()
	if tmp != nil {
		enc.Func = make(map[replacedString]replacedInt, len(tmp))
		for k, v := range tmp {
			enc.Func[replacedString(k)] = replacedInt(v)
		}
	}
	return &enc, nil
}

func (x *X) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type X struct {
		Map         map[replacedString]replacedInt
		Named       namedMap2
		NoConv      map[string]int
		NoConvNamed namedMap
	}
	var dec X
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Map != nil {
		x.Map = make(map[string]int, len(dec.Map))
		for k, v := range dec.Map {
			x.Map[string(k)] = int(v)
		}
	}
	if dec.Named != nil {
		x.Named = make(namedMap, len(dec.Named))
		for k, v := range dec.Named {
			x.Named[string(k)] = int(v)
		}
	}
	if dec.NoConv != nil {
		x.NoConv = dec.NoConv
	}
	if dec.NoConvNamed != nil {
		x.NoConvNamed = dec.NoConvNamed
	}
	return nil
}
