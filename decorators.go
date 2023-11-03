package draftjs

import "fmt"

type Decorator interface {
	RenderBeginning(data map[string]string) string
	RenderEnding(data map[string]string) string
}

type LinkDecorator struct {
}

func (decorator *LinkDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("<a href=\"%s\" target=\"_blank\">", data["url"])
}

func (decorator *LinkDecorator) RenderEnding(data map[string]string) string {
	return "</a>"
}

type ImageDecorator struct {
}

func (decorator *ImageDecorator) RenderBeginning(data map[string]string) string {
	if alt, ok := data["alt"]; ok {
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", data["data"], alt)
	}
	return fmt.Sprintf("<img src=\"%s\">", data["data"])
}

func (decorator *ImageDecorator) RenderEnding(data map[string]string) string {
	return "</img>"
}

type BlockImageDecorator struct {
}

func (decorator *BlockImageDecorator) RenderBeginning(data map[string]string) string {
	if alt, ok := data["alt"]; ok {
		return fmt.Sprintf("<img src=%s alt=\"%s\">", data["data"], alt)
	}
	return fmt.Sprintf("<figure><img src=%s>", data["data"])
}

func (decorator *BlockImageDecorator) RenderEnding(data map[string]string) string {
	return "</img></figure>"
}

type InlineImageDecorator struct {
}

func (decorator *InlineImageDecorator) RenderBeginning(data map[string]string) string {
	if alt, ok := data["alt"]; ok {
		return fmt.Sprintf("<img src=%s alt=\"%s\">", data["data"], alt)
	}
	return fmt.Sprintf("<img src=%s>", data["data"])
}

func (decorator *InlineImageDecorator) RenderEnding(data map[string]string) string {
	return "</img>"
}

type BlockAudioDecorator struct {
}

func (decorator *BlockAudioDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("<figure><audio controls><source src=\"%s\" type=\"audio/mpeg\">", data["data"])
}

func (decorator *BlockAudioDecorator) RenderEnding(data map[string]string) string {
	return "</audio></figure>"
}

type InlineAudioDecorator struct {
}

func (decorator *InlineAudioDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("<audio controls><source src=\"%s\" type=\"audio/mpeg\">", data["data"])
}

func (decorator *InlineAudioDecorator) RenderEnding(data map[string]string) string {

	return "</audio>"
}

type AudioDecorator struct {
}

func (decorator *AudioDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("<audio controls><source src=\"%s\" type=\"audio/mpeg\">", data["data"])
}

func (decorator *AudioDecorator) RenderEnding(data map[string]string) string {
	return "</audio>"
}

type BlockMathJaxDecorator struct {
}

func (decorator *BlockMathJaxDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("\\[%s\\]", data["data"])
}

func (decorator *BlockMathJaxDecorator) RenderEnding(data map[string]string) string {
	return ""
}

type InlineMathJaxDecorator struct {
}

func (decorator *InlineMathJaxDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("\\(%s\\)", data["data"])
}

func (decorator *InlineMathJaxDecorator) RenderEnding(data map[string]string) string {
	return ""
}
