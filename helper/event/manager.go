package event

var eventMap = make(map[string][]*Handler, 16)

type Handler struct {
	Handle func(interface{}) error
	Next   *Handler
}

func RegHandler(h *Handler, event string) {
	if _, ok := eventMap[event]; !ok {
		eventMap[event] = make([]*Handler, 1)
	}
	eventMap[event] = append(eventMap[event], h)
}

func Notify(event string, msg interface{}) {
	for _, handler := range eventMap[event] {
		go func(h *Handler) {
			err := h.Handle(msg)
			for ; err == nil && h.Next != nil; {
				h = h.Next
				err = h.Handle(msg)
			}
		}(handler)
	}
}
