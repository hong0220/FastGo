package tokenbucket

//
//import (
//	"fmt"
//	"github.com/hong0220/fastgo/pkg/middleware/tokenbucket"
//	"testing"
//	"time"
//)
//
//func TestTokenBucket(t *testing.T) {
//	bucket := tokenbucket.NewBucket(5, time.Second)
//	for i := 0; i < 1000; i++ {
//		time.Sleep(time.Millisecond * 100)
//		go DoFunc(bucket, i)
//	}
//	for {
//		fmt.Println("....")
//	}
//}
//
//func DoFunc(bucket *tokenbucket.Bucket, index int) {
//	if bucket.GetToken() {
//		fmt.Printf("getToken success : %d\n", index)
//	} else { // 丢弃
//		fmt.Printf("getToken fail : %d\n", index)
//	}
//}
