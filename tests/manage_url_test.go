package tests

import (
	"testing"

	"github.com/EmeraldLS/url-shortener/controller"
)

func TestRandomID(t *testing.T) {
	got := controller.Random_id()
	want := "NKMrCd"
	if want != got {
		t.Errorf("FAILED: Want %v but got %v", want, got)
	} else {
		t.Logf("PASSED: Want %v and got %v", want, got)
	}
}

func TestShortenURL(t *testing.T) {
	// req := &out.URL_SHORTENER_REQUEST{
	// 	OriginalUrl: "https://www.facebook.com",
	// }
}
