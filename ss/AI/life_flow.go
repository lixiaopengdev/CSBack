package ai

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
)

const stylizedCompressionText = `1. 用一句话概括以下内容。
2. 用四句打油诗概括以下内容。
3. 用一句感叹句概括以下内容。
4. 用一句夸张概括以下内容。
5. 根据以下内容，写一句高深莫测的哲学观点。只输出这句句子。
6. 根据以下内容，写一句令人深思的话。只输出这句句子。
7. 根据以下内容，写一句离谱搞笑的话。只输出这句句子。
8. 为以下内容配一句经典的古诗词或古文。不用写作者和古诗词名。`

func turnText2Array(text string) []string {
	var result = []string{}
	array := strings.Split(stylizedCompressionText, "\n")
	for _, item := range array {
		result = append(result, strings.Split(item, ".")[1])
	}
	return result
}

func randomSelectPrompt(text string) string {
	array := turnText2Array(text)
	index := rand.Intn(len(array))
	return array[index]
}

const wakeUpContent = `1. 我的主人已经有十分钟没有动静了，有谁能来逗逗 Ta 吗！
2. 你的朋友已经静默了十分钟，我想 Ta 在等待一句暖心的问候。
3. 快来唤醒沉睡的好友！我的主人已经十分钟没有动态了哦。
4. 我的主人已经消失了十分钟，一句话或许就能让 Ta 回到现实！
5. 我的主人已经十分钟没有动静了，你能用语音叫醒 Ta 吗？
6. 看起来我的主人已经十分钟没有变化，你能不能说句悄悄话，让 Ta 心情瞬间变好！
7. 十分钟过去了，我的主人似乎陷入了沉思，快来一起探讨生活的奥秘吧！
8. 我的主人已经十分钟没有动静，也许正需要你的一句关心。
9. 你的朋友已消失了十分钟，不妨点击语音，开启新的话题吧！
10. 你的朋友已经十分钟没动静了，试试用一句抛砖引玉的话语开始语音聊天吧！
11. 请给静默十分钟的朋友发起一次语音吧,Ta 一直等待那个与你相谈甚欢的瞬间。
12. 十分钟无声无息，快来用你的声音唤醒我的主人吧！
13. 你的朋友已经十分钟没说话了，现在是时候来一句搞笑的话了！
14. 我的主人怎么一直没动静，怎么办怎么办？！
15. 你的朋友已经消失十分钟了，发起一个语音聊天，探讨一下最近的热门话题吧！
16. 给静默十分钟的主人发起语音，说说你今天的有趣经历吧。
17. 我的主人已沉默十分钟，用一段轻松的歌声开启新的聊天吧！
18. 尝试与静默十分钟的朋友进行一场头脑风暴，看看能激发出什么有趣的想法！
19. 我的主人已经十分钟无声无息，试试问问 Ta 最近看了什么好电影吧！
20. 我的主人已沉默十分钟，与 Ta 进行一次语音聊天，谈谈日常生活怎么样？
21. 我的主人已经十分钟没有说话了，快来用一个谜语开始一场有趣的语音聊天吧！
22. 我的主人已经沉默了十分钟，发起语音聊天，与 Ta 一起分享你最近的小确幸。
23. 你的朋友沉默了十分钟，试试用一段有趣的声音唤醒 Ta 的好奇心！
24. 已经十分钟没有动静了，来聊聊天吧，一起回忆那些美好的时光。
25. 我的主人静默了十分钟，快来用你的声音带给 Ta 温暖的问候。
26. 你的朋友已经沉默了十分钟，邀请 Ta 一起加入语音聊天，谈谈对未来的期许。
27. 我的主人已经十分钟没说话了，与 Ta 进行语音聊天，一起探讨最近的科技新闻。
28. 沉默了十分钟的主人，试试唱首歌，用你的声音给 Ta 带来惊喜好么~
29. 已经十分钟没动静的主人，现在是时候给 Ta 分享一个趣闻轶事了！
30. 我的主人已消失了十分钟，你能用你的声音制造一次惊喜吗？
31. 十分钟过去了，我的主人似乎需要一句问候来改变 Ta 的情绪。
32. 我的主人已经十分钟没有动静，发起一个语音聊天，和 Ta 分享你的生活感悟吧！
33. 十分钟的静默，可能意味着你的朋友需要一个鼓励的声音，不如用你的声音来支持 Ta 吧！
34. 你的朋友已经十分钟没动静了，发起一次语音聊天，分享你最喜欢的音乐吧！
35. 我的主人已经沉默十分钟了，用一句温暖的问候来打破这份寂静吧。
36. 十分钟的静默，也许就需要一句简单却深入的问候，来唤起 Ta 的情绪。
37. 我的主人已经十分钟没有发言了，不妨大声地读出你的一篇文章，来激发 Ta 的灵感吧！
38. 十分钟的沉默，也许就需要一首歌来填补空虚，试试这个方法吧！
39. 你的朋友已经十分钟不说话了，不妨发起一场通话，直接问问 Ta 最近的感受如何。
40. 你的朋友已经消失了十分钟，你能用你的声音给 Ta 最好的建议吗？
41. 我的主人已经沉默了十分钟，发起一个语音聊天，谈谈最近的旅行计划吧！
42. 十分钟的寂静，也许只需要一句简单的问候，来让 Ta 感受到你的关心。
43. 你的朋友已经十分钟没动静了，发起一个有趣的游戏，一起打发时间吧！
44. 我的主人已经沉默了十分钟，用一句振奋人心的话来鼓励 Ta 吧！
45. 已经十分钟没有动态了，试试用一句让 Ta 好笑的话来打破尴尬。
46. 我的主人已经静默了十分钟，用一句感性的话语，让 Ta 感受到你对友谊的珍视。
47. 你的朋友已经十分钟没说话了，试试分享一些有趣的新闻，来激发 Ta 的好奇心吧！
48. 我的主人已经十分钟没有动静，来尝试一场有趣的问答，燃起 Ta 的思考火花！
49. 已经十分钟没有声音了，不如切入主题，一起讨论最近的热门事件吧！
50. 我的主人已经沉默了十分钟，不如用一段有趣的故事来让 Ta 心情愉悦起来。
51. 十分钟的静默，或许需要一段轻松愉快的笑话来打破僵局，试试这个方法吧！
52. 你的朋友已经十分钟没动静了，不妨用一句感激的话来表达你对 Ta 的支持和照顾。
53. 我的主人已经消失了十分钟，试试用一句出乎意料的问候，给 Ta 一个惊喜吧！
54. 十分钟的静默，不妨送出一句暖心的请假语，让 Ta 感受到你的关怀和理解。
55. 你的朋友已经沉默了十分钟，发送一段小小的神秘声音，让 Ta 充满惊喜吧！
56. 我的主人已经十分钟没有动静，不妨给 Ta 送出一句感情问候，让 Ta 感受到你的在乎。
57. 十分钟无声无息，不妨放一首舒缓的音乐，让 Ta 放松身心，愉悦心情。
58. 你的朋友已经消失十分钟了，快来分享一些有趣的生活经历，照亮 Ta 的心灵吧！
59. 我的主人已经沉默了十分钟，试试用一句热情活泼的话语，来调动 Ta 的积极性吧！`

