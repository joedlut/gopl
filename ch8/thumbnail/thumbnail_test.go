package thumbnail_test

import (
	"log"
	"os"
	"sync"

	"ch8/thumbnail"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 错误的写法
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		//incorrect
		go thumbnail.ImageFile(f) // NOTE: ignoring errors
	}
}

// makeThumbnails3 makes thumbnails of the specified files in parallel.
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f)
	}

	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			//上面的循环goroutine还没有执行完，有一个报错就返回
			return err // NOTE: incorrect: goroutine leak!
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		//有一个报错就返回nil
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)

	//goroutine计数器
	var wg sync.WaitGroup // number of working goroutines

	for f := range filenames {
		//在goroutine worker执行前执行
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			//有一个报错同样就直接返回nil
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	// closer  跟上面的主goroutine必须是并行执行
	go func() {
		wg.Wait()
		//等所有的worker goroutine执行完后执行 close(sizes)
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
