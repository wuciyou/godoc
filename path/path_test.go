package path

import (
	"path"
	"testing"
)

// 删除多余的字符，返回最终解析后的路径
func TestClean(t *testing.T) {
	t.Log(path.Clean("/src///../testgo////////thinkgo/main.go")) // /testgo/thinkgo/main.go
}

// 获取uri 中的路径和文件名
func TestSplit(t *testing.T) {
	pathstr, filename := path.Split(path.Clean("/src///../testgo////////thinkgo/main.go"))
	t.Logf("pathstr:%s \n filename:%s", pathstr, filename)

	/**
	* return :
	* pathstr:/testgo/thinkgo/
	* filename:main.go
	 */
}
