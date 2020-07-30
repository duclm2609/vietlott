package mongo

import "dev.duclm/vietlott/parser/domain"

func (h handler) Get() (domain.ParserConfig, error) {
	//TODO: replace with implementation
	return domain.ParserConfig{
		Mega645Selector: domain.Mega645Selector{
			Url: "https://vietlott.vn/vi/trung-thuong/ket-qua-trung-thuong/645.html",
			DrawInfo: domain.DrawInfoSelector{
				Base:     "#divLeftContent > div > div.header > div > div > h5",
				DrawId:   "b:nth-child(1)",
				DrawDate: "b:nth-child(2)",
			},
			JackpotPrizeSelector: "#divRightContent > div > div.chitietketqua_table.maga645_table > div.gt_jackpot > div > div.col-md-7 > div > h3",
			JackpotSelector:      "#divLeftContent > div > div.content > div > div.day_so_ket_qua_v2",
			JackpotWinner:        "#divRightContent > div > div.chitietketqua_table.maga645_table > div.table-responsive > table > tbody > tr:nth-child(1) > td:nth-child(3)",
		},
	}, nil
}
