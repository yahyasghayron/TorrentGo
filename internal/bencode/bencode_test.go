package bencode

import (
  "testing"
  "reflect"
)

func TestEncode(t *testing.T) {
	// Test case 1: Simple map encoding
	data := map[string]interface{}{
		"foo":  1,
		"bar":  "baz",
		"list": []interface{}{1, 2, 3},
	}
	expected := "d3:bar3:baz3:fooi1e4:listli1ei2ei3eee"
	encoded := Encode(data)
	if encoded != expected {
		t.Errorf("Expected %s but got %s", expected, encoded)
	}

	// Test case 2: Empty map encoding
	data = map[string]interface{}{}
	expected = "de"
	encoded = Encode(data)
	if encoded != expected {
		t.Errorf("Expected %s but got %s", expected, encoded)
	}
}

func TestDecode(t *testing.T) {
	// Test case 1: Simple map decoding
	data := []byte("d3:bar3:baz3:fooi1e4:listli1ei2ei3eee")
	expected := map[string]interface{}{
		"foo":  int64(1),
		"bar":  "baz",
		"list": []interface{}{int64(1), int64(2), int64(3)},
	}
	decoded := Decode(data)
	if !reflect.DeepEqual(decoded, expected) {
		t.Errorf("Expected %#v but got %#v", expected, decoded)
	}

	// Test case 2: Empty map decoding
	data = []byte("de")
	expected = map[string]interface{}{}
	decoded = Decode(data)
	if !reflect.DeepEqual(decoded, expected) {
		t.Errorf("Expected %#v but got %#v", expected, decoded)
	}
}
