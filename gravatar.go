package gravatar

import (
  "crypto/md5"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "strconv"
  "strings"
)

// GravatarProfile stores profile information associated with an email.
type GravatarProfile struct {
  AboutMe string

  Accounts []struct {
    ShortName string
    Domain    string
    Url       string
    Verified  bool `json:",string"`
    Username  string
    Display   string
  }

  CurrentLocation string

  DisplayName string

  Emails []struct {
    Primary bool `json:",string"`
    Value   string
  }

  Hash string

  Id int `json:",string"`

  Ims []struct {
    Type  string
    Value string
  }

  Name struct {
    Family    string
    Formatted string
    Given     string
  }

  PhoneNumbers []struct {
    Type  string
    Value string
  }

  Photos []struct {
    Type  string
    Value string
  }

  PreferredUsername string

  ProfileBackground struct {
    Color    string
    Position string
    Repeat   string
  }

  ProfileUrl string

  ThumbnailUrl string

  Urls []struct {
    Title string
    Value string
  }
}

const gravatarHost = "gravatar.com"

// List of (optional) values for a "default action" option for GetAvatar.
// Each option defines what GetAvatar has to do in case of non-existing
// user.
const (
  // DefaultError defaults to an error.
  DefaultError = "404"

  // DefaultIdentIcon defaults to a generated geometric pattern.
  DefaultIdentIcon = "identicon"

  // DefaultMonster defaults to a generated 'monster' with different colors
  // and faces.
  DefaultMonster = "monsterid"

  // DefaultMysteryMan defaults to a simple, cartoon-style silhouetted outline
  // of a person.
  DefaultMysteryMan = "mm"

  // DefaultRetro defaults to a generated 8-bit arcade-style pixelated faces.
  DefaultRetro = "retro"

  // DefaultWavatar defaults to a generated faces with differing features and
  // backgrounds.
  DefaultWavatar = "wavatar"
)

var client = new(http.Client)

// EmailHash converts an email to lowercase and returns its MD5 hash as hex
// string.
func EmailHash(email string) string {
  m := md5.New()
  io.WriteString(m, strings.ToLower(email))
  return fmt.Sprintf("%x", m.Sum(nil))
}

// GetAvatar does a HTTP(S) request and returns an avatar image.
//
// Optional arguments include Default* (default actions) and image size.
func GetAvatar(scheme, emailHash string, opts ...interface{}) (data []byte, err error) {
  url := &url.URL{
    Scheme: scheme,
    Host:   gravatarHost,
    Path:   "/avatar/" + emailHash,
  }
  values := url.Query()

  for _, opt := range opts {
    switch o := opt.(type) {
    case string:
      values.Add("d", o)

    case int:
      values.Add("s", strconv.Itoa(o))
    }
  }

  url.RawQuery = values.Encode()

  err = run(url, get_avatar(&data))
  return
}

// GetProfile does a HTTP(S) request (based on `scheme` argument) and returns
// gravatar profile.
func GetProfile(scheme, emailHash string) (g GravatarProfile, err error) {
  url := &url.URL{
    Scheme: scheme,
    Host:   gravatarHost,
    Path:   "/" + emailHash + ".json",
  }

  err = run(url, unmarshal_json(&g))
  return
}

func get_avatar(dst *[]byte) func([]byte) error {
  return func(data []byte) (err error) {
    *dst = data[:]
    return
  }
}

func run(url *url.URL, f func([]byte) error) (err error) {
  var res *http.Response
  res, err = client.Get(url.String())

  if err == nil {
    var data []byte
    defer res.Body.Close()

    if data, err = ioutil.ReadAll(res.Body); err == nil {
      if res.StatusCode == http.StatusOK {
        err = f(data)
      } else {
        err = errors.New(string(data))
      }
    }
  }

  return
}

func unmarshal_json(g *GravatarProfile) func([]byte) error {
  return func(data []byte) (err error) {
    obj := struct {
      Entries []GravatarProfile `json:"entry"`
    }{[]GravatarProfile{}}

    if err = json.Unmarshal(data, &obj); err == nil && len(obj.Entries) > 0 {
      *g = obj.Entries[0]
    }

    return
  }
}

// Local Variables:
// indent-tabs-mode: nil
// tab-width: 2
// End:
// ex: set tabstop=2 shiftwidth=2 expandtab:
