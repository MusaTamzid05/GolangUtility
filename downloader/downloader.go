package downloader 

import (
	"fmt"
	"io/ioutil"
	"time"
	"net/http"
	"net"
)

type Downloader struct {
	response *http.Response
	timeoutDuration  time.Duration 
}


func (d *Downloader) Download(url string , timeout int) string {

	d.timeoutDuration= time.Duration(timeout) * time.Second
	var err error

	d.response , err = d.getResponse(url) 

	if err != nil {

		fmt.Println("Error : " , err)
		return ""

	}

	defer d.response.Body.Close()
	return d.getHTML()
}


func (d *Downloader) getHTML() string {

	fmt.Println("Status Code : " , d.response.StatusCode)
	bytes , _ := ioutil.ReadAll(d.response.Body)
	return string(bytes)
}


func (d *Downloader) getResponse(url string) (*http.Response , error ){


	t := http.Transport { Dial : d.timeoutHandler, }
	client := http.Client{ Transport : &t, }

	return client.Get(url)

}

func (d *Downloader) timeoutHandler(network string , host string) (net.Conn , error) {

	conn , err := net.DialTimeout(network ,  host , d.timeoutDuration)

	if err != nil {
		return nil , err
	}

	conn.SetDeadline(time.Now().Add(d.timeoutDuration))
	return conn , nil
}



func main() {
	fmt.Println("vim-go")
}
