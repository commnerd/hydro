package cmd

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5DirEmpty(t *testing.T) {
	makeDir("foo")
	md5, _ := MD5Dir("foo")
	os.Remove("foo")

	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e", md5)
}

func TestMD5Dir(t *testing.T) {
	makeDir("foo")
	defer os.RemoveAll("foo")

	contents := []byte("This is a test")
	os.WriteFile("foo"+PathSeparatorString+"bar.txt", contents, 0644)

	md5, _ := MD5Dir("foo")

	assert.Equal(t, "4bc59800f57b6b0bb652b28c9df27ef9", md5)

}

func TestMD5DirWithEmptyFiles(t *testing.T) {
	makeDir("foo")
	defer os.RemoveAll("foo")
	makeDir("bar")
	defer os.RemoveAll("bar")

	contents := []byte("")
	os.WriteFile("foo"+PathSeparatorString+"bar.txt", contents, 0644)
	os.WriteFile("bar"+PathSeparatorString+"foo.txt", contents, 0644)

	md5foo, _ := MD5Dir("foo")
	md5bar, _ := MD5Dir("bar")

	assert.Equal(t, "1e666ef753adea14e91aabb315ea713f", md5foo)
	assert.Equal(t, "acb5d227c70841c560438b10d2fa5a60", md5bar)
}
