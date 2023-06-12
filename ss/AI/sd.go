package ai

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

const (
	// SDBASEURL = "http://192.168.50.168:7861/sdapi/v1/%s"
	// SDBASEURL = "http://zingy.land:7860/sdapi/v1/%s"
	sdengine = "sd"
)

func InitSD() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	SDHOST, SDHOSTExist := os.LookupEnv("SDHOST")
	if SDHOSTExist || len(SDHOST) > 0 {
		SDBASEURL = fmt.Sprintf("http://%s/sdapi/v1/", SDHOST) + "%s"
	}

}

var SDBASEURL = "http://zingy.land:7860/sdapi/v1/%s"
var SDT2IBASEURL = fmt.Sprintf(SDBASEURL, "txt2img")
var SDI2IBASEURL = fmt.Sprintf(SDBASEURL, "img2img")
var SDCLIPBASEURL = fmt.Sprintf(SDBASEURL, "interrogate")
var SDSRBASEURL = fmt.Sprintf(SDBASEURL, "extra-single-image")

type SDImageGenerationResponseBody struct {
	Images     []string `json:"images"`
	Parameters struct {
	} `json:"parameters"`

	Info string `json:"info"`
}

type SDI2ISuperResolutionResponseBody struct {
	HtmlInfo string `json:"html_info"`
	Image    string `json:"image"`
}

type SDT2ImageGenerationRequestBody struct {
	Prompt            string  `json:"prompt"`
	Steps             int     `json:"steps"`
	NegativePrompt    string  `json:"negative_prompt"`
	SamplerIndex      string  `json:"sampler_index"`
	EnableHR          bool    `json:"enable_hr"`
	DenoisingStrength float32 `json:"denoising_strength"`
	HrScale           int     `json:"hr_scale"`
	HrUpscaler        string  `json:"hr_upscaler"`
	Width             int     `json:"width"`
	Height            int     `json:"height"`
	CfgScale          int     `json:"cfg_scale"`
	HrSecondPassSteps int     `json:"hr_second_pass_steps"`
}

type SDI2ImageGenerationRequestBody struct {
	InitImages     []string `json:"init_images"`
	Prompt         string   `json:"prompt"`
	Steps          int      `json:"steps"`
	NegativePrompt string   `json:"negative_prompt"`
	SamplerIndex   string   `json:"sampler_index"`
}

// 超分
type SDI2ISuperResolutionRequestBody struct {
	ResizeMode                int     `json:"resize_mode"`
	UpscalingResize           float64 `json:"upscaling_resize"`
	Upscaler1                 string  `json:"upscaler_1"`
	ExtrasUpscaler2Visibility float32 `json:"extras_upscaler_2_visibility"`
	Image                     string  `json:"image"`
}

type SDPNGINFORequestBody struct {
	Image string `json:"image"`
	Model string `json:"model"`
}

type SDClipINFOResptBody struct {
	Caption string `json:"caption"`
}

func TrySDT2I(prompt string, upscale float64, sdBaseUrlFmt string) (string, error) {
	// Create a Resty Client
	client := resty.New()

	// POST JSON string
	// No need to set content type, if you have client level setting
	reqBody := SDT2ImageGenerationRequestBody{
		Prompt:         prompt,
		NegativePrompt: "EasyNegative,ng_deepnegative_v1_75t",
		SamplerIndex:   "DPM++ SDE Karras",
		Steps:          28,
		EnableHR:       false,
		Width:          768,
		Height:         768,
		CfgScale:       8,
		// EnableHR:          true,
		// DenoisingStrength: 0.6,
		// HrScale:           2,
		// HrUpscaler:        "4x_fatal_Anime_500000_G",
		// Width:             768,
		// Height:            768,
		// CfgScale:          8,
		// HrSecondPassSteps: 10,
		//16:9
		// Width:910
		// Height:512
	}
	r := &SDImageGenerationResponseBody{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(r).
		Post(fmt.Sprintf(sdBaseUrlFmt, "txt2img"))
	if err != nil {
		return "", err
	} else {
		if len(r.Images) == 0 {
			return "", errors.New("sd resp: " + resp.String())
		}
		fmt.Println("bs64 str", r.Images[0][0:20])
		imgbs64, _ := TrySuperResolution(r.Images[0], upscale, sdBaseUrlFmt)
		return imgbs64, nil
	}
}

func TrySDI2I(bs64, prompt string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	imgs := []string{bs64}
	// POST JSON string
	// No need to set content type, if you have client level setting
	r := &SDImageGenerationResponseBody{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(SDI2ImageGenerationRequestBody{InitImages: imgs, Prompt: prompt + "masterpiece, best quality, ultra-detailed", NegativePrompt: "EasyNegative", SamplerIndex: "DPM++ 2M Karras", Steps: 40}).
		SetResult(r).
		Post(SDI2IBASEURL)
	if err != nil {
		return "", err
	} else {
		if len(r.Images) == 0 {
			return "", errors.New("sd resp: " + resp.String())
		}
		fmt.Println("bs64 str", r.Images[0][0:20])
		return r.Images[0], nil
	}

}

func TrySuperResolution(image string, upscale float64, sdBaseUrlFmt string) (string, error) {
	// 创建一个Resty客户端
	client := resty.New()
	// imgs := []string{image}
	r := &SDI2ISuperResolutionResponseBody{}
	// 定义超分接口请求的参数
	reqBody := SDI2ISuperResolutionRequestBody{
		ResizeMode:      0,
		UpscalingResize: upscale,
		Upscaler1:       "R-ESRGAN 4x+ Anime6B",
		Image:           image,
	}

	// 发送POST请求，并获取响应结果
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetResult(r).
		Post(fmt.Sprintf(sdBaseUrlFmt, "extra-single-image"))

	if err != nil {
		return "", err
	} else {
		if len(r.Image) == 0 {
			return "", errors.New("sd resp: " + resp.String())
		}
		fmt.Println("bs64 str", r.Image[0:20])
		return r.Image, nil
	}
}

func TryCLIPINFO(bs64 string) (string, error) {
	// Create a Resty Client
	client := resty.New()
	// POST JSON string
	// No need to set content type, if you have client level setting
	r := &SDClipINFOResptBody{}
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(SDPNGINFORequestBody{Image: bs64, Model: "clip"}).
		SetResult(r).
		Post(SDCLIPBASEURL)
	if err != nil {
		return resp.String(), err
	} else {
		if r.Caption != "" {
			return r.Caption, nil
		} else {
			return resp.String(), err
		}

	}

}
