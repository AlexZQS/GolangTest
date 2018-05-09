package mock

import "fmt"

/**
 * @description 
 * @time 2018/4/16 22:45
 * @version 
 */
type Retriever struct {
	Contents string
}

func (r *Retriever) Get(url string) string {
	r.Contents += url
	return fmt.Sprintf(r.Contents)
}
