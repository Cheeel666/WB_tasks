package handlers

import (
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type TestUser struct {
}

func sendBanRequest(count int) {
	var body io.Reader
	client := &http.Client{}
	for i := 0; i < count; i++ {

		url := "http://localhost:8080/api/v1/user/" + strconv.Itoa(rand.Intn(10000)) + "/block"
		req, err := http.NewRequest(http.MethodPut, url, body)
		if err != nil {
			logrus.Error(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		resp, err := client.Do(req)
		if err != nil {
			logrus.Error(err)
		}

		if resp.StatusCode != 200 {
			// logrus.Info("Request finished with status: ", resp.StatusCode)

		}
	}
}

// TestUserBan in service
func (t *TestUser) TestUserBan(c *fasthttp.RequestCtx) {
	// countWorkersInterface := c.Value("workerCount")
	// countWorkers, err := strconv.Atoi(fmt.Sprintf("%v", countWorkersInterface))

	// if err != nil {
	// 	c.Write([]byte("invalid number"))
	// 	return
	// }
	countWorkers := 10
	for i := 0; i < countWorkers; i++ {
		go sendBanRequest(10)
	}
	c.Write([]byte("Finished testing ban users"))
}

// // TestUserUnban in service
// func (t *TestUser) TestUserUnban(c *fasthttp.RequestCtx) {
// 	countWorkers := c.Value("workerCount")

// 	c.Write([]byte("Finished testing unban users"))
// }
