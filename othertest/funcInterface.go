package main

/**
 * @description 
 * @time 2018/5/10 0:02
 * @version 
 */
type Handler interface {
	Do(k,v interface{})
}

func Each(m map[interface{}]interface{}, handler Handler)  {
	if m != nil && len(m) > 0{
		for k, v := range m {
			handler.Do(k,v)
		}
	}
}