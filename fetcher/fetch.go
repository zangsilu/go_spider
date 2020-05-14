package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	//response, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//defer response.Body.Close()

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36") //浏览器中User-agent
	req.Header.Set("Cookie", "sid=9dabc394-526d-4928-8530-dadd6ca815bd; Hm_lvt_2c8ad67df9e787ad29dbd54ee608f5d2=1589259270; FSSBBIl1UgzbN7NO=5ZS3Dv_.M7wyjvtMaxYr7yHe6VtwAlZeIPBo4KGFJaUPUOsx.wg1gZVZ7JrSqNOkhcRrmIxvcYNml1yw95daYjA; Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2=1589293748; FSSBBIl1UgzbN7NP=5U.jawm__7olqqqmncV70YqIUwnNpSpwX3byS_HUfl_nDLNGzlerPcUDhfkLKZNO68mDqj4b0Lnjqio8zU1hS8IwRD8ft97XqXEPWlMyRa4XM0P5_B.jH.ie5mhVhym4Lngiur9cZhgnKuyR7AI6CUBHgrbz.H0feSvqIwppYIEjVz751I9Ih5cLp7qXdDlLMypR9vt2WVIA8Z8KEl9U_w.Ht1Ro.SVKkyy4TMatDOuctjzFOh7Kbly.3x.1AkTm_93DkPOaBrgNFHzreiXJtZb")
	response, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error: status code", response.StatusCode)
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}
