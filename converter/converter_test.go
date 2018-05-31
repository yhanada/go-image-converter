package converter

import (
	"strings"
	"testing"
)

func TestGetImageTypeJpeg(t *testing.T) {
	result := GetImageType("jpg")
	if result != JPEG {
		t.Error("Failed: not jpeg")
	}
}

func TestGetImageTypeTable(t *testing.T) {
	t.Helper()
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

func TestGetTargetFiles(t *testing.T) {
	t.Helper()
	cases := []struct {
		dir       string
		imageType ImageType
		count     int
	}{
		{dir: "../testdata", imageType: JPEG, count: 2},
		{dir: "../testdata", imageType: GIF, count: 0},
		{dir: "../testdata/foo", imageType: JPEG, count: 1},
		{dir: "../testdata/empty", imageType: JPEG, count: 0},
	}

	for _, c := range cases {
		files := GetTargetFiles(c.imageType, c.dir)
		if len(files) != c.count {
			t.Error("Invalid target files:", c)
		}
	}
}

func TestConvert(t *testing.T) {
	t.Helper()
	src := "../testdata/test.jpg"
	dest, err := Convert(PNG, src)
	if err != nil {
		t.Error("Failed to convert image")
	}

	if strings.Compare(dest, "../testdata/test.png") != 0 {
		t.Error("Invalid destination filepath")
	}
}

func TestConvertErrorCases(t *testing.T) {
	t.Helper()
	cases := []struct {
		targetType ImageType
		dir        string
	}{
		{targetType: PNG, dir: "../testdata/notexist.jpg"},
		{targetType: PNG, dir: "../testdata/test.text"},
		{targetType: NONE, dir: "../testdata/test.jpg"},
	}

	for _, c := range cases {
		_, err := Convert(c.targetType, c.dir)
		if err == nil {
			t.Error("Failed to check error:", c)
		}
	}

}
