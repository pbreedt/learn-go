package json

import "testing"

func TestMarshal(t *testing.T) {
	m := Metric{Type: "the-type", Payload: Payload{Value: "any value"}}
	s, err := Marshal(m)
	if err != nil {
		t.Fatalf("expected err = nil, got:%v", err)
	}
	if s != `{"type":"the-type","payload":{"value":"any value"}}` {
		t.Fatalf("expected Metric.Payload.Value to be '{\"type\":\"the-type\",\"payload\":{\"value\":\"any value\"}}', got:%s", s)
	}
}

func TestUnmarshal(t *testing.T) {
	json := `[
	  {
		"type": "number-type",
		"payload": {"value":123}
	  },
	  {
		"type": "string-type",
		"payload": {"value":"abc"}
	  }	  
	]`
	mslice, err := Unmarshal(json)
	if err != nil {
		t.Fatalf("expected err = nil, got:%v", err)
	}
	if len(mslice) <= 0 {
		t.Fatalf("expected slice len > 0, got:%d", len(mslice))
	}
	if mslice[0].Type != "number-type" {
		t.Fatalf("expected Type of Metric 0 to be 'number-slice', got:%s", mslice[0].Type)
	}
	if mslice[0].Payload.Value.(float64) != 123 {
		t.Fatalf("expected Payload.Value of Metric 0 to be 123, got:%v (%T)", mslice[0].Payload.Value, mslice[0].Payload.Value)
	}

	if mslice[1].Type != "string-type" {
		t.Fatalf("expected Type of Metric 0 to be 'string-slice', got:%s", mslice[0].Type)
	}
	if mslice[1].Payload.Value.(string) != "abc" {
		t.Fatalf("expected Payload.Value of Metric 0 to be 'abc', got:%v (%T)", mslice[0].Payload.Value, mslice[0].Payload.Value)
	}
}
