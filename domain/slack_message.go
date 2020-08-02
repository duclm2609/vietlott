package domain

import (
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

func mapJackpotNumber(prize Jackpot) Text {
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

func MapFrom(mega Mega645Result) SlackMessage {
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

func MapFromCompareResult(result Mega645CompareResult) SlackMessage {
	var blocks []Block
	if result.JackpotPrize[0] != 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf(":tada: *Trúng Jackpot rồiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiiii!!!!!!!!!!!!!!!!!!!*\n%+v", result.JackpotPrize),
			},
		})
	}
	if len(result.FirstPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải nhất: %d", len(result.FirstPrize)),
			},
		})
	}
	if len(result.SecondPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải nhì: %d", len(result.SecondPrize)),
			},
		})
	}
	if len(result.ThirdPrize) > 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: fmt.Sprintf("Giải ba: %d", len(result.ThirdPrize)),
			},
		})
	}
	if len(blocks) == 0 {
		blocks = append(blocks, Block{
			Type: BlockTypeSection,
			Text: Text{
				Type: TextTypeMarkdown,
				Text: "Chưa trúng phát nào, chúc may mắn lần sau nhé ông bạn!",
			},
		})
	}
	return SlackMessage{Blocks: blocks}
}