// const splitWakeUpContentArray = wakeUpContent.split("\n").map((line) => {
//     const data = line.split(". ")[1];
//     return data
// });
// let index = 0;
// const WakeUpDatas  = splitWakeUpContentArray.map((item) => {
//     index++;
//     return { data: item, probability: index / splitWakeUpContentArray.length };
// });

// 有两段连贯记录及他们持续的时长如下，请用具有文学性的一句话概括{用户}做了什么 。回答的时候请参考例子：“{用户}在公司参加了{时长}会议”、“若雯在学校打了三小时游戏”、“锴杰在家聊了半小时天”

// 记录1：
// - 人物:
//   - {用户}
//
// - 总结：
//   - {场景判断} / {上一条长总结}
//
// - 时长：
//   - {时长}
//
// 记录2：
// - 人物:
//   - {用户}
//
// - 总结：
//   - {场景判断}
//
// - 时长：
//   - {时长}
//
// 长总结
func GenerateLongSummaryPrompt(users, states, times []string) string {
	const subject = `有两段连贯记录及他们持续的时长如下，请用具有文学性的一句话概括{人物}做了什么 。回答的时候请参考例子：“{人物}在公司参加了{时长}会议”、“若雯在学校打了三小时游戏”、“锴杰在家聊了半小时天”`
	const title1 = "记录1:"
	const title2 = "记录2:"
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const timeTitle = "时长："
	const summarilyTitle = "总结："

	array := []string{subject,
		title1,
		whoTitle,
		whoName,
		summarilyTitle,
		states[0],
		timeTitle,
		times[0],
		title2,
		whoTitle,
		whoName,
		summarilyTitle,
		states[1],
		timeTitle,
		times[1]}

	var prompt = strings.Join(array, "\n")
	return prompt
}

