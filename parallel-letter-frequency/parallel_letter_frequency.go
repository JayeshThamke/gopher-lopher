package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the text concurrently
// and returns FreqMap.
func ConcurrentFrequency(texts []string) FreqMap {
	freqMap := FreqMap{}

	mutex := &sync.Mutex{}
	var wg sync.WaitGroup

	for i := 0; i < len(texts); i++ {

		// for each iteration we will be adding 1 waitgroup
		// and indicating completion
		wg.Add(1)
		go func(text string, wg *sync.WaitGroup) {
			for _, r := range text {
				mutex.Lock()
				freqMap[r]++
				mutex.Unlock()
			}
			wg.Done()
		}(texts[i], &wg)
	}
	wg.Wait()
	return freqMap
}
