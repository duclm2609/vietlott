package slack

import (
	"dev.duclm/vietlott/parser/domain"
	"fmt"
)

const BlockTypeSection = "section"
const TextTypeMarkdown = "mrkdwn"
const DateLayout = "02/01/2006"

type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type Block struct {
	Type string `json:"type"`
	Text Text   `json:"text"`
}

type SlackMessage struct {
	Blocks []Block `json:"blocks"`
}

func mapDrawIdText(drawId string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":fire: Kỳ quay: *%s*", drawId),
	}
}

func mapDrawDateText(drawDate string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":date: Ngày quay thưởng: *%s*", drawDate),
	}
}

func mapJackpotPrize(prize string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":dollar: Giá trị Jackpot: *%s* VNĐ", prize),
	}
}

func mapJackpotNumber(prize domain.Jackpot) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":slot_machine: Jackpot: *%s*", prize),
	}
}

func mapNumberOfJackpot(number string) Text {
	return Text{
		Type: TextTypeMarkdown,
		Text: fmt.Sprintf(":medal: Số lượng giải Jackpot: *%s*", number),
	}
}

func MapFrom(mega domain.Mega645Result) SlackMessage {
	return SlackMessage{
		Blocks: []Block{
			{
				Type: BlockTypeSection,
				Text: mapDrawIdText(mega.DrawId),
			},
			{
				Type: BlockTypeSection,
				Text: mapDrawDateText(mega.DrawDate.Format(DateLayout)),
			},
			{
				Type: BlockTypeSection,
				Text: mapJackpotPrize(mega.Prize),
			},
			{
				Type: BlockTypeSection,
				Text: mapJackpotNumber(mega.Jackpot),
			},
			{
				Type: BlockTypeSection,
				Text: mapNumberOfJackpot(mega.Winner),
			},
		},
	}
}
