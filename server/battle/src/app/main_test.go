package main_test

import (
	"app"
	"app/connection"
	"app/typhenapi/core"
	webapi "app/typhenapi/web/submarine"
	websocketapi "app/typhenapi/websocket/submarine"
	"bytes"
	"errors"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var p = fmt.Println

type clientSession struct {
	conn         *connection.Connection
	api          *websocketapi.WebSocketAPI
	disconnected chan struct{}
}

func newClientSession() *clientSession {
	serializer := typhenapi.NewJSONSerializer()
	session := &clientSession{
		conn:         connection.New(),
		disconnected: make(chan struct{}),
	}
	session.api = websocketapi.New(session, serializer, nil)
	session.conn.DisconnectHandler = func() {
		session.disconnected <- struct{}{}
	}
	session.conn.BinaryMessageHandler = func(data []byte) {
		if err := session.api.DispatchMessageEvent(data); err != nil {
			main.Log.Error(err)
		}
	}
	return session
}

func (s *clientSession) Send(data []byte) error {
	return s.conn.WriteBinaryMessage(data)
}

func (s *clientSession) connect(url string) error {
	_, err := s.conn.Connect(strings.Replace(url, "http", "ws", 1), nil)
	return err
}

func (s *clientSession) close() {
	s.conn.Close()
}

type webAPITransporter struct {
	serializer typhenapi.Serializer
}

func (m *webAPITransporter) RoundTrip(request *http.Request) (*http.Response, error) {
	response := &http.Response{Header: make(http.Header), Request: request}
	response.Header.Set("Content-Type", "application/json")
	data, _ := ioutil.ReadAll(request.Body)
	typhenType, statusCode := m.Routes(request.URL.Path, data)
	err := typhenType.Coerce()
	if err != nil {
		return response, err
	}

	response.StatusCode = statusCode
	body, _ := typhenType.Bytes(m.serializer)
	response.Body = ioutil.NopCloser(bytes.NewReader(body))
	if response.StatusCode >= 400 {
		return nil, errors.New("Server Error")
	}
	return response, nil
}

func newWebAPIMock(url string) *webapi.WebAPI {
	main.WebAPIRoundTripper = &webAPITransporter{typhenapi.NewJSONSerializer()}
	return main.NewWebAPI(url)
}

func newTestServer() (*httptest.Server, *main.Server) {
	main.Log.Level = logrus.PanicLevel
	main.WebAPIRoundTripper = &webAPITransporter{typhenapi.NewJSONSerializer()}
	gin.SetMode(gin.TestMode)
	rawServer := main.NewServer()
	server := httptest.NewServer(rawServer)
	return server, rawServer
}
