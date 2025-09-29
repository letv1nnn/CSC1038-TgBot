package botui

import "strings"

const About string = `The bot will provide multiple solutions to a lab problem: one using basic syntax and another using advanced syntax. It will also explain some C programming topics covered in the lecture. Additionally, by using the provided button, the bot will return a list of resources for Computer Science.`
const Resources string = `
	Computer Systems
	https://www.youtube.com/watch?v=Keducx5bp-g&list=PL0j-r-omG7i0-mnsxN5T4UcVS1Di0isqf&index=17

	https://www.nand2tetris.org

	OS
	https://pages.cs.wisc.edu/~remzi/OSTEP/

	https://pages.cs.wisc.edu/~remzi/Classes/537/Fall2021/

	https://pdos.csail.mit.edu/6.S081/2021/schedule.html

	Algo
	https://www.youtube.com/watch?v=oFVYVzlvk9c&list=PLUl4u3cNGP63EdVPNLG3ToM6LaEUuStEY&index=13

	https://t.me/tenfoundation/1217

	Math
	https://www.3blue1brown.com/#lessons

	https://ocw.mit.edu/course-lists/scholar-courses/

	https://www.khanacademy.org

	https://mathacademy.com/adult-students

	Networking

	http://gaia.cs.umass.edu/kurose_ross/wireshark.php

	https://www.youtube.com/playlist?list=PLoCMsyE1cvdWKsLVyf6cPwCLDIZnOj0NS

	DB
	https://15445.courses.cs.cmu.edu/fall2024/assignments.html

	https://www.youtube.com/watch?v=niLwbfE3V9Q&list=PLSE8ODhjZXjYDBpQnSymaectKjxCy6BYq&index=20

	Distributed Systems
	https://pdos.csail.mit.edu/6.824/schedule.html

	https://www.youtube.com/watch?v=UEAMfLPZZhE&list=PLeKd45zvjcDFUEv_ohr_HdUFe97RItdiB

	Security
	https://pwn.college/

	https://web.stanford.edu/class/cs253/

	https://61600.csail.mit.edu/2023/
`

func EscapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(text)
}
