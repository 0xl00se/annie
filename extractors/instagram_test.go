package extractors

import (
	"testing"

	"github.com/iawia002/annie/config"
	"github.com/iawia002/annie/test"
)

func TestInstagram(t *testing.T) {
	config.InfoOnly = true
	tests := []struct {
		name string
		args test.Args
	}{
		{
			name: "Video",
			args: test.Args{
				URL:   "https://www.instagram.com/p/BYQ0PMWlAQY",
				Title: "王薇雅🇨🇳🇺🇸 on Instagram：“我的Ins是用来分享#lifestyle 一些正能量健身旅游等，请那些负能量离我远点！谢谢😀😀BTW，我从来不否认我P图微调，谁都想展现自己最完美的一面在网上.不要再给我粉丝私信黑我了，那么空能不能多读书看报增加些涵养🙂🙂🙂”",
				Size:  1469037,
			},
		},
		{
			name: "Image Single",
			args: test.Args{
				URL:   "https://www.instagram.com/p/Bei7whzgfMq",
				Title: "王薇雅🇨🇳🇺🇸 on Instagram：“Let go of what u can no longer keep. Protect what’s still worth keeping. ✨✨✨”",
				Size:  144348,
			},
		},
		{
			name: "Image Album",
			args: test.Args{
				URL:   "https://www.instagram.com/p/BdZ7sPTgchP",
				Title: "王薇雅🇨🇳🇺🇸 on Instagram：“2018的第一餐，吃得很满足🐷#happynewyear #🎆 #🎊”",
				Size:  10353828,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Instagram(tt.args.URL)
			test.Check(t, tt.args, data)
		})
	}
}
