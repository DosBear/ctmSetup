package config

type Soft struct {
	File    string
	URL     string
	Folder  string
	Name    string
	Checked bool
	Version string
}

func GetConfig() []Soft {
	var softList []Soft

	softList = append(softList, Soft{
		File:    "setup_dl.exe",
		URL:     "https://ftp.ctm.ru/DCL/SFX/setup_dl.exe",
		Folder:  "DCL",
		Name:    "ВЭД-Декларант",
		Checked: true,
		Version: ""})
	softList = append(softList, Soft{
		File:    "setup_me.exe",
		URL:     "https://ftp.ctm.ru/MONITOR_ED/SFX/setup_me.exe",
		Folder:  "MONITOR_ED",
		Name:    "Монитор ЭД",
		Checked: true,
		Version: ""})
	softList = append(softList, Soft{
		File:    "setup_pa.exe",
		URL:     "https://ftp.ctm.ru/PAYMENT/SFX/setup_pa.exe",
		Folder:  "PAYMENT",
		Name:    "ВЭД-Платежи",
		Checked: true,
		Version: ""})
	softList = append(softList, Soft{
		File:    "setup_vi.exe",
		URL:     "https://ftp.ctm.ru/VEDINFO/SFX/setup_vi.exe",
		Folder:  "VEDINFO",
		Name:    "ВЭД-Инфо",
		Checked: true,
		Version: ""})
	softList = append(softList, Soft{
		File:    "setup_al.exe",
		URL:     "https://ftp.ctm.ru/ALPHABET/SFX/setup_al.exe",
		Folder:  "ALPHABET",
		Name:    "ВЭД-Алфавит",
		Checked: true,
		Version: ""})
	softList = append(softList, Soft{
		File:    "setup_st.exe",
		URL:     "https://ftp.ctm.ru/STS/SFX/setup_st.exe",
		Folder:  "STS",
		Name:    "ВЭД-Склад",
		Checked: true,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_cl.exe",
		URL:     "https://ftp.ctm.ru/CONTROL/SFX/setup_cl.exe",
		Folder:  "CONTROL",
		Name:    "ВЭД-Контроль",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_lw.exe",
		URL:     "https://ftp.ctm.ru/CONTROLS32/SFX/setup_lw.exe",
		Folder:  "CONTROLS32",
		Name:    "ВЭД-Контроль ГТД",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_td.exe",
		URL:     "https://ftp.ctm.ru/TRANSP/SFX/setup_td.exe",
		Folder:  "TD",
		Name:    "Транспортные документы",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_rk.exe",
		URL:     "https://ftp.ctm.ru/RAILATLAS/SFX/setup_rk.exe",
		Folder:  "RAILATLAS",
		Name:    "Rail-Атлас",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_ri.exe",
		URL:     "https://ftp.ctm.ru/RAILINFO/SFX/setup_ri.exe",
		Folder:  "RAILINFO",
		Name:    "Rail-Инфо",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_rt.exe",
		URL:     "https://ftp.ctm.ru/RAILTRF/SFX/setup_rt.exe",
		Folder:  "RAILTARIF",
		Name:    "Rail-Тариф",
		Checked: false,
		Version: ""})

	softList = append(softList, Soft{
		File:    "setup_rt.exe",
		URL:     "https://ftp.ctm.ru/RAILTRFRU/SFX/setup_rr.exe",
		Folder:  "RAILTARIFRU",
		Name:    "Rail-Тариф Россия",
		Checked: false,
		Version: ""})

	return softList
}
