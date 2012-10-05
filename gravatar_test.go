package gravatar

import (
  "testing"
)

// TODO(ftrvxmtrx): write real tests

func TestEmailHash(t *testing.T) {
  h := EmailHash("ftrvxmtrx@gmail.com")

  if h != "d96ba36eb0d406aea53f3868cd06fca8" {
    t.Error(h)
  }
}

func TestGetAvatar(t *testing.T) {
  for _, scheme := range []string{"http", "https"} {
    if a, err := GetAvatar(scheme, "d96ba36eb0d406aea53f3868cd06fca8"); err != nil {
      t.Error(a, err)
    }

    if a, err := GetAvatar(scheme, "0"); err != nil {
      t.Error(a, err)
    }

    if a, err := GetAvatar(scheme, "0", DefaultError); err == nil {
      t.Error(a, err)
    }

    if a, err := GetAvatar(scheme, "0.png", DefaultIdentIcon, 256); err != nil {
      t.Error(a, err)
    }

    if a, err := GetAvatar(scheme, "0.png", RatingX, DefaultIdentIcon, 256); err != nil {
      t.Error(a, err)
    }
  }
}

func TestGetAvatarURL(t *testing.T) {
  if url := GetAvatarURL("http", "d96ba36eb0d406aea53f3868cd06fca8"); url == nil {
    t.Error(url)
  }
}

func TestGetProfile(t *testing.T) {
  for _, scheme := range []string{"http", "https"} {
    if e, err := GetProfile(scheme, "d96ba36eb0d406aea53f3868cd06fca8"); err != nil {
      t.Error(e, err)
    }

    if e, err := GetProfile(scheme, "0"); err == nil {
      t.Error(e, err)
    }
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
