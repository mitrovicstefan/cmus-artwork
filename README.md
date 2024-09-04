# Cmus pixel artwork

![image](https://github.com/user-attachments/assets/78eb33d0-513e-43a8-bf89-fff469f22d1a)

Small [cmus](https://github.com/cmus/cmus) util for displaying artwork in pixelated form. Also allows pause/play/next/previous commands. The purpose is exclusively aesthetic.
Only tested and supported on linux and mac.

It will print `cover.jpg` or `cover.png` or `cover.webp` image from a folder where your song is located as pixelated ascii image art. It refreshes every 10 secs so if your artwork doesn't change give it up to 10 seconds.

## Installation

Download the latest release using wget or curl

```
wget https://github.com/mitrovicstefan/cmus-artwork/releases/download/v1.0.0/cmus-artwork
```

```
curl https://github.com/mitrovicstefan/cmus-artwork/releases/download/v1.0.0/cmus-artwork
```

Move the file to your local binaries

```
sudo mv cmus-artwork /usr/local/bin/
```

Run it!

```
cmus-artwork
```

## Commands
You can use `j` to pause/play, `h` to rewind/play previous and `l` to play next song. `CMD + C` or `CTRL + C` closes the app.

## Notes
- If your artwork looks broken into random lines, you either resized the terminal or your terminal aspect ratio does not suit the cover art. If you just resized the terminal, give it up to 10 seconds to refresh the cover art and adjust to new dimensions. If it's broken and you didn't resize it, your terminals aspect ratio is off. Increase width until your cover art fits, and give it up to 10 seconds to auto-refresh.
- If you see this guy -  `¯\_(ツ)_/¯` it means:
  1. cmus is not open
  2. there is no `cover.png/cover.jpg/cover.webp` in the same directory where the song is.

## Config
There is currently no conf file support. If people actually use this, I'll expand the functionality and create config for custom key bindings, custom character maps, grayscale support, themes etc. Please star this project and/or write feature requests so I know that people use this. Thank you!

## Building and development

If you wish to build the binary yourself, it's extremely easy! First make sure that you have [go installed locally](https://go.dev/doc/install).

Once it's installed, clone the repository -

```
git clone https://github.com/mitrovicstefan/cmus-artwork
cd cmus-artwork
```

Run `main.go` to automatically install all the dependencies -

```
go run main.go
```

To build a binary run -

```
go build
```

From here you can start developing and messing around with the code. PRs are ALWAYS welcome!

## Feature requests and bug reports

To request a feature or report a bug open a new issue in this repo. I'll get to it in my free time :)

## Thank you
[moby/term](https://github.com/moby/term/tree/main) codebase for mac support\
[TheZoraiz/ascii-image-converter](https://github.com/TheZoraiz/ascii-image-converter) is library I use to convert images to ascii
