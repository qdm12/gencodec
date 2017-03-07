// generated by github.com/fjl/gencodec, do not edit.

package mapconv

import (
	"encoding/json"
	"errors"
)

func (x *X) MarshalJSON() ([]byte, error) {
	type XJSON struct {
		Map   map[replacedString]replacedInt
		Named namedMap2
	}
	var enc XJSON
	enc.Map = make(map[replacedString]replacedInt, len(x.Map))
	for k, v := range x.Map {
		enc.Map[replacedString(k)] = replacedInt(v)
	}
	enc.Named = make(namedMap2, len(x.Named))
	for k, v := range x.Named {
		enc.Named[replacedString(k)] = replacedInt(v)
	}
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type XJSON struct {
		Map   map[replacedString]replacedInt
		Named namedMap2
	}
	var dec XJSON
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	var x0 X
	if dec.Map == nil {
		return errors.New("missing required field 'map' for X")
	}
	x0.Map = make(map[string]int, len(dec.Map))
	for k, v := range dec.Map {
		x0.Map[string(k)] = int(v)
	}
	if dec.Named == nil {
		return errors.New("missing required field 'named' for X")
	}
	x0.Named = make(namedMap, len(dec.Named))
	for k, v := range dec.Named {
		x0.Named[string(k)] = int(v)
	}
	*x = x0
	return nil
}

func (x *X) MarshalYAML() (interface{}, error) {
	type XYAML struct {
		Map   map[replacedString]replacedInt
		Named namedMap2
	}
	var enc XYAML
	enc.Map = make(map[replacedString]replacedInt, len(x.Map))
	for k, v := range x.Map {
		enc.Map[replacedString(k)] = replacedInt(v)
	}
	enc.Named = make(namedMap2, len(x.Named))
	for k, v := range x.Named {
		enc.Named[replacedString(k)] = replacedInt(v)
	}
	return &enc, nil
}

func (x *X) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type XYAML struct {
		Map   map[replacedString]replacedInt
		Named namedMap2
	}
	var dec XYAML
	if err := unmarshal(&dec); err != nil {
		return err
	}
	var x0 X
	if dec.Map == nil {
		return errors.New("missing required field 'map' for X")
	}
	x0.Map = make(map[string]int, len(dec.Map))
	for k, v := range dec.Map {
		x0.Map[string(k)] = int(v)
	}
	if dec.Named == nil {
		return errors.New("missing required field 'named' for X")
	}
	x0.Named = make(namedMap, len(dec.Named))
	for k, v := range dec.Named {
		x0.Named[string(k)] = int(v)
	}
	*x = x0
	return nil
}
