package twitch

import (
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/zigzter/league-predictions/utils"
)

func GenerateRandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", b)
}

var serverAddr = "localhost:9001"
var clientId = "qwodvkbo5rdiq1bnycx63cr3981gn6"

type (
	ServerStartMsg     struct{}
	ServerStartedMsg   struct{}
	AuthOpenMsg        struct{}
	AuthOpenedMsg      struct{}
	TokenReceiveMsg    struct{}
	TokenReceivedMsg   struct{}
	ProcessCompleteMsg struct{}
)

func StartLocalServer(ready chan<- struct{}, externalMsgs chan tea.Msg) tea.Cmd {
	return func() tea.Msg {
		http.HandleFunc("/token/", func(w http.ResponseWriter, r *http.Request) {
			token := strings.TrimPrefix(r.URL.Path, "/token/")
			if token != "" {
				utils.SaveConfig(utils.TwitchTokenKey, token)
				fmt.Fprintln(w, "Token received, you can close this window.")
				externalMsgs <- TokenReceivedMsg{}
			} else {
				fmt.Fprintln(w, "Failed to retrieve token.")
				externalMsgs <- TokenReceivedMsg{}
			}
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			filePath := "./twitch/token.html"
			http.ServeFile(w, r, filePath)
		})

		httpServer := &http.Server{Addr: serverAddr}
		go func() {
			ready <- struct{}{}
			if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("ListenAndServe error: %v", err)
			}
		}()
		return ServerStartedMsg{}
	}
}

func PromptTwitchAuth() tea.Cmd {
	return func() tea.Msg {
		scopes := []string{
			"channel:moderate",
		}
		redirectUrl := serverAddr
		scope := strings.Join(scopes, " ")
		state := GenerateRandomString(10)
		url := fmt.Sprintf(
			"https://id.twitch.tv/oauth2/authorize?response_type=token&client_id=%s&redirect_uri=http://%s&scope=%s&state=%s",
			clientId,
			redirectUrl,
			scope,
			state,
		)
		utils.OpenBrowser(url)
		return AuthOpenedMsg{}
	}
}
