package apps

import (
	"crypto/tls"
	"fmt"
	"github.com/cloudfoundry/cf-test-helpers/v2/cf"
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/cloudfoundry/cf-acceptance-tests/helpers/app_helpers"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers/random_name"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = AppsDescribe("WebSocket", func() {
	var appName string
	var wsURL string

	BeforeEach(func() {
		appName = random_name.CATSRandomName("APP")

		// Push WebSocket test app to CF
		Expect(cf.Cf("push", appName, "--no-start", "-p", "../../assets/websocket", "-f", "../../assets/websocket/manifest.yml").Wait()).To(Exit(0))
		Expect(cf.Cf("start", appName).Wait()).To(Exit(0))

		// Get the app URL and create the WebSocket URL
		appInfo := app_helpers.AppReport(appName)
		wsURL = fmt.Sprintf("wss://%s", appInfo.Routes[0].Host)
	})

	AfterEach(func() {
		// Clean up app after test
		Expect(cf.Cf("delete", appName, "-f").Wait()).To(Exit(0))
	})

	It("establishes a WebSocket connection and echoes messages", func() {
		// Use Gorilla WebSocket for client connection
		dialer := websocket.Dialer{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Skip TLS verification in tests
		}

		// Connect to WebSocket server
		conn, _, err := dialer.Dial(wsURL, http.Header{})
		Expect(err).NotTo(HaveOccurred(), "Failed to connect to WebSocket server")
		defer conn.Close()

		// Send a message over WebSocket
		message := "Hello WebSocket!"
		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		Expect(err).NotTo(HaveOccurred(), "Failed to send message over WebSocket")

		// Read the response (should echo the message back)
		_, response, err := conn.ReadMessage()
		Expect(err).NotTo(HaveOccurred(), "Failed to read message from WebSocket server")
		Expect(string(response)).To(Equal("echo: " + message))

		// Confirm WebSocket closes cleanly
		err = conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		Expect(err).NotTo(HaveOccurred())
	})
})
