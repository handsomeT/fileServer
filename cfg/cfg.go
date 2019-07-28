package cfg

const File = "cfg.json"

// 最
const MaxUploadSize = 200 * 1024 * 1024

type Cfg struct {
	Dir  string
	Port int
}
