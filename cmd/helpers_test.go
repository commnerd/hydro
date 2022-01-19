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

func TestMD5File(t *testing.T) {
	contents := []byte("This is a test")
	os.WriteFile("foo.txt", contents, 0644)
	defer os.RemoveAll("foo.txt")

	md5, _ := MD5File("foo.txt")

	assert.Equal(t, "ce114e4501d2f4e2dcea3e17b546f339", md5)
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

func TestIsHydroRepoFalse(t *testing.T) {
	assert.False(t, IsHydroRepo("."))
}

func TestIsHydroRepoTrue(t *testing.T) {
	initDirs("foo")
	defer os.RemoveAll("foo")
	assert.True(t, IsHydroRepo("foo"))
}

func TestInHydroRepoFalse(t *testing.T) {
	makeDir("foo")
	makeDir("foo" + PathSeparatorString + "bar")
	defer os.RemoveAll("foo")
	assert.False(t, InHydroRepo("foo"+PathSeparatorString+"bar"))
}

func TestInHydroRepoTrue(t *testing.T) {
	initDirs("foo")
	defer os.RemoveAll("foo")
	for _, initDir := range initialDirs {
		makeDir("foo" + PathSeparatorString + initDir + PathSeparatorString + "bar")
		assert.True(t, InHydroRepo("foo"+PathSeparatorString+initDir+PathSeparatorString+"bar"))
	}
}
