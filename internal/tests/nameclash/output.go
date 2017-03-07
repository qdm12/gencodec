// generated by github.com/fjl/gencodec, do not edit.

package nameclash

import (
	"encoding/json"

	errors0 "github.com/fjl/gencodec/internal/clasherrors"
	json0 "github.com/fjl/gencodec/internal/clashjson"
)

func (y *Y) MarshalJSON() ([]byte, error) {
	type YJSON0 struct {
		Foo    *json0.Foo   `optional:"true"`
		Foo2   *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var enc0 YJSON0
	enc0.Foo = &y.Foo
	enc0.Foo2 = &y.Foo2
	enc0.Bar = &y.Bar
	enc0.Gazonk = &y.Gazonk
	enc0.Over = (*enc)(&y.Over)
	return json.Marshal(&enc0)
}

func (y *Y) UnmarshalJSON(input []byte) error {
	type YJSON0 struct {
		Foo    *json0.Foo   `optional:"true"`
		Foo2   *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var dec YJSON0
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	var x Y
	if dec.Foo != nil {
		x.Foo = *dec.Foo
	}
	if dec.Foo2 != nil {
		x.Foo2 = *dec.Foo2
	}
	if dec.Bar != nil {
		x.Bar = *dec.Bar
	}
	if dec.Gazonk != nil {
		x.Gazonk = *dec.Gazonk
	}
	if dec.Over != nil {
		x.Over = int(*dec.Over)
	}
	*y = x
	return nil
}

func (y *Y) MarshalYAML() (interface{}, error) {
	type YYAML struct {
		Foo    *json0.Foo   `optional:"true"`
		Foo2   *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var enc0 YYAML
	enc0.Foo = &y.Foo
	enc0.Foo2 = &y.Foo2
	enc0.Bar = &y.Bar
	enc0.Gazonk = &y.Gazonk
	enc0.Over = (*enc)(&y.Over)
	return &enc0, nil
}

func (y *Y) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type YYAML struct {
		Foo    *json0.Foo   `optional:"true"`
		Foo2   *json0.Foo   `optional:"true"`
		Bar    *errors0.Foo `optional:"true"`
		Gazonk *YJSON       `optional:"true"`
		Over   *enc         `optional:"true"`
	}
	var dec YYAML
	if err := unmarshal(&dec); err != nil {
		return err
	}
	var x Y
	if dec.Foo != nil {
		x.Foo = *dec.Foo
	}
	if dec.Foo2 != nil {
		x.Foo2 = *dec.Foo2
	}
	if dec.Bar != nil {
		x.Bar = *dec.Bar
	}
	if dec.Gazonk != nil {
		x.Gazonk = *dec.Gazonk
	}
	if dec.Over != nil {
		x.Over = int(*dec.Over)
	}
	*y = x
	return nil
}
