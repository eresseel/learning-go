package main

func CheckNumber(num int) []string {
	messages := []string{}

	if num >= 10 {
		messages = append(messages, "Num is more than 10")
		if num > 15 {
			messages = append(messages, "Num is also more than 15")
		}
	} else {
		messages = append(messages, "Num is less than 10")
	}

	return messages
}
