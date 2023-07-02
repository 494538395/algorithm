package speed

//
//import (
//	"sync/atomic"
//	"time"
//)
//
//var (
//	SpeedWriter = newSpeedWriter()
//)
//
//var (
//	PreV1   int64
//	PreV2   int64
//	Current int64
//)
//
//type speedWriter struct {
//}
//
//func newSpeedWriter() *speedWriter {
//	return &speedWriter{}
//}
//
//func (s *speedWriter) Write(p []byte) (n int, err error) {
//	atomic.AddInt64(&Current, int64(len(p)))
//	return len(p), nil
//}
//
//func (s speedWriter) Save() {
//	PreV1 = PreV2
//	PreV2 = Current
//}
//
//func (s speedWriter) GetSpeed() int64 {
//	return PreV2 - PreV1
//}
//
//func main() {
//	go func() {
//		router.Run()
//	}()
//
//	controllers.Start = time.Now()
//
//	// 定时更新 v1 ，v2
//	go func() {
//		for {
//			time.Sleep(time.Second)
//			docker.SpeedWriter.Save()
//		}
//	}()
//	cmd.Execute()
//}
