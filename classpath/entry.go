package classpath
import "os"
import "strings"
const pathListSeparator = string(os.PathListSeparator) // 存放路径分隔符
type Entry interface {
	readClass(className string) ([]byte, Entry, error) // 寻找、加载class文件
	String() string // toString 返回变量的字符串表示
}
func newEntry(path string) Entry{
	if string.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if string.HasSuffix(path, "*"){
		return newWildcardEntry(path)
	}
	if string.HasSuffix(path, ".jar") || string.HasSuffix(path, ".JAR") || 
		string.HasSuffix(path, ".zip") || string.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}