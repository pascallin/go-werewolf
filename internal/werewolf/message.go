package werewolf

import "strings"

type Message struct {
	Content string
	Username string
}

var (
	separator = ":"
)

func MessageDecode(str string) Message {
	arr := strings.Split(str, separator)
	return Message{Content: arr[0]}
}

func MessageEncode(msg Message) string {
	return msg.Content
}