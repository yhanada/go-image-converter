package converter

import "testing"

func TestGetImageTypeJpeg(t *testing.T) {
	result := GetImageType("jpg")
	if result != JPEG {
		t.Error("Failed: not jpeg")
	}
}

func TestGetImageTypeTable(t *testing.T) {
	cases := []struct {
		input    string
		expected ImageType
	}{
		{input: "jpg", expected: JPEG},
		{input: "JPG", expected: JPEG},
		{input: "jpeg", expected: JPEG},
		{input: "JPEG", expected: JPEG},
		{input: "gif", expected: GIF},
		{input: "GIF", expected: GIF},
		{input: "png", expected: PNG},
		{input: "PNG", expected: PNG},
		{input: "hoge", expected: NONE},
		{input: "", expected: NONE},
	}

	for _, c := range cases {
		if GetImageType(c.input) != c.expected {
			t.Error("Invalid ImageType")
		}
	}
}
