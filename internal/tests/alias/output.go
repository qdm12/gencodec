// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package alias

import (
	"encoding/json"
)

// MarshalJSON marshals as JSON.
func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		A Aliased
		B AliasedTwice
	}
	var enc X
	enc.A = x.A
	enc.B = x.B
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		A *Aliased
		B *AliasedTwice
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.A != nil {
		x.A = *dec.A
	}
	if dec.B != nil {
		x.B = *dec.B
	}
	return nil
}