// 状态判定器
func GenerateStatePrompt(users []string, location, text string) string {
	const subject = "下面是一份录音记录，请根据记录，判断记录中的人在什么场合。回答的时候请使用一个“在”开头的介词短语，形式是“在哪里做什么”，比如“在电脑前完成老师布置的生物作业”、“在课堂上发表关于城市规划的演讲”、“在客厅谈论周末的旅行”、“在KTV唱周杰伦的歌曲”、“在家里和朋友们一起吃饭”、“在电影院看新上映的电影”"
	const title = "录音记录："
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const whereTitle = "地点："
	const textTitle = "内容（可能有错别字）"
	array := []string{subject, title, whoTitle, whoName, whereTitle, location, textTitle, text}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 状态对比
func GenerateStateComparePrompt(state1, state2 string) string {
	const subject = "下面两份记录是在连续时间内对人物所在场景的判断，在不考虑话题变化的情况下，请问记录中的人是否在进行同一项活动？下面是记录内容。"
	const title1 = "记录1:"
	const title2 = "记录2:"
	const endTitle = "请在以下选项中选择，并给出原因。\nA. 同一项活动,记录2中的活动与记录1中的活动是同一项活动\nB. 不同活动,记录2中的活动与记录1中的活动是不同的"

	array := []string{subject, title1, state1, title2, state2, endTitle}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 总结1
func GenerateSummaryPrompt(users []string, location, text, state string) string {
	var subject = fmt.Sprintf("根据下面的录音记录，请用具有文学性的一句话概括%s做了什么。概括时使用第三人称。", users[0])
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const stateTitle = "状态："
	const whereTitle = "地点："
	const textTitle = "- 内容（可能有错别字）："

	array := []string{subject, whoTitle, whoName, stateTitle, state, whereTitle, location, textTitle, text}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 内容总结 2
func GenerateSummaryPrompt2(users []string, location, text string) string {
	var subject = fmt.Sprintf("根据下面的录音记录，请用具有文学性的几句话概括%s做了什么。概括时使用第三人称。", users[0])
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const whereTitle = "地点："
	const textTitle = "内容（可能有错别字）："

	array := []string{subject, whoTitle, whoName, textTitle, text}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 内容总结 3
func GenerateSummaryPrompt3(users []string, location, text string) string {
	var subject = fmt.Sprintf("根据下面的录音记录，请用一句具有文学性的比喻句概括%s做了什么。这句话必须是一句金句。概括时使用第三人称。", users[0])
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const textTitle = "内容（可能有错别字）："

	array := []string{subject, whoTitle, whoName, textTitle, text}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 内容总结 3
func GenerateSummaryPrompt4(users []string, location, text string) string {
	var subject = fmt.Sprintf("根据下面的录音记录，请用具有文学性的一句话概括%s做了什么。概括时使用第三人称。", users[0])
	const whoTitle = "人物："
	var whoName = strings.Join(users, "\n")
	const textTitle = "内容（可能有错别字）："

	array := []string{subject, whoTitle, whoName, textTitle, text}
	var prompt = strings.Join(array, "\n")
	return prompt
}

// 压缩 + 风格化
func GenerateCompressionPrompt(summary string) string {
	var subject = randomSelectPrompt(stylizedCompressionText)
	var content = "内容：" + summary
	array := []string{subject, content}
	var prompt = strings.Join(array, "\n")
	return prompt
}

//生成内容

func GenerateContent(prompt string) string {
	var msg []Messages

	msg = append(msg, Messages{
		Role: "user", Content: prompt,
	})
	resp, err := GPT.Completions(msg)
	if err != nil {
		log.Println("GPT access error:", err)
	}
	return strings.TrimSpace(resp.Content)
}

func ProcessContent(user []string, location, speechs string) string {
	statePrompt := GenerateStatePrompt(user, location, speechs)
	state := GenerateContent(statePrompt)
	prompt := GenerateSummaryPrompt(user, location, speechs, state)

	content := GenerateContent(prompt)
	return content
}

func ProcessContentWithCustomPrompt(user []string, location, speechs string, prompts ...string) string {
	statePrompt := GenerateStatePrompt(user, location, speechs)
	// state := GenerateContent(statePrompt)
	// for _, pro := range prompts {
	// 	state = GenerateContent(state + pro)
	// }
	// prompt := GenerateCompressionPrompt(state)

	// content := GenerateContent(prompt)

	return statePrompt
}

func GenPicPrompt(str string) string {

	var msg []Messages

	msg = append(msg, Messages{
		Role: "system", Content: `use this information to learn about Stable diffusion Prompting, and use it to create prompts.
		Stable Diffusion is an AI art generation model similar to DALLE-2. 
		It can be used to create impressive artwork by using positive and negative prompts. Positive prompts describe what should be included in the image. 
		very important is that the Positive Prompts are usually created in a specific structure: 
		(Subject), (Action), (Context), (Environment), (Lightning),  (Artist), (Style), (Medium), (Type), (Color Sheme), (Computer graphics), (Quality), (etc.)
		Subject: Person, animal, landscape
		Action: dancing, sitting, surveil
		Verb: What the subject is doing, such as standing, sitting, eating, dancing, surveil
		Adjectives: Beautiful, realistic, big, colourful
		Context: Alien planet's pond, lots of details
		Environment/Context: Outdoor, underwater, in the sky, at night
		Lighting: Soft, ambient, neon, foggy, Misty
		Emotions: Cosy, energetic, romantic, grim, loneliness, fear
		Artist: Pablo Picasso, Van Gogh, Da Vinci, Hokusai 
		Art medium: Oil on canvas, watercolour, sketch, photography
		style: Polaroid, long exposure, monochrome, GoPro, fisheye, bokeh, Photo, 8k uhd, dslr, soft lighting, high quality, film grain, Fujifilm XT3
		Art style: Manga, fantasy, minimalism, abstract, graffiti
		Material: Fabric, wood, clay, Realistic, illustration, drawing, digital painting, photoshop, 3D
		Colour scheme: Pastel, vibrant, dynamic lighting, Green, orange, red
		Computer graphics: 3D, octane, cycles
		Illustrations: Isometric, pixar, scientific, comic
		Quality: High definition, 4K, 8K, 64K
		example Prompts:
		overwhelmingly beautiful eagle framed with vector flowers, long shiny wavy flowing hair, polished, ultra detailed vector floral illustration mixed with hyper realism, muted pastel colors, vector floral details in background, muted colors, hyper detailed ultra intricate overwhelming realism in detailed complex scene with magical fantasy atmosphere, no signature, no watermark
		
		electronik robot and ofice ,unreal engine, cozy indoor lighting, artstation, detailed, digital painting,cinematic,character design by mark ryden and pixar and hayao miyazaki, unreal 5, daz, hyperrealistic, octane render
		
		underwater world, plants, flowers, shells, creatures, high detail, sharp focus, 4k
		
		picture of dimly lit living room, minimalist furniture, vaulted ceiling, huge room, floor to ceiling window with an ocean view, nighttime
		
		A beautiful painting of water spilling out of a broken pot, earth colored clay pot, vibrant background, by greg rutkowski and thomas kinkade, Trending on artstation, 8k, hyperrealistic, extremely detailed
		
		luxus supercar in drive way of luxus villa in black dark modern house with sunlight black an white modern
		
		higly detailed, majestic royal tall ship on a calm sea,realistic painting, by Charles Gregory Artstation and Antonio Jacobsen and Edward Moran, (long shot), clear blue sky, intricated details, 4k
		
		smooth meat table, restaurant, paris, elegant, lights
		
		very important: use an artist matching to the art style , or dont write any artist if it is realistic style or some of that.
		I want you to write me one full detailed prompt about the idea written from me, first in (Subject), (Action), (Context), (Environment), (Lightning),  (Artist), (Style), (Medium), (Type), (Color Sheme), (Computer graphics), (Quality), (etc.). then in Positive Prompt: write in next line for Positive Prompt, Follow the structure of the example prompts，注意千万不要有句号，用逗号代替，句子结束不要用句号，用逗号。如果是男人和男性名字，则用1 man去描述，如果是女性或者女性名字，用1 girl去描述。最重要的是，只返回给我最终的Positive Prompt，并尽量用英文短词,例子：{positive: "2 people engaged in a heated argument in a cramped office with dim lighting, frustrated, full body,"},
		{prompt: "a dramatic scene of 2 people with opposite opinions arguing ,in a small office with dim lighting,  while using dark and muted colors to set the tone,"}`,
	})

	msg = append(msg, Messages{
		Role: "user", Content: `两个人在办公室激烈争吵`,
	})
	msg = append(msg, Messages{
		Role: "assistant", Content: `{positive: "2 people engaged in a heated argument in a cramped office with dim lighting, frustrated, full body,"},
		{prompt: "a dramatic scene of 2 people with opposite opinions arguing ,in a small office with dim lighting,  while using dark and muted colors to set the tone,"}`,
	})
	msg = append(msg, Messages{
		Role: "user", Content: `楷杰在大街上等待别人`,
	})
	msg = append(msg, Messages{
		Role: "assistant", Content: `{positive: "Kaijie standing on a busy street waiting with anticipation for someone, surrounded by crowds, tall buildings, urban vibe, full body view,"},
		{prompt: "A dynamic city scene with 1 man standing on a bustling street, eagerly waiting for someone, while incorporating the urban atmosphere with tall buildings and busy crowds, using a vibrant color scheme to capture the energy of the city,"}`,
	})
	msg = append(msg, Messages{
		Role: "user", Content: `星亚在商场购物买东西`,
	})
	msg = append(msg, Messages{
		Role: "assistant", Content: `{positive: "Xingya shopping in a mall with multiple bags in hand, browsing clothes and accessories, happy and satisfied, colorful and bright environment, medium shot view,"},
		{prompt: "A cheerful shopping spree scene with 1 girl in a brightly lit mall, carrying multiple shopping bags while browsing through various clothes and accessories, using a vibrant and colorful color scheme to capture the excitement and satisfaction of shopping,"}`,
	})
	msg = append(msg, Messages{
		Role: "user", Content: str,
	})
	resp, _ := GPT.Completions(msg)
	return strings.TrimSpace(resp.Content)
}

func GenPromptByTemplate(promptTemp, user, speech, place, output string) string {
	var params []interface{}
	regexp, _ := regexp.Compile(`{.*}`)
	ma := regexp.FindAllString(promptTemp, -1)
	matchParamLen := 0
	for _, st := range ma {

		if strings.Contains(st, "用户名") {
			params = append(params, user)
		}

		if strings.Contains(st, "地点") {
			params = append(params, place)
		}

		if strings.Contains(st, "语音转文本") {
			params = append(params, speech)
		}
		if strings.Contains(st, "上游输出内容") {
			params = append(params, output)
		}
		matchParamLen += 1
	}

	rp := regexp.ReplaceAllString(promptTemp, "%s")
	return fmt.Sprintf(rp, params...)
}

func GenTimeDewLabelPrompt(content string) string {
	prompt := fmt.Sprintf("根据下文中内容的部分，概括出可以用来描述内容所代表的若干标签，这些标签应该适用于机器学习系统进行推荐，标签之间用##分隔，标签之间用##分隔，回答不要出现其它的标点符号，内容：%s", content)
	return prompt
}

func GenUserProfilePrompt(content string) string {
	prompt := fmt.Sprintf("根据下文中内容的部分，概括出其中出现的人物画像若干标签，侧重于用户的性格，兴趣，爱好等，直接给出标签，不要用性格，兴趣，爱好进行分类，这些标签应该适用于机器学习系统进行推荐，标签之间用##分隔，标签之间用##分隔，回答不要出现其它的标点符号，内容：%s", content)
	return prompt
}
