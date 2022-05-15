package telegram

type Response struct {
	Result []Update `json: "result"`
}

type Update struct {
	Update_Id int `json: "update_id"`
	Message  Message `json: "message"`
	Callback_Query CallbackQuery `json: "callback_query"`
}

type Message struct {
	Message_Id int `json: message_id`
	From From `json: "from"`
	Date int `json: "date"`
	Text string `json: "text"`
}

type From struct {
	Id int `json: "id"`
	Is_Bot bool `json: "is_bot"`
	First_Name string `json: "first_name"`
	Username string `json: "username"`
}

type CallbackQuery struct {
	Id string
	Message Message
	Data string 
}

// type InlineKeyboard struct {
// 	Inline_Keyboard [][]InlineButton `json: "inline_keyboard"`
// }

// type InlineButton struct {
// 	Text string `json: "text"`
// 	Callback_Data string `json: "callback_data"`
// }
