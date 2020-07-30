db.createUser(
    {
        user: "test",
        pwd: "12345678",
        roles: [
            {
                role: "readWrite",
                db: "vietlott"
            }
        ]
    }
);

db.parser_config.insert(
    {
        "mega645_selector": {
            "url": "https://vietlott.vn/vi/trung-thuong/ket-qua-trung-thuong/645",
            "draw_info_selector": {
                "base": "#divLeftContent > div > div.header > div > div > h5",
                "draw_id": "b:nth-child(1)",
                "draw_date": "b:nth-child(2)",
            },
            "jackpot_prize_selector": "#divRightContent > div > div.chitietketqua_table.maga645_table > div.gt_jackpot > div > div.col-md-7 > div > h3",
            "jackpot_selector": "#divLeftContent > div > div.content > div > div.day_so_ket_qua_v2",
            "jackpot_winner_selector": "#divRightContent > div > div.chitietketqua_table.maga645_table > div.table-responsive > table > tbody > tr:nth-child(1) > td:nth-child(3)"
        }
    }
)