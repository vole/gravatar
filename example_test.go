package gravatar_test

import (
  "bytes"
  "fmt"
  gr "github.com/ftrvxmtrx/gravatar"
  "image"
  _ "image/png"
)

func ExampleGetAvatar() {
  // get avatar image (128x128) using HTTP transport
  emailHash := gr.EmailHash("ftrvxmtrx@gmail.com")
  raw, err := gr.GetAvatar("http", emailHash, 128)

  // get avatar image (default size, png format) with fallback to "retro"
  // generated avatar.
  // use https transport
  emailHash = "cfcd208495d565ef66e7dff9f98764da.png"
  raw, err = gr.GetAvatar("https", emailHash, gr.DefaultRetro)

  if err == nil {
    var cfg image.Config
    var format string

    rawb := bytes.NewReader(raw)
    cfg, format, err = image.DecodeConfig(rawb)
    fmt.Println(cfg, format)
  }
}

func ExampleGetProfile() {
  // get profile using HTTPS transport
  emailHash := gr.EmailHash("ftrvxmtrx@gmail.com")
  profile, err := gr.GetProfile("https", emailHash)

  if err == nil {
    fmt.Println(profile.PreferredUsername)
    fmt.Println(profile.ProfileUrl)
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
